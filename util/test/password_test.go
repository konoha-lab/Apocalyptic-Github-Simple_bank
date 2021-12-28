package util

import (
	_util "simple_bank/util"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := _util.RandomString(6)

	hashedPassword, err := _util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = _util.CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestWrongPassword(t *testing.T) {
	password := _util.RandomString(6)

	hashedPassword, err := _util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	wrongPassword := _util.RandomString(6)
	err = _util.CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
