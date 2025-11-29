package jwt

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

const UserIDKey = "user_id"
const RDS_KEY_TOKEN = "token:"

func GetRDSKeyToken(userID string) string {
	return RDS_KEY_TOKEN + userID
}

var JwtSecret []byte
var JwtExpire = time.Hour * 24

func InitSecret(secret string, expire time.Duration) {
	JwtSecret = []byte(secret)
	JwtExpire = expire
}

type Claims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateToken(uid string) (string, error) {
	claims := &Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JwtExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetUserID(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", false
	}
	return uid, ok
}

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func GetReqID(ctx context.Context) (string, bool) {
	reqID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", false
	}
	return reqID, ok
}

func SetReqID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, UserIDKey, reqID)
}

func ExtractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", errors.New("invalid authorization header")
	}

	return token, nil
}
