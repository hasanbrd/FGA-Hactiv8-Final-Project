package photo

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo.UserID = userID

	err := ctx.ShouldBindJSON(&Photo)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&Photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createPhoto{
		photo: photo{
			ID:       Photo.ID,
			Title:    Photo.Title,
			Caption:  Photo.Caption,
			PhotoUrl: Photo.PhotoUrl,
			UserID:   Photo.UserID,
		},
		CreatedAt: time.Now(),
	})
}

func GetAllPhotos(ctx *gin.Context) {
	db := database.GetDB()
	photos := []models.Photo{}

	err := db.Find(&photos).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("Email", "Username")
	}).Find(&photos)

	ctx.JSON(http.StatusOK, photos)
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo.UserID = userID

	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	Photo.ID = uint(photoId)

	err := ctx.ShouldBindJSON(&Photo)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatePhoto{
		photo: photo{
			ID:       Photo.ID,
			Title:    Photo.Title,
			Caption:  Photo.Caption,
			PhotoUrl: Photo.PhotoUrl,
			UserID:   Photo.UserID,
		},
		UpdatedAt: Photo.UpdatedAt,
	})
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}

	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	photo.ID = uint(photoId)

	err := db.Where("id = ?", photoId).Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
