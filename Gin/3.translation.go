package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func main() {
	if err := InitializeTranslate(); err != nil {
		panic(err)
	}

	g := gin.Default()

	g.POST("login", LoginHandler)
	g.POST("register", RegisterHandler)
	g.Run(":1234")
}

func RegisterHandler(c *gin.Context) {
	var r RegisterForm

	err := c.ShouldBindJSON(&r)

	if err != nil {
		// error类型判断
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, gin.H{
				"code":  40001,
				"msg":   "注册失败",
				"error": err.Error(),
			})
			return
		}
		// error 是ValidationErrors类型的
		c.JSON(200, gin.H{
			"code":  40004,
			"error": err.Translate(translator),
		})
		return
	}
	c.JSON(200, gin.H{
		"Code": 0,
		"Msg":  "success!",
		"data": r,
	})

}

type LoginForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re-password" binding:"eqfield=password"`

	//UserName string `json:"username" validate:"required,min=3,max=7"`
	//Password string `json:"password" validate:"required,len=8"`
	//
	// eqfield=password : 表示re-password 必须与password相等
	//RePassword string `json:"re-password" validate:"eqfield=password"`
	//RePassword string `json:"re-password" binding:"required,len=8"`
}

type RegisterForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	Password string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func LoginHandler(c *gin.Context) {
	var l LoginForm
	err := c.ShouldBindJSON(&l)

	if err != nil {
		//c.JSON(200, gin.H{
		//	"Code": 4001,
		//	"Msg":  "登录失败,请检测",
		//})
		// translate all error at once

		errs := err.(validator.ValidationErrors)
		c.JSON(200, gin.H{
			"Code": 4001,
			"Msg":  errs.Translate(translator),
		})
		return
	}

	c.JSON(200, gin.H{
		"Code": 0,
		"Msg":  "success!",
		"data": l,
	})
}

var translator ut.Translator

func InitializeTranslate() (err error) {

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 默认的TagName是binding,此处可以改成validate, 需要与结构体LoginForm一致，否则校验不会生效
		//validator.SetTagName("validate")

		validator.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("json") // 通过反射的方式获取字段的tag中的json值!!!
		})

		z := zh.New()
		uni := ut.New(z, z)
		// 把获取到translator赋值给全局
		translator, _ = uni.GetTranslator("zh")
		err = zh_translations.RegisterDefaultTranslations(validator, translator)
		return err
	}
	return err

	// 下面这样不行！！！！！！！！！！

	//validator := validator.New()
	//validator.SetTagName("binding")
	//
	//validator.RegisterTagNameFunc(func(field reflect.StructField) string {
	//	return field.Tag.Get("json") // 通过反射的方式获取字段的tag中的json值!!!
	//})
	//
	//z := zh.New()
	//uni := ut.New(z, z)
	//// 把获取到translator赋值给全局
	//translator, _ = uni.GetTranslator("zh")
	//err = zh_translations.RegisterDefaultTranslations(validator, translator)
	//return err
}
