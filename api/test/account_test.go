package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockapi "simple_bank/db/mock"
	"simple_bank/db/models"
	"simple_bank/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handle := mockapi.NewMockHandler(ctrl)
	handle.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := newTestServer(t, handle)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
	// fmt.Println(account)
}

func TestGetAccountAPI2(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(handler *mockapi.MockHandler)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.ID,
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(models.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(models.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStubs: func(handler *mockapi.MockHandler) {
				handler.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
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

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account models.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount models.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
	// fmt.Println(gotAccount)
}

func randomAccount() models.Account {
	return models.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
