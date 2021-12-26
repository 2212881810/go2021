package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	// 调用了bookabledate
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

/**
验证函数
*/
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {

	date, ok := fl.Field().Interface().(time.Time)

	if ok { // ok为true, 说明fl的field字段可以转换成time.Time类型，然后对其进行验证

		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main2() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//创建一个新的validator指针
		//validate := validator.New()

		// 将bookableDate这个函数注册到Validate实例上去
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
