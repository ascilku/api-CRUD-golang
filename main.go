package main

import (
	"api-satu/auth"
	"api-satu/campaign"
	"api-satu/handler"
	"api-satu/respons"
	"api-satu/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/resauceDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// migration
		// db.Debug().AutoMigrate(&user.User{})
		// db.Debug().AutoMigrate(&campaign.Campaign{})
		// db.Debug().AutoMigrate(&campaign.CampaignImage{})
		// repository user
		newRepository := user.NewRepository(db)
		newService := user.NewService(newRepository)
		newAuth := auth.NewAuth()
		newHandler := handler.NewHandler(newService, newAuth)
		//repository campaign
		campaignRepository := campaign.NewRepository(db)
		campaignService := campaign.NewService(campaignRepository)
		campaignHandlerCampign := handler.NewHandlerCampign(campaignService)
		//repository active campaign image

		// router
		router := gin.Default()
		api := router.Group("v1")
		api.POST("user", newHandler.CreateHandler)
		api.POST("auth-user", newHandler.AuthUserHandler)
		api.POST("check-email", newHandler.CheckEmailUserHandler)
		api.POST("update-image", authMiddleware(newAuth, newService), newHandler.UpdateImageUserHandler)
		// router campaign
		api.GET("campaign", campaignHandlerCampign.FindCampaignUserHand)
		router.Run()

	}
}

func authMiddleware(auth auth.Auth, service user.Service) gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			responsValue := respons.ResponsValue("Unauthorized Bearer", http.StatusUnauthorized, "Error", nil)
			g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
			return
		}
		var tokenString = ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		validationToken, err := auth.ValidationToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responsValue := respons.ResponsValue("Unauthorized Validate Token", http.StatusUnauthorized, "Error", errorMessage)
			g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
			return
		} else {
			claim, ok := validationToken.Claims.(jwt.MapClaims)
			if !ok || !validationToken.Valid {
				responsValue := respons.ResponsValue("Unauthorized", http.StatusUnauthorized, "Error", nil)
				g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
				return
			}
			userID := int(claim["user_id"].(float64))
			getUserByID, err := service.GetUserByID(userID)
			if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				responsValue := respons.ResponsValue("Unauthorize", userID, "Error", errorMessage)
				g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
				return
			}
			g.Set("current_user", getUserByID)
		}
		// responsValue := respons.ResponsValue(tokenString, http.StatusOK, "Error", nil)
		// g.AbortWithStatusJSON(http.StatusOK, responsValue)
		// return
	}
}

// func authMiddleware(auth auth.Auth, service user.Service) gin.HandlerFunc {
// 	return func(g *gin.Context) {
// 		authHeader := g.GetHeader("Authorization")
// 		if !strings.Contains(authHeader, "Bearer") {
// 			responsValue := respons.ResponsValue("Unautorized Bearer", http.StatusUnauthorized, "Error", nil)
// 			g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
// 			return
// 		}
// 		var tokenString = ""
// 		arrToken := strings.Split(authHeader, " ")
// 		if len(arrToken) == 2 {
// 			tokenString = arrToken[1]
// 		}
// 		validationToken, err := auth.ValidationToken(tokenString)
// 		if err != nil {
// 			errorMessage := gin.H{"error": err.Error()}
// 			responsValue := respons.ResponsValue("Unautorized validation", http.StatusUnauthorized, "Error", errorMessage)
// 			g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
// 			return
// 		} else {
// 			claim, ok := validationToken.Claims.(jwt.MapClaims)
// 			if !ok || !validationToken.Valid {
// 				responsValue := respons.ResponsValue("Unautorized", http.StatusUnauthorized, "Error", nil)
// 				g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
// 				return
// 			}
// 			userID := claim["user_id"].(float64)
// 			getUserByID, err := service.GetUserByID(int(userID))
// 			if err != nil {
// 				errorMessage := gin.H{"error": err.Error()}
// 				responsValue := respons.ResponsValue("Unautorized", http.StatusUnauthorized, "Error", errorMessage)
// 				g.AbortWithStatusJSON(http.StatusUnauthorized, responsValue)
// 				return
// 			}
// 			g.Set("current_user", getUserByID)
// 		}
// 	}
// }
