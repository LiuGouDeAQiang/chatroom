package logic

import (
	model "chatroom/app/models"
	"chatroom/app/tools"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id           int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Uid          string `gorm:"column:uid;type:bigint(20)" json:"uid" form:"uid"`
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	Email        string `gorm:"column:email;type:varchar(255)" json:"email"`
	Phone        string `gorm:"column:phone;type:varchar(255)" json:"phone"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func PostLogin(context *gin.Context) {
	var user User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(200, tools.ECode{
			Code:    0,
			Message: err.Error(),
		})
	}

	uid, _ := strconv.ParseInt(user.Uid, 10, 64)
	ret := model.GetUser(uid)
	if ret.Uid < 1 || ret.Password != encryptV1(user.Password) {
		context.JSON(200, tools.UserErr)
		return
	}
	jwt, _ := model.GetJwt(ret.Uid)
	context.JSON(200, tools.ECode{
		Message: "登陆成功",
		Token:   jwt,
	})
	return
}

type CUser struct {
	Name      string `json:"name" form:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password" form:"password"`
	//Password2 string `json:"password_2"`
}

func CreateUser(context *gin.Context) {
	var user CUser
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(200, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	if user.Name == "" {
		context.JSON(http.StatusOK, tools.ParamErr)
		return
	}

	password := len(user.Password)
	if password != 6 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "密码限定六位",
		})
		return
	}
	newUser := model.User{
		Name:        user.Name,
		Password:    encryptV1(user.Password),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	//VerifyCodeHandler(context)
	newUser.Uid = tools.GetUID()
	if err := model.CreateUser(&newUser); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "新用户创建失败！", //这里有风险
		})
		return
	}

	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "创建成功,你的账户是" + strconv.FormatInt(newUser.Uid, 10),
	})

	return
}
func encryptV1(pwd string) string {
	// 创建一个 MD5 哈希对象
	hash := md5.New()
	// 将密码转换为字节数组并计算哈希值
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	// 打印加密后的密码
	// 返回加密后的密码字符串
	return hashString
}
