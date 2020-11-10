package structs

//AuthLoginRequest ... auth login request
type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//AuthDetails ... auth details
type AuthDetails struct {
	AuthID    uint
	AUthEmail string
}

//AccessDetails ... access details
type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}

//AuthLoginResponse ... auth login response
type AuthLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//TokenDetails ... token details
type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}
