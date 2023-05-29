package apis

import (
	model "first_gin-gorm_proj/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录
func Login(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Login()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"data": result.Name,
	})
}

// 注册
func Register(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": id,
	})
}

// 修改密码
func Update(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Update()
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "修改失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

// 删除用户
func Delete(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("username")
	result, err := user.Delete(user.Name)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
