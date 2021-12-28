package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	mockapi "simple_bank/db/mock"
	"simple_bank/db/models"
	_repo "simple_bank/db/repository"
	_util "simple_bank/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user models.User, password string) {
	pass := "secret"
	_, err := _util.HashPassword(pass)
	require.NoError(t, err)

	return models.User{
		Username: _util.RandomString(6),
		FullName: _util.RandomString(12),
		Email:    _util.RandomEmail(),
	}, pass
}

type eqCreateUserParamsMatcher struct {
	arg      _repo.CreateUserParams
	password string
}

func EqCreateUserParams(arg _repo.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(_repo.CreateUserParams)
	if !ok {
		return false
	}

	err := _util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)

	hashedPassword, err := _util.HashPassword(password)
	require.NoError(t, err)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(handler *mockapi.MockHandler)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				arg := _repo.CreateUserParams{
					Username:       user.Username,
					HashedPassword: hashedPassword,
					FullName:       user.FullName,
					Email:          user.Email,
				}
				handler.EXPECT().
					// CreateUser(gomock.Any(), gomock.Eq(arg)).
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(models.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(models.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username":  "invalid-user#1",
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     "invalid-email",
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"username":  user.Username,
				"password":  "123",
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := mockapi.NewMockHandler(ctrl)
			tc.buildStubs(handler)

			server := newTestServer(t, handler)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user models.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser models.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
	// fmt.Println(gotAccount)
}
