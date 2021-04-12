package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tzk-chris/Gin-Vue/common"
)

//func Hello(ctx *gin.Context) {
//	ctx.String(200, "Hello gin")
//}

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {

	db := common.InitDB()
	defer db.Close()

	router := gin.Default() // 比gin.New()多了logger和recovery
	//view := gin.New()
	//Router(router)
	router = CollectRoute(router)
	router.LoadHTMLGlob("template/**/*") // 定位二级目录，三级目录：template/**/**/*

	panic(router.Run(":9090")) // fixed port

	//router.GET("/", Hello)
	//router.GET("/user", view.User)
	//router.GET("/user/:id", view.GetFrontendParams) // :冒号占位符
	////router.GET("/user2/*id", view.GetFrontendParams2) // *星号占位符 无参数也访问成功 但不推荐
	//router.GET("/query", view.GetFrontendQuery)
	//router.GET("/queryarr", view.GetFrontendQueryArr)
	//router.GET("/querymap", view.GetFrontendQueryMap)
	//router.GET("/touseradd", view.UserAdd)
	//router.POST("/douseradd", view.PostUserAdd)

}



