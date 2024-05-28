package handler

import (
	"api-satu/auth"
	"api-satu/respons"
	"api-satu/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service user.Service
	auth    auth.Auth
}

func NewHandler(service user.Service, auth auth.Auth) *handler {
	return &handler{service, auth}
}

func (h *handler) CreateHandler(g *gin.Context) {
	var keyCreateUser user.CreateUser
	err := g.ShouldBindJSON(&keyCreateUser)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}
		failedResponsValue := respons.ResponsValue("Failed Request Create User", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, failedResponsValue)
	} else {
		createService, err := h.service.CreateService(keyCreateUser)
		if err != nil {

			errorMessage := gin.H{"errors": err.Error()}
			failedResponsValue := respons.ResponsValue("Failed Request Create User", http.StatusBadRequest, "Error", errorMessage)
			g.JSON(http.StatusBadRequest, failedResponsValue)
		} else {
			userFormatter := user.UserFormatter(createService, "token token")
			failedResponsValue := respons.ResponsValue("Succes Request Create User", http.StatusOK, "Succes", userFormatter)
			g.JSON(http.StatusOK, failedResponsValue)
		}
	}
}

func (h *handler) AuthUserHandler(g *gin.Context) {
	var authUser user.AuthUser
	err := g.ShouldBindJSON(&authUser)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		errorMessage := gin.H{"errors": errors}
		responsValue := respons.ResponsValue("Failed Login User", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responsValue)
	} else {
		AuthUser, err := h.service.AuthUser(authUser)
		if err != nil {
			responsValue := respons.ResponsValue("Failed Login User", http.StatusBadRequest, "Error", err.Error())
			g.JSON(http.StatusBadRequest, responsValue)
		} else {
			if err != nil {
				responsValue := respons.ResponsValue("Failed Login User", http.StatusBadRequest, "Error", err.Error())
				g.JSON(http.StatusBadRequest, responsValue)
			} else {
				generateToken, err := h.auth.GenerateToken(AuthUser.ID)
				if err != nil {
					responsValue := respons.ResponsValue("Failed Login User", http.StatusBadRequest, "Error", err.Error())
					g.JSON(http.StatusBadRequest, responsValue)
				} else {
					UserFormatter := user.UserFormatter(AuthUser, generateToken)
					responsValue := respons.ResponsValue("Succes Login User", http.StatusOK, "Succes", UserFormatter)
					g.JSON(http.StatusOK, responsValue)
				}
			}

		}
	}
}

func (h *handler) CheckEmailUserHandler(g *gin.Context) {
	var checkEmail user.CheckEmailUser
	err := g.ShouldBindJSON(&checkEmail)
	if err != nil {
		errorMessage := gin.H{"errors": user.ErrorValidation(err)}
		responsValue := respons.ResponsValue("Errors entry email find by email", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responsValue)
	} else {
		checkByEmail, err := h.service.CheckByEmail(checkEmail)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responsValue := respons.ResponsValue("Errors entry email find by email", http.StatusUnprocessableEntity, "Error", errorMessage)
			g.JSON(http.StatusUnprocessableEntity, responsValue)
		} else {
			data := gin.H{"is_available": checkByEmail}
			checkMessage := "email blom ada"
			if checkByEmail {
				checkMessage = "email sudah ada"
			}
			responsValue := respons.ResponsValue(checkMessage, http.StatusOK, "Succes", data)
			g.JSON(http.StatusOK, responsValue)
		}

	}
}

func (h *handler) UpdateImageUserHandler(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": err.Error()}
		responValue := respons.ResponsValue("Failed Upload Image", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responValue)
	} else {
		path := "image/" + file.Filename
		err := g.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"is_uploaded": err.Error()}
			responValue := respons.ResponsValue("Failed Upload Image", http.StatusBadRequest, "Error", errorMessage)
			g.JSON(http.StatusBadRequest, responValue)
		} else {
			currentUser := g.MustGet("current_user").(user.User)
			userID := currentUser.ID
			_, err := h.service.UpdateImage(userID, path)
			if err != nil {
				errorMessage := gin.H{"is_uploaded": err.Error()}
				responValue := respons.ResponsValue("Failed Upload Image", http.StatusBadRequest, "Error", errorMessage)
				g.JSON(http.StatusBadRequest, responValue)
			} else {
				data := gin.H{"is_uploaded": userID}
				responValue := respons.ResponsValue("Success Upload Image", http.StatusOK, "Success", data)
				g.JSON(http.StatusOK, responValue)
			}
		}
	}
}
