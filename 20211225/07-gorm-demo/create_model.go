package main

import (
	"time"

	"gorm.io/gorm"
)

//type User struct {
//	// 约束大于配置，gorm中， id默认会搞成主键
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     *time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

//type User struct {
//	gorm.Model // gorm帮我们搞了一个Model, 定义好主键id列，createdAt , updatedAt
//	Name       string
//	Age      int
//}

//https://gorm.io/zh_CN/docs/models.html#embedded_struct
// 等效于
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Age       int
	Gender    *string `gorm:"not null"` // 非空限制
}

func createUser() {

	GLOBAL_DB.AutoMigrate(&User{})

}
