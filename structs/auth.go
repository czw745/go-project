package structs

//AuthLoginRequest ...
type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//AuthLoginResponse ...
type AuthLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
