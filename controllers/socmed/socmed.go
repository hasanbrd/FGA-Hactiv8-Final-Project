package socmed

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia.UserID = userID

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Debug().Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	})
}

func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedias := []models.SocialMedia{}

	err := db.Find(&socialMedias).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username")
	}).Find(&socialMedias)

	ctx.JSON(http.StatusOK, socialMedias)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia.UserID = userID

	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	socialMedia.ID = uint(socialMediaId)

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Model(&socialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updateSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia.ID = uint(socialMediaId)

	err := db.Where("id = ?", socialMedia.ID).Delete(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
