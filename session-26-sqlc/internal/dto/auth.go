package dto

type SignupRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
