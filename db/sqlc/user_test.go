package db

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		ID:    1,
		Name:  "xiaoming",
		Email: "11232",
	}
	user, err := testQueries.CreateUser(context.Background(), args)
	if err != nil {
		t.Log("create user error", err.Error())
	}
	t.Log(user)
}
