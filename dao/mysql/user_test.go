package mysql

import (
	"reflect"
	"testing"
	"web-study/entity"
	"web-study/settings"
)

func TestClose(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Close()
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		cfg *settings.MySQLConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUp(t *testing.T) {
	type args struct {
		p *entity.ParamLoginUp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoginUp(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("LoginUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMd5Password(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Password(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("Md5Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectByUsername(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
