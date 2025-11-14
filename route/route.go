package route

import (
	"blog/common"
	"blog/service/domain"
	postservice "blog/service/post"
	userservice "blog/service/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InitRoute 初始化路由
func InitRoute(router *gin.Engine) {

	group1 := router.Group("/user")
	group1.POST("/register", Register)
	group1.POST("/login", Login)

	group2 := router.Group("/post", auth)
	group2.POST("/add", AddPost)
	group2.POST("/view", ViewPost)
	group2.POST("/query", QueryPost)
	group2.POST("/update", UpdatePost)
	group2.POST("/delete", DeletePost)

	group3 := router.Group("/comment", auth)
	group3.POST("/post", PostComment)
	group3.POST("/all", AllComment)

}

// 用户权限验证
func auth(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		log.Error("auth token is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token is empty",
		})
		return
	}

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	parseToken, err := common.ParseToken(token)
	if err != nil {
		log.Error("auth token is invalid")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token is invalid",
		})
	}
	if parseToken != nil {
		c.Set("UserId", parseToken.Id)
	}

	c.Next()
}

// Register 用户注册
func Register(c *gin.Context) {

	request := domain.RegisterRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("Register gin bind error", err)
		common.Fail(c, err)
		return
	}

	userData, err := userservice.Register(request)
	if err != nil {
		log.Error("Register Register error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, userData)
}

// Login 用户登录
func Login(c *gin.Context) {

	request := domain.LoginRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("Register gin bind error", err)
		common.Fail(c, err)
		return
	}

	userData, err := userservice.Login(request)
	if err != nil {
		log.Error("Register Register error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, userData)
}

// AddPost 发布文章
func AddPost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("AddPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	request.UserId = c.GetInt64("UserId")
	err = postservice.Add(request)
	if err != nil {
		log.Error("AddPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}

// ViewPost 查看文章
func ViewPost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("ViewPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	if request.Id == 0 {
		log.Error("id is blank ")
		err1 := errors.New("id is blank ")
		common.Fail(c, err1)
		return
	}

	data, err := postservice.View(request)
	if err != nil {
		log.Error("ViewPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, data)
}

// QueryPost 文章列表
func QueryPost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("QueryPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	data, err := postservice.Query(request)
	if err != nil {
		log.Error("QueryPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, data)
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("UpdatePost gin bind error", err)
		common.Fail(c, err)
		return
	}

	request.UserId = c.GetInt64("UserId")
	err = postservice.Update(request)
	if err != nil {
		log.Error("UpdatePost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("DeletePost gin bind error", err)
		common.Fail(c, err)
		return
	}

	if request.Id == 0 {
		log.Error("id is blank ")
		err1 := errors.New("id is blank ")
		common.Fail(c, err1)
		return
	}

	err = postservice.Delete(request)
	if err != nil {
		log.Error("DeletePost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}

// PostComment 发布评论
func PostComment(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("AddPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	request.UserId = c.GetInt64("UserId")
	err = postservice.Add(request)
	if err != nil {
		log.Error("AddPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}

// AllComment 文章的所有评论
func AllComment(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("AddPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	request.UserId = c.GetInt64("UserId")
	err = postservice.Add(request)
	if err != nil {
		log.Error("AddPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}
