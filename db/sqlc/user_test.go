package db

import (
	"context"
	"sync"
	"testing"
)

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		ID:    1,
		Name:  "user1",
		Email: "122",
		Money: 100,
	}
	user, err := testQueries.CreateUser(context.Background(), args)
	if err != nil {
		t.Log("create user error", err.Error())
	}
	t.Log(user)

	args = CreateUserParams{
		ID:    2,
		Name:  "user2",
		Email: "122",
		Money: 300,
	}
	user, err = testQueries.CreateUser(context.Background(), args)
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

func TestAddMoney(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 1)
	if err != nil {
		t.Log("获取用户失败")
	}
	t.Log(user)
	user, err = testQueries.UpdateUserMoney(context.Background(), UpdateUserMoneyParams{
		Money: 10,
		ID:    1,
	})
	if err != nil {
		t.Log("更新失败")
	}
	t.Log(user)
}

func TestStore_TransferMoney(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			err := TestStore.TransferMoney(1, 2, 100)
			if err != nil {
				t.Log(err)
			} else {
				t.Log("转账成功")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	users, _ := testQueries.ListUsers(context.Background())
	t.Log("最后账户信息", users)
}
