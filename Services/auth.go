package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

//AuthLogin ... auth login
func AuthLogin(info *structs.AuthLoginRequest) (result structs.AuthLoginResponse, res structs.Response, err error) {
	ad, res, err := CheckInfo(info)
	if err != nil {
		res.Message = err.Error()
		return
	}
	td, res, err := CreateToken(ad)
	if err != nil {
		res.Message = err.Error()
		return
	}
	res, err = CreateAuth(ad, td)
	if err != nil {
		res.Message = err.Error()
		return
	}
	result.AccessToken = td.AccessToken
	result.RefreshToken = td.RefreshToken
	return
}

//AuthInfo ... auth info
func AuthInfo(userID uint64) (ur *models.UserResponse, res structs.Response, err error) {
	ur = &models.UserResponse{}
	var user models.User
	if err = config.DB.Model(&user).Preload("Roles.Permissions").Where("id = ?", userID).First(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	ur.ID = user.ID
	ur.Name = user.Name
	ur.Email = user.Email
	ur.Status = user.Status
	ur.Roles = user.Roles
	return
}

//CheckInfo ... check info
func CheckInfo(info *structs.AuthLoginRequest) (ad *structs.AuthDetails, res structs.Response, err error) {
	var user models.User
	ad = &structs.AuthDetails{}
	if err = config.DB.Where("email = ?", info.Email).First(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	if err = CheckPasswordHash(info.Password, user.Password); err != nil {
		res.Message = err.Error()
		return
	}
	ad.AuthID = user.ID
	ad.AUthEmail = user.Email
	return
}

//CreateToken ... create token
func CreateToken(ad *structs.AuthDetails) (td *structs.TokenDetails, res structs.Response, err error) {
	td = &structs.TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = ad.AuthID
	atClaims["exp"] = td.AccessTokenExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		res.Message = err.Error()
		return
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = ad.AuthID
	rtClaims["exp"] = td.RefreshTokenExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//CreateAuth ... create auth
func CreateAuth(ad *structs.AuthDetails, td *structs.TokenDetails) (res structs.Response, err error) {
	at := time.Unix(td.AccessTokenExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RefreshTokenExpires, 0)
	now := time.Now()

	errAccess := config.Client.Set(td.AccessUUID, ad.AuthID, at.Sub(now)).Err()
	if errAccess != nil {
		res.Message = errAccess.Error()
		return
	}
	errRefresh := config.Client.Set(td.RefreshUUID, ad.AuthID, rt.Sub(now)).Err()
	if errRefresh != nil {
		res.Message = errRefresh.Error()
		return
	}
	return
}

//DeleteAuth ... delete auth
func DeleteAuth(info *structs.AuthLoginRequest) (int64, error) {
	deleted, err := config.Client.Del(info.Email).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

//AuthLogout ... auth logout
func AuthLogout(c *gin.Context) {

}
