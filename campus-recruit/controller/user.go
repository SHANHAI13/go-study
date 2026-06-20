package controller

import (
	"net/http"

	"campus-recruit/middleware"
	"campus-recruit/model"
	"campus-recruit/utils"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var req model.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// TODO: 调用 service 层处理注册逻辑
	// user, err := service.Register(&req)

	utils.Success(c, gin.H{"msg": "注册接口已就绪，service 层待实现"})
}

// 登录
func Login(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// TODO: 1. 查询用户 2. 验证密码 3. 生成 Token
	// 示例返回，实际开发替换为 service 层调用
	token, _ := middleware.GenerateToken(1, req.Username, "student")
	utils.Success(c, model.LoginResp{
		Token: token,
	})
}

// 获取当前用户信息
func GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")
	role, _ := c.Get("role")

	utils.Success(c, gin.H{
		"user_id":  userID,
		"username": username,
		"role":     role,
	})
}
