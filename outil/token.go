package outil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	keyDefault = "pfomoy#274/vp,"
)

var (
	methodDefault = jwt.SigningMethodHS256
)

func DefaultMethod() jwt.SigningMethod {
	return methodDefault
}

func TokenGenWithClaims(claims map[string]any, method jwt.SigningMethod, key any) (string, error) {
	if key == nil {
		key = keyDefault
	}
	return jwt.NewWithClaims(method, jwt.MapClaims(claims)).SignedString(key)
}

func TokenParse(tokenStr string, key any) (claims map[string]any, err error) {
	if key == nil {
		key = keyDefault
	}
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return
	}

	if token.Valid {
		err = jwt.ErrInvalidKey
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析结果无法转换为jwt.MapClaims")
		return
	}

	return claims, nil
}

func TokenParseField(tokenStr string, key any, field string) (value any, err error) {
	claims, err := TokenParse(tokenStr, key)
	if err != nil {
		return
	}

	value, ok := claims[field]
	if !ok {
		err = errors.New("token中没有该字段")
	}
	return
}
