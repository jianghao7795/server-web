package request

type LoginForm struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" from:"password"`
}
