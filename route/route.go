package route

import (
	"blog/common"
	"blog/service/domain"
	postservice "blog/service/post"
	userservice "blog/service/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

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
	// 解析Token
	token1, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 检查Token的签名方法是否是我们所期望的算法，这里我们期望的是HS256算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于验证Token的签名
		getString := viper.GetString("jwt.password")
		return []byte(getString), nil
	})
	if err != nil || !token1.Valid {
		log.Error("auth token is invalid")
	}

	claims := token1.Claims
	if claims != nil && claims.Valid() == nil {
		mapClaims := claims.(jwt.MapClaims)
		if mapClaims != nil {
			userId := mapClaims["id"].(float64)
			c.Header("UserId", strconv.FormatInt(int64(userId), 10))
			c.Set("UserId", strconv.FormatInt(int64(userId), 10))
		}
	}

	c.Next()
}

// 用户注册
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

// 用户登录
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

// 发布文章
func AddPost(c *gin.Context) {
	request := domain.PostRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		log.Error("AddPost gin bind error", err)
		common.Fail(c, err)
		return
	}

	userId := c.GetHeader("UserId")
	if userId == "" {
		userId = c.GetString("UserId")
	}
	request.UserId, _ = strconv.ParseInt(userId, 10, 64)
	err = postservice.Add(request)
	if err != nil {
		log.Error("AddPost error", err)
		common.Fail(c, err)
		return
	}

	common.Success(c, "")
}

// 查看文章
func ViewPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 文章列表
func QueryPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 更新文章
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 删除文章
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 发布评论
func PostComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 文章的所有评论
func AllComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
