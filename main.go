package main

import (
	"net/http"
	"server-vanstartup/handler"
	"server-vanstartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/vanstartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	api := r.Group("/api/v1")

	api.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to api VanStartup",
		})
	})

	api.POST("/users", userHandler.RegisteUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email-checkers", userHandler.CheckEmailAvaibility)

	r.Run()
}
