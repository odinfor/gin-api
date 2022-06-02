package logic

import (
	"gin-api/internal/svc"
	"gin-api/internal/types"
	"testing"
)

func TestUserLogic_AddUser(t *testing.T) {
	type fields struct {
		svcCtx *svc.ServiceContext
	}
	type args struct {
		params *types.AddUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "测试添加用户",
			fields: fields{
				svcCtx: svc.NewServiceContext(),
			},
			args: args{params: &types.AddUserRequest{
				Name:    "测试用户1",
				Address: "测试用户地址1",
				Age:     20,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLogic{
				svcCtx: tt.fields.svcCtx,
			}
			if err := u.AddUser(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserLogic_DiffUser(t *testing.T) {
	type fields struct {
		svcCtx *svc.ServiceContext
	}
	type args struct {
		params *types.DiffUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "测试删除用户用例",
			fields: fields{
				svcCtx: svc.NewServiceContext(),
			},
			args:    args{params: &types.DiffUserRequest{Id: 10}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLogic{
				svcCtx: tt.fields.svcCtx,
			}
			if err := u.DiffUser(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("DiffUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
