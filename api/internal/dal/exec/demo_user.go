package exec

// 演示demo curd示例

import (
	"context"
	"fmt"
	"gin-api/internal/dal/model"
	"gin-api/internal/dal/query"
	"gorm.io/gen/examples/dal"
)

func Create(ctx context.Context) error {
	q := query.Use(dal.ConnectDB("root:root@tcp(localhost:3306)/demo-user?charset=utf8mb4&parseTime=True").Debug())
	return q.WithContext(ctx).DemoUser.Create(&model.DemoUser{Name: "aa", Address: "cc", Age: 10})
}

func QueryLast(ctx context.Context) error {
	q := query.Use(dal.ConnectDB("root:root@tcp(localhost:3306)/demo-user?charset=utf8mb4&parseTime=True").Debug())
	u, err := q.WithContext(ctx).DemoUser.Last()
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}

func Del(ctx context.Context, id int32) error {
	q := query.Use(dal.ConnectDB("root:root@tcp(localhost:3306)/demo-user?charset=utf8mb4&parseTime=True").Debug())
	rsp, err := q.WithContext(ctx).DemoUser.DO.Where(q.DemoUser.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	fmt.Println(rsp)
	return nil
}
