package request

import (
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

// GetAudience implements jwt.Claims.
func (c CustomClaims) GetAudience() (jwt.ClaimStrings, error) {
	panic("unimplemented")
}

// GetExpirationTime implements jwt.Claims.
func (c CustomClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetIssuedAt implements jwt.Claims.
func (c CustomClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetIssuer implements jwt.Claims.
func (c CustomClaims) GetIssuer() (string, error) {
	panic("unimplemented")
}

// GetNotBefore implements jwt.Claims.
func (c CustomClaims) GetNotBefore() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetSubject implements jwt.Claims.
func (c CustomClaims) GetSubject() (string, error) {
	panic("unimplemented")
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}
