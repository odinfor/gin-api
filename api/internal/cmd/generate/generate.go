package main

// gorm gen 使用demo
// 执行main方法，通过gorm gen生成相关代码和模型的映射关系.
// 执行后在OutPath路径下生成model和query两个目录.
// demo_user是因为database中已经存在的表同步所产生。mytables是ddl同步所产生

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
	"gorm.io/gorm"
)

const mytableSQL = "CREATE TABLE IF NOT EXISTS `mytables` (" +
	"    `ID` int(11) NOT NULL," +
	"    `username` varchar(16) DEFAULT NULL," +
	"    `age` int(8) NOT NULL," +
	"    `phone` varchar(11) NOT NULL," +
	"    INDEX `idx_username` (`username`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"

func prepare(db *gorm.DB) {
	db.Exec(mytableSQL)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../api/internal/dal/query",
	})

	dal.DB = dal.ConnectDB("root:root@tcp(localhost:3306)/demo-user?charset=utf8mb4&parseTime=True").Debug()
	g.UseDB(dal.DB)

	// 生成所有的模型。通过表生成对应的映射关系.
	// g.GenerateModel("people") -- generate a model struct map to table `people` in database
	// g.GenerateModelAs("people", "People") -- generate a struct and specify struct's name
	// g.GenerateModel("people", gen.FieldIgnore("address"), gen.FieldType("id", "int64")) -- add option to ignore field
	// g.GenerateAllTable() -- generate all tables, ex: g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
