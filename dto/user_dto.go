package dto

import "github.com/tzk-chris/Gin-Vue/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// 自定义好返回的数据，这里只返回 用户名、电话
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

//{
//	"code": 200,
//	"data": {
//	"user": {
//	"name": "OuxsjWlViI",
//	"telephone": "12345678900"
//	}
//	}
//}