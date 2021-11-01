package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

var issuer = "wiki-global-manage"

//
type CustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

//Token jwt服务
type Token struct {
	rwlock     sync.RWMutex
	privateKey []byte
}

//获取
func (srv *Token) Get() []byte {
	return srv.get()
}

//获取
func (srv *Token) get() []byte {
	srv.rwlock.RLock()
	defer srv.rwlock.RUnlock()
	srv.privateKey = []byte("wikiglobal-2020")
	return srv.privateKey
}

//
func (srv *Token) put(newKey []byte) {
	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()

	srv.privateKey = newKey
}

//初始化
func (srv *Token) InitToken(privateKey string) {
	srv.put([]byte(privateKey))
}

//Decode解码
func (srv *Token) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.get(), nil
	})

	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode 将 User 用户信息加密为 JWT 字符串
// expireTime := time.Now().Add(time.Hour * 24 * 3).Unix() 三天后过期
func (srv *Token) Encode(userId, expireTime int64) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(srv.get())
}
