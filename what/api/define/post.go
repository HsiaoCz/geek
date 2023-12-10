package define

// user singup
type UserS struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"repassword"`
	Email      string `json:"email"`
}

// user login
type UserL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
