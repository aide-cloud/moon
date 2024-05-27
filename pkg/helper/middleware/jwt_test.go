package middleware

import (
	"testing"

	"github.com/aide-cloud/moon/pkg/vobj"
)

func TestNewJwtClaims(t *testing.T) {
	token, err := NewJwtClaims(&JwtBaseInfo{
		User:     1,
		Role:     1,
		Team:     1,
		TeamRole: vobj.RoleSuperAdmin,
	}).GetToken()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(token)
}