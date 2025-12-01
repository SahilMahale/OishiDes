package server

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"strings"
	"time"

	"github.com/SahilMahale/OishiDes/oishi-server/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func getAuthToken(c *fiber.Ctx) string {
	return strings.Split(c.Get("Authorization"), "Bearer ")[1]
}

func checkIfAdmin(role string) bool {
	return role == string(constants.Admin)
}
func readPrivateKeyFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	decodedKey, _ := pem.Decode(buffer)

	pk, errP := x509.ParsePKCS8PrivateKey(decodedKey.Bytes)
	if errP != nil {
		return errP
	}
	pKey, ok := pk.(*rsa.PrivateKey)
	if !ok {
		panic("PrivateKey type error")
	}
	privateKey = pKey
	return nil
}

func readPublicKeyFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	decodedKey, _ := pem.Decode(buffer)

	pk, errP := x509.ParsePKIXPublicKey(decodedKey.Bytes)
	if errP != nil {
		return errP
	}
	pubKey, ok := pk.(*rsa.PublicKey)
	if !ok {
		panic("PrivateKey type error")
	}
	publicKey = pubKey
	return nil
}

func makeTokenWithClaims(checkIfAdmin bool, username string) (token string, err error) {
	// create claims for user
	claims := MyCustomClaims{
		Name: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	if checkIfAdmin {
		claims.Type = string(constants.Admin)
	} else {
		claims.Type = string(constants.User)
	}
	// create token for the claims
	ctoken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	token, errp := ctoken.SignedString(privateKey)
	if errp != nil {
		log.Error(errp)
		return "", errp
	}
	return token, nil
}

func getClaimsForThisCall(token string) (claims *MyCustomClaims, err error) {
	claims = &MyCustomClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
