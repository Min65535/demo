package decode

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtToken interface {
	Encode(uid, rid uint) (string, error)
	Decode(token string) (uint, uint, error)
}

type jwtToken struct {
	pvtKey *rsa.PrivateKey
	pubKey *rsa.PublicKey
}

func NewJwtToken(pvtKeyPEM string) (JwtToken, error) {
	pvtKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pvtKeyPEM))
	if err != nil {
		return nil, errors.New("parse jwt private key" + err.Error())
	}
	return &jwtToken{pvtKey: pvtKey, pubKey: &pvtKey.PublicKey}, nil
}

func (t *jwtToken) Encode(uid, rid uint) (string, error) {
	if t.pvtKey == nil {
		return "", errors.New("尚未初始化jwt的key")
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"uid": uid,
		"rid": rid,
	}).SignedString(t.pvtKey)
}

func (t *jwtToken) Decode(token string) (uint, uint, error) {
	keyFunc := func(jt *jwt.Token) (interface{}, error) {
		if _, ok := jt.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("jwt token method err:" + token)
		}
		return t.pubKey, nil
	}

	jt, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return 0, 0, err
	}

	claims := jt.Claims.(jwt.MapClaims)
	if _, ok := claims["uid"]; !ok {
		return 0, 0, errors.New("token is not expected")
	}
	if _, ok := claims["rid"]; !ok {
		return 0, 0, errors.New("token is not expected")
	}
	return uint(claims["uid"].(float64)), uint(claims["rid"].(float64)), nil
}

type jwtTokenErr struct {
}

func (j *jwtTokenErr) Encode(_, _ uint) (string, error) {
	return "", errors.New("不存在的校验")
}

func (j *jwtTokenErr) Decode(_ string) (uint, uint, error) {
	return 0, 0, errors.New("不存在的校验")
}

type JwtSets map[string]JwtToken

func (js *JwtSets) Get(key string) JwtToken {
	j, ok := (*js)[key]
	if ok {
		return j
	}
	return &jwtTokenErr{}
}
