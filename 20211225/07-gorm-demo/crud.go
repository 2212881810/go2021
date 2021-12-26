package main

import "fmt"

func create_user() {

	//var gender = "abc"
	//res := GLOBAL_DB.Create(&User{Name: "郑钦锋", Age: 25, Gender: &gender})

	//
	//user.ID             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数

	//var gender = "abc"
	//// 只创建Name和Gender字段
	//res := GLOBAL_DB.Select("Name", "Gender").Create(&User{Name: "郑钦锋", Gender: &gender})

	//var gender = "abc"
	//// 除了Age之外的都创建
	//res := GLOBAL_DB.Omit("Age").Create(&User{Gender: &gender})

	// 批量插入
	var gender = "abc"
	var users = []User{{Name: "admin_01", Gender: &gender}, {Name: "admin_02", Gender: &gender}, {Name: "admin_03", Gender: &gender}}
	res := GLOBAL_DB.Create(&users)
	// 批量创建100条数据
	//GLOBAL_DB.CreateInBatches(users, 100)

	if res.Error != nil {
		fmt.Println("create error: ", res.Error)
		return
	}

	fmt.Println(res.RowsAffected)
}

func find_user() {
	u := User{}
	// 查询表中的第一条数据 ，默认按主键排序，如果没有主键，按表中第1列排序
	res := GLOBAL_DB.First(&u)
	fmt.Printf("查询到%d条数据\n", res.RowsAffected)
	fmt.Println(u)

	// 随便取一条数据，怎么一直是第1条？？？？
	u2 := User{}
	GLOBAL_DB.Take(&u2)
	fmt.Println(u2)

	// 取表中最后一条数据
	u3 := User{}
	GLOBAL_DB.Last(&u3)
	fmt.Println(u3)

	// 按条件进行查询
	u4 := User{ID: 4}
	GLOBAL_DB.Find(&u4)
	fmt.Println(u4)

	// 按条件查询多条数据，使用字符串进行查询拼接， 用切片进行装结果
	var users []User
	GLOBAL_DB.Where("name = ?", "郑钦锋").Find(&users)
	fmt.Println(users)

	// 按条件查询，使用struct进行条件过滤，（效果我上面一样）
	var user2s []User
	GLOBAL_DB.Where(&User{Name: "郑钦锋"}).Find(&user2s)
	fmt.Println(user2s)

	// 查询指定的字段属性
	var user3s []User
	GLOBAL_DB.Select("Age", "CreatedAt").Where(&User{Name: "admin_01"}).Find(&user3s)
	fmt.Println(user3s)

	// 使用原生sql查询
	type Result struct {
		Name string
		Age  int
	}
	// 如果只有一个结果用struct接收就行， 如果有多个结果，使用切片
	var result []Result
	GLOBAL_DB.Raw("select name , age from users where name = ? ", "郑钦锋").Scan(&result)
	fmt.Println(result)

	// 模糊查询
	var user4s []User
	// 这种查询Select("Name")好像没生效
	GLOBAL_DB.Where("name like ?", "admin%").Find(&user4s)
	fmt.Println(user4s)
}
