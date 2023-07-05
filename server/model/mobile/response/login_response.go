package response

import "server/model/mobile"

type LoginResponse struct {
	User      mobile.MobileUser `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
