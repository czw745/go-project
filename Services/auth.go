package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
)

//AuthLogin ... auth login
func AuthLogin(info *structs.AuthLoginRequest) (result structs.AuthLoginResponse, res structs.Response, err error) {
	res, err = CheckInfo(info)
	if err != nil {
		res.Message = err.Error()
		return
	}
	result, res, err = CreateToken()
	if err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//CheckInfo ... check info
func CheckInfo(info *structs.AuthLoginRequest) (res structs.Response, err error) {
	var user models.User
	if err = config.DB.Where("email = ?", info.Email).First(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	if err = CheckPasswordHash(info.Password, user.Password); err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//CreateToken ... create token
func CreateToken() (result structs.AuthLoginResponse, res structs.Response, err error) {
	result.AccessToken = "hsadfhdlfjsdfkdshfsflkfhadshflsfsadjkh"
	result.RefreshToken = "uwrweyrewuiryiuewyreyroweyiruweyroyeqwer"
	return
}
