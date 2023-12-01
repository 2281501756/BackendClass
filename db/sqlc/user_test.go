package db

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		ID:    2,
		Name:  "xiaohu",
		Email: "122",
	}
	user, err := testQueries.CreateUser(context.Background(), args)
	if err != nil {
		t.Log("create user error", err.Error())
	}
	t.Log(user)
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 1)
	if err != nil {
		t.Log("no user")
	}
	t.Log(user)
}

func TestListUsers(t *testing.T) {
	users, err := testQueries.ListUsers(context.Background())
	if err != nil {
		t.Log("no users")
	}
	t.Log(users)
}
