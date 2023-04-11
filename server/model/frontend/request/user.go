package request

type LoginForm struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" from:"password"`
}

type RegisterUser struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" from:"password"`
	Introduction string `json:"introduction" form:"introduction"`
	Content      string `json:"content" form:"content"`
}
