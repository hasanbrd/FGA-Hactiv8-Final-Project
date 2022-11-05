package user

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	user := models.User{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, registerResponse{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	})
}

func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	user := models.User{}
	password := ""

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		fmt.Println(err)
		fmt.Println("error disini")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email or Pasword",
		})
		return
	}

	comparesPass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparesPass {
		fmt.Println("error dicompare")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email or Pasword",
		})
		return
	}

	jwt := helpers.GenarateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, loginRespon{
		Token: jwt,
	})
}

func UpdateUser(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	user := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	user.ID = userID

	err := db.Model(&user).Where("id = ?", userID).Updates(models.User{
		Email:    user.Email,
		Username: user.Username,
	}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Where("id = ?", userID).First(&user)

	ctx.JSON(http.StatusOK, updateResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

func DeleteUser(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.User{}

	userID := uint(userData["id"].(float64))
	err := db.Delete(user, uint(userID)).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
