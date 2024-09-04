package response

type AuthResponse struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func NewAuthResponse(token string) AuthResponse {
	return AuthResponse{
		Token:  token,
		Prefix: "Bearer",
	}
}
