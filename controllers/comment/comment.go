package comment

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	comment := models.Comment{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comment.UserID = userID

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Model(&photo).Where("id = ?", comment.PhotoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo ID not found",
		})
		return
	}

	err = db.Debug().Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createComment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	})
}

func GetAllComments(ctx *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}

	err := db.Find(&comments).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Preload("Photo").Find(&comments)

	ctx.JSON(http.StatusOK, comments)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comment.UserID = userID

	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	comment.ID = uint(commentId)

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Model(&comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: comment.Message,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Model(&comment).Where("id = ?", comment.ID).First(&comment)

	ctx.JSON(http.StatusOK, updateComment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}

	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	comment.ID = uint(commentId)

	err := db.Where("id = ?", comment.ID).Delete(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
