package test

import (
	"fmt"
	"mini-douyin/dao"
	"mini-douyin/models"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		user models.UserDao
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "demo01", args: args{user: models.UserDao{UserName: "www", Password: "w123"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dao.Register(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindOneSimple(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{name: "demo01", args: "asdsa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := dao.FindOneSimple(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("FindOneSimple() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//根据用户的id查找用户
func TestSelectUserById(t *testing.T) {
	var userId int64
	userId = 2
	user, err := dao.SelectUserByID(userId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)
}
