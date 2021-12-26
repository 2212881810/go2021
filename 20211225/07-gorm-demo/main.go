package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBAL_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(192.168.79.70:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		//DefaultStringSize:         256,                                                                                 // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false,                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		//SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "tb_", // 表名前缀，`User`表为`t_users`
		//	SingularTable: false, // 使用单数表名，启用该选项后，`User` 表将是`user`
		//	//NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		//	NoLowerCase: false, // 大写,tb_Users这种，还是设置成false吧
		//},
		//DisableAutomaticPing:                     true, // 禁用ping去检测数据库的可用性
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束,如果不设置使用AutoMigrate自动建表，会创建外键约束
	})

	if err != nil {
		fmt.Println("连接mysql 错误:", err)
		return
	}

	//fmt.Println(db)
	//
	//type User struct {
	//	Name string
	//}
	//
	//type UserInfo struct {
	//	Name string
	//}
	//
	//type Student struct {
	//	Name   string
	//	Age    int
	//	gender bool
	//}
	//
	//// 连接池配置
	//sqlDB, err := db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//
	////stats := sqlDB.Stats()
	////fmt.Println(stats)
	//
	//// 1. 创建表
	//db.AutoMigrate(&User{}, &Student{})
	//// 建表时添加后缀， 此处创建两张表
	////db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Student{})
	//migrator := db.Migrator()
	//
	//// 2. 判断是否存在表
	//fmt.Println(migrator.HasTable(&User{}))
	//
	//// 3. 获取当前数据库名称
	//fmt.Println(migrator.CurrentDatabase())
	//
	//// 4. 如果存在表则删除（删除时会忽略、删除外键约束)
	////db.Migrator().DropTable(&User{})
	////db.Migrator().DropTable("user_infos")
	//
	//// 5. 重命名表
	//if !migrator.HasTable(&User{}) {
	//	db.Migrator().RenameTable(&User{}, &UserInfo{})
	//}
	//
	//if !migrator.HasTable(&UserInfo{}) {
	//	db.Migrator().RenameTable(&UserInfo{}, &User{})
	//}

	//db.Migrator().RenameTable("tb_users", "user_infos")
	GLOBAL_DB = db

	//fmt.Println("----------------------------")
	// 1.创建表
	//createUser()

	// 2. add 一条数据
	//create_user()

	// 3. 查询数据
	//find_user()

	// 4. 一对一的关系
	//belongsTo()
	// 4.1 belongsTo 新增数据
	belongsToCreate()
}
