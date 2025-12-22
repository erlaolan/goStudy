package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/models"
	"go_blog/utils"
	"strconv"
)

type CommentController struct{}

type CreateCommentRequest struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid Post ID")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "User is not logged in")
		return
	}
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
		return
	}

	var comment = models.Comment{
		UserID:  userID.(uint),
		PostID:  uint(postID),
		Content: req.Content,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		utils.InternalServerError(c, "Failed to create comment")
		return
	}

	config.DB.Preload("User").Preload("Post").First(&comment, comment.ID)
	utils.Success(c, comment)

}

func (cc *CommentController) GetCommentList(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid Post ID")
		return
	}
	//验证文章是否存在
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	var comments []models.Comment
	if err := config.DB.Preload("User").
		Where("post_id = ?", postID).
		Limit(pageSize).
		Offset(offset).
		Find(&comments).Error; err != nil {
		utils.InternalServerError(c, "Failed to get comment list")
		return
	}
	var total int64
	config.DB.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&total)

	utils.Success(c, gin.H{
		"total":     total,
		"comments":  comments,
		"page":      page,
		"page_size": pageSize,
	})
}
