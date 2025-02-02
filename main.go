package main

import (
	"net/http"
	"server-vanstartup/auth"
	"server-vanstartup/campaign"
	"server-vanstartup/handler"
	"server-vanstartup/helper"
	"server-vanstartup/payment"
	"server-vanstartup/transaction"
	"server-vanstartup/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	authService := auth.NewJWTService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()

	r.Static("/images", "./images")

	api := r.Group("/api/v1")

	api.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to api VanStartup",
		})
	})

	api.POST("/users", userHandler.RegisteUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email-checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService) , userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService) , campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	r.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "FAILED", nil))
			return
		}
	
		tokenString := ""
		arrToken := strings.Split(tokenString, "Bearer ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		jwtToken, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "FAILED", nil))
			return
		}

		claim, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "FAILED", nil))
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "FAILED", nil))
			return
		}

		c.Set("currentUser", user)
	}
}
