package main

import (
	"gorm.io/gorm"
)

// User /*
// belongs to : 一对一的模型
//*

type Boy struct {
	gorm.Model
	Name      string
	CompanyID int // 通过这行代码，会自动建立Boy 与Company的逻辑外键关系
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func belongsTo() {
	// 1 建表:gorm会自动创建Boy和Company两张表，并且建立好这两张表之间的关联关系
	GLOBAL_DB.AutoMigrate(&Boy{})
}

/**
在添加Boy数据时，gorm会自动帮我们添加一条company的数据，并建立好他们之间的映射关系
*/
func belongsToCreate() {

	/*

		// 清空这两张表中的数据,Unscoped 表示永久性删除
		// gorm 默认不会执行全表删除操作，若想全表删除，可以采用以下几种方式

		// 1. 通过where 1=1 来执行, 测试不生效
		//GLOBAL_DB.Unscoped().Where("1 = 1").Delete(&Boy{}, &Company{})

		// 2. 使用原生sql没有问题，而且是真正的删除，并非逻辑删除
		//GLOBAL_DB.Exec("delete from boys")
		//GLOBAL_DB.Exec("delete from companies")

		// 3. 第三种全表数据删除的方式，注意：只是逻辑删除
		//GLOBAL_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Boy{}, &Company{})
		// 物理删除，但是只能删除前面那张表，好奇怪
		//GLOBAL_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Company{}, &Boy{})

	*/

	GLOBAL_DB.Exec("delete from boys")
	GLOBAL_DB.Exec("delete from companies")

	//tx := GLOBAL_DB.Unscoped().Delete(&Boy{}, &Company{})
	//if tx.Error != nil {
	//	fmt.Println("delete all data error :", tx.Error)
	//}

	//b := Boy{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name:      "郑钦锋",
	//	CompanyID: 1,
	//	Company: Company{
	//		ID:   1,
	//		Name: "平安科技",
	//	},
	//}
	//// add b时，gorm会自动建立好两张表中数据的映射关系
	//GLOBAL_DB.Create(&b)

	c1 := Company{
		ID:   2,
		Name: "中国银行",
	}
	b1 := Boy{
		Name:    "苟大爷",
		Company: c1, // 通过对象建立关系
	}
	// 表真的建立了b1和c1之间的外键关联关系,垃圾.
	// 所以一定要将DisableForeignKeyConstraintWhenMigrating 设置成true
	GLOBAL_DB.Create(&b1)
	//
	//// 直接指定id，新增数据
	//b2 := Boy{
	//	Name:      "苟小爷",
	//	CompanyID: 2,
	//}
	//GLOBAL_DB.Create(&b2)
	//
	//// 只添加指定的列数据
	b3 := Boy{
		Name:      "小明",
		CompanyID: 2,
	}
	// 只会添加Name列表数据，也不会建立该boy与company之间的关联关系！
	GLOBAL_DB.Select("Name").Create(&b3)
}
