package main

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/router"
)

func main() {

	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "root:123Zhaozelong@tcp(139.159.220.124:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local",
	//	DefaultStringSize:         256,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	//
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//a, _ := gormadapter.NewAdapterByDB(db)
	//e, _ := casbin.NewEnforcer("conf/rbac_model.conf", a)
	//
	//e.LoadPolicy()
	//
	//e.Enforce("alice", "data1", "read")
	//
	//e.SavePolicy()

	r := router.InitRouter()

	r.Run(":8080")
}
