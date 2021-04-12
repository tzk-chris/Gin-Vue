package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tzk-chris/Gin-Vue/common"
	"github.com/tzk-chris/Gin-Vue/dto"
	"github.com/tzk-chris/Gin-Vue/model"
	"github.com/tzk-chris/Gin-Vue/response"
	"github.com/tzk-chris/Gin-Vue/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 { // 如果名称没有传， 就一个随机10位的字符串
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	// 创建用户，密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.FailInternal(ctx, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 返回结果
	response.Success(ctx, nil, "注册成功")

}

func Login(ctx *gin.Context)  {
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	log.Println(telephone, password)
	// 判断手机号是否存在 database
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在，请注册")
		return
	}
	//  判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}
	// 发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.FailInternal(ctx, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user") // 从上下文中取到信息

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephont string) bool {
	var user model.User
	db.Where("telephone = ?", telephont).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}