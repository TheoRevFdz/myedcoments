package commons

import (
	"CursoGolang/myedcoments/models"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a PublicKey")
	}
}

// GenerateJWT (genera el token para el cliente)
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Theo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}

	return result
}
