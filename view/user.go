package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id       int
	UserName string
	Age      int
	Addr     string
}

func User(ctx *gin.Context)  {
	// 结构体渲染
	user_info := UserInfo{1, "chrix", 18, "csc"}
	ctx.HTML(200, "user/user.html", user_info)

}

func GetFrontendParams(ctx *gin.Context)  {
	// 获取前端参数 test
	user_id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello, %s", user_id)
}

func GetFrontendParams2(ctx *gin.Context)  {
	// 获取前端参数 test  星号占位符
	user_id2 := ctx.Param("id")
	ctx.String(http.StatusOK, "hello, %s", user_id2)
}

func GetFrontendQuery(ctx *gin.Context)  {
	// 获取前端参数名 test http://127.0.0.1:9090/user3?id=123
	user_id := ctx.Query("id")
	ctx.String(http.StatusOK, "chrix, %s", user_id)
}

func GetFrontendQueryArr(ctx *gin.Context)  {
	// 获取前端参数名 test http://127.0.0.1:9090/queryarr?id=123,124,125
	ids := ctx.QueryArray("id")
	ctx.String(http.StatusOK, "chrix, %s", ids)
}

func GetFrontendQueryMap(ctx *gin.Context)  {
	// 获取前端参数名 test http://127.0.0.1:9090/querymap?user[name]=chrix&user[age]=18
	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK, "ur name, %s", user)
}

// 获取post请求数据
func UserAdd(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "user/user_add.html", nil)
}

func PostUserAdd(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)

	ctx.String(http.StatusOK, "添加成功")
}