package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/models"
	"go_blog/utils"
	"strconv"
)

type PostController struct{}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
		return
	}
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}
	if err := config.DB.Create(&post).Error; err != nil {
		utils.InternalServerError(c, "Fail to create Post")
		return
	}
	// 预加载用户信息
	config.DB.Preload("User").First(&post, post.ID)

	utils.Success(c, post)
}

// 获取所有文章列表
func (pc *PostController) GetPostList(c *gin.Context) {
	var posts []models.Post

	//分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	// 查询文章列表 预加载用户信息
	if err := config.DB.Preload("User").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error; err != nil {
		utils.InternalServerError(c, "Fail to get Post list")
		return
	}
	//获取总数
	var total int64

	config.DB.Model(&models.Post{}).Count(&total)

	utils.Success(c, gin.H{
		"total":     total,
		"posts":     posts,
		"page":      page,
		"page_size": pageSize,
	})
}

// 获取单个文章
func (pc *PostController) GetPost(c *gin.Context) {

	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid Post ID")
		return
	}
	var post models.Post
	if err := config.DB.Preload("User").Preload("Comments.User").First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
	}
	utils.Success(c, post)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid Post ID")
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
	}
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
		return
	}
	// 检查是否文章作者
	if post.UserID != userID.(uint) {
		utils.Forbidden(c, "You are not allowed to update this Post")
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	if err := config.DB.Save(&post).Error; err != nil {
		utils.InternalServerError(c, "Fail to update Post")
	}

	config.DB.Preload("User").First(&post, post.ID)
	utils.Success(c, post)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid Post ID")
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
	}
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
	}
	if post.UserID != userID.(uint) {
		utils.Forbidden(c, "You are not allowed to delete this Post")
		return
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		utils.InternalServerError(c, "Fail to delete Post")
	}
	utils.Success(c, gin.H{"message": "Post deleted"})
}
