package logic

import (
	"errors"
	"gin-api/internal/model"
	"gin-api/internal/svc"
	"gin-api/internal/types"
	"go.uber.org/zap"
)

func AddOne(a int) (int, error) {
	if a < 10 || a > 100 {
		return 0, errors.New("超过范围")
	}
	return a + 1, nil
}

func GetMap(a int) (map[string]string, error) {
	if a < 10 || a > 100 {
		return nil, errors.New("超过范围")
	}
	return map[string]string{
		"aa":   "cccc",
		"key2": "this key2's value",
	}, nil
}

func GetSlice(a int) ([]int, error) {
	if a < 10 || a > 100 {
		return nil, errors.New("超过范围")
	}
	return []int{1, 2, 3, 4}, nil
}

func GetFloat(a int) (float32, error) {
	if a < 10 || a > 100 {
		return 0, errors.New("超过范围")
	}
	return 3.543, nil
}

func GetInterface(a int) (interface{}, error) {
	if a < 10 || a > 100 {
		return 0, errors.New("超过范围")
	}
	return map[string]interface{}{
		"key1": []int{1, 2, 3, 4},
		"key2": "this is key2's value",
		"key3": map[string]string{
			"a": "1111",
			"b": "2222",
		},
	}, nil
}

type UserLogic struct {
	svcCtx *svc.ServiceContext
}

func NewUserLogic(svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		svcCtx: svcCtx,
	}
}

func (u *UserLogic) AddUser(params *types.AddUserRequest) error {
	user := &model.DemoUser{
		Name:    params.Name,
		Address: params.Address,
		Age:     params.Age,
	}
	zap.L().Info("this is test add user info zap log")
	if _, err := u.svcCtx.DemoUserModel.Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserLogic) DiffUser(params *types.DiffUserRequest) error {
	return u.svcCtx.DemoUserModel.DeleteById(params.Id)
}
