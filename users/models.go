package users

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Token string `json:"token"`
}
