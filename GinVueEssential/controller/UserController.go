package controller

import (
	"net/http"

	"herrkung.com/GinVueEssential/response"

	"herrkung.com/GinVueEssential/dto"

	"golang.org/x/crypto/bcrypt"

	"herrkung.com/GinVueEssential/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"herrkung.com/GinVueEssential/model"
)

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	//c.JSON(http.StatusOK, gin.H{"code": 200, "user": dto.ToUserDto(user.(model.User))})
	response.Response(c, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//1 get parameter
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//2 data auth
	if len(telephone) != 11 {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "telephone length must be 11"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "telephone length must be 11")
		return
	}
	if len(password) < 6 {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "password length must be greater than 5"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password length must be greater than 5")
		return
	}
	//3 user is exists?
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "user is NOT existed"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "user is NOT existed")
		return
	}
	//4 password is correct?
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 400, "Msg": "password is NOT correct"})
		response.Response(c, http.StatusInternalServerError, 400, nil, "password is NOT correct")
		return
	}
	//5 release token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "release token fail")
		return
	}
	//6 return login successful result
	//c.JSON(200, gin.H{"token": token, "Msg": "login Successful"})
	response.Success(c, gin.H{"token": token}, "login Successful")

}

func Register(c *gin.Context) {
	DB := common.GetDB()
	//1 get parameter
	//name := c.PostForm("name")
	//telephone := c.PostForm("telephone")
	//password := c.PostForm("password")
	var requestUser = model.User{}
	_ = c.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//2 data auth
	if len(telephone) != 11 {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "telephone length must be 11"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "telephone length must be 11")
		return
	}
	if len(password) < 6 {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "password length must be greater than 5"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password length must be greater than 5")
		return
	}
	if len(name) == 0 {
		name = "DefaultUserName"
	}
	//3 user is exists?
	if isTelephoneExist(DB, telephone) {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "Msg": "user is already existed"})
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "user is already existed")
		return
	}
	//4 create user on database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "Msg": "encrypt user password ERROR"})
		response.Response(c, http.StatusInternalServerError, 500, nil, "encrypt user password ERROR")
		return
	}
	newUser := &model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	DB.Create(&newUser) //create a user in database
	//5 return register successful result
	//c.JSON(200, gin.H{"Msg": "Register Successful"})
	response.Success(c, nil, "Register Successful")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	//按照telephone索引用户，找到后将其信息填充到user中
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
