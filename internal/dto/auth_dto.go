package dto

type GetJWTIput struct {
	Username string `json:"username,omitempty" :"username"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
