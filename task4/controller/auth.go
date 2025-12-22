package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/models"
	"go_blog/utils"
)

type AuthController struct{}

type RegisterRequest struct {
	UserName string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRequest struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// 用户注册
func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	var existUser models.User
	//检查用户名是否存在
	if err := config.DB.Where("username = ?", req.UserName).First(&existUser).Error; err == nil {
		utils.BadRequest(c, "UserName already exist")
		return
	}
	//检查邮箱是否存在
	if err := config.DB.Where("email = ?", req.Email).First(&existUser).Error; err == nil {
		utils.BadRequest(c, "Email already exist")
		return
	}
	user := models.User{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		utils.InternalServerError(c, "Fail to create user")
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "Fail to generate token")
		return
	}
	utils.Success(c, AuthRequest{
		Token: token,
		User:  user,
	})
}

// 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	var user models.User
	//查找用户
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.Unauthorized(c, "Invalid UserName or Password")
		return
	}
	//验证密码
	if !user.CheckPassword(req.Password) {
		utils.Unauthorized(c, "Invalid UserName or Password")
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "Fail to generate token")
	}
	utils.Success(c, AuthRequest{
		Token: token,
		User:  user,
	})
}

// 获取用户信息
func (ac *AuthController) GetProfile(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		utils.Unauthorized(c, "User is not logged in")
		return
	}
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "User not found")
		return
	}
	utils.Success(c, user)
}
