package util

import (
	"context"
	"testing"
	"time"
	"webook/internal/interface/web"
	"webook/internal/repository/dao"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const tokenExpiration = time.Minute * 10

func CreateUserAndGetToken(t *testing.T, db *gorm.DB, userAgent string) (dao.User, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := dao.User{
		Email:    "test@example.com",
		Password: "abc123$$$",
	}
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		assert.FailNow(t, "failed to create user", err)
		return dao.User{}, "", err
	}

	token, err := signJwtToken(user.ID, userAgent, tokenExpiration)
	if err != nil {
		assert.FailNow(t, "failed to sign jwt token", err)
		return dao.User{}, "", err
	}

	return user, token, nil
}

func signJwtToken(userId int64, userAgent string, expiration time.Duration) (string, error) {
	uc := web.UserClaims{
		Uid:       userId,
		UserAgent: userAgent,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
	return token.SignedString(web.JWTKey)
}
