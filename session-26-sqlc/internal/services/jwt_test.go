package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJWTService_CreateJWT(t *testing.T) {
	err := os.Setenv("secret", "sometestsecret")
	assert.Nil(t, err)

	jwtsvc := JWTService{}
	jwt, err := jwtsvc.CreateJWT("test username")
	assert.Nil(t, err)
	assert.Greater(t, len(jwt), 0)

}
