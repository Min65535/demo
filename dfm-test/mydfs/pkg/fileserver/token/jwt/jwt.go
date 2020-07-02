package jwt

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"time"
)

type Jwt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJwt() *Jwt {
	skPEM := devSK
	pkPEM := devPK
	if qyenv.GetUseDocker() == 2 {
		skPEM = privatekeyPEM
		pkPEM = publickeyPEM
	}

	sk, err := jwt.ParseRSAPrivateKeyFromPEM(skPEM)
	if err != nil {
		panic("jwt private key PEM error")
		return nil
	}
	pk, err := jwt.ParseRSAPublicKeyFromPEM(pkPEM)
	if err != nil {
		panic("jwt public key PEM error")
		return nil
	}
	return &Jwt{PrivateKey: sk, PublicKey: pk}
}

// generate a token from userID
func (t *Jwt) GenerateToken(userID uint, expiredAt time.Duration) (accessToken string, err error) {
	exp := time.Now().Add(expiredAt)
	// jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"exp": exp.Unix(), "userID": userID})
	// sign the jwt token
	accessToken, err = token.SignedString(t.PrivateKey)
	if err != nil {
		// todo: log error
	}
	return
}

// get userID from token
func (t *Jwt) DecodeToken(token string) (userID uint, err error) {
	// get map claims
	claims, err := t.claimsFromToken(token)
	if err != nil {
		return
	}
	if _, ok := claims["userID"]; !ok {
		err = errors.New("token is not expected")
		return
	}
	userID = uint(claims["userID"].(float64))
	return userID, err
}

// get map from token
func (t *Jwt) claimsFromToken(tokenString string) (jwt.MapClaims, error) {
	// parse token
	jwtToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			// todo: log error
			return
		}
		return t.PublicKey, nil
	})

	// get claims
	var claims jwt.MapClaims
	if jwtToken == nil || jwtToken.Claims == nil {
		return claims, errors.New("jwtToken error")
	}

	claims = jwtToken.Claims.(jwt.MapClaims)
	return claims, err
}

// todo: get privatekey & publickey from apollo
var (
	privatekeyPEM = []byte(`-----BEGIN RSA PRIVATE KEY-----
xxxxxxx
-----END RSA PRIVATE KEY-----`)

	publickeyPEM = []byte(`-----BEGIN PUBLIC KEY-----
xxxxxxx
-----END PUBLIC KEY-----`)

	devSK = []byte(`-----BEGIN RSA PRIVATE KEY-----
xxxxxxx
-----END RSA PRIVATE KEY-----`)
	devPK = []byte(`-----BEGIN PUBLIC KEY-----
xxxxxxx
-----END PUBLIC KEY-----`)
)
