package util

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"os"
	"time"

	jwe "github.com/square/go-jose"
)

var (
	privateKey *rsa.PrivateKey
	encrypter  jwe.Encrypter
)

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	strPrivateKey := ExportRsaPrivateKeyAsPemStr(privateKey)
	fprivate, err := os.Create("key.private")
	if err != nil {
		panic(err)
	}
	defer fprivate.Close()

	_, err = fprivate.WriteString(strPrivateKey)
	if err != nil {
		panic(err)
	}

	strPublicKey,err := ExportRsaPublicKeyAsPemStr(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	fPublic, err := os.Create("key.public")
	if err != nil {
		panic(err)
	}
	defer fPublic.Close()

	_, err = fPublic.WriteString(strPublicKey)
	if err != nil {
		panic(err)
	}

	encrypter, err = jwe.NewEncrypter(jwe.A128GCM, jwe.Recipient{Algorithm: jwe.RSA_OAEP, Key: &privateKey.PublicKey}, nil)
	if err != nil {
		panic(err)
	}
}

type CustomClaims struct {
	IDUser      string `json:"user_id"`
	Name        string `json:"user_name"`
	AccessLevel int    `json:"access_level"`
	Exp         int64  `json:"exp"`
}

func NewClaim(id, name string, accessLevel int) *CustomClaims {
	return &CustomClaims{
		IDUser:      id,
		Name:        name,
		AccessLevel: accessLevel,
		Exp:         time.Now().Add(time.Hour * 2).Unix(),
	}
}

func CreateJWE(c *CustomClaims) (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	object, err := encrypter.Encrypt(data)
	if err != nil {
		return "", err
	}

	return object.CompactSerialize()
}

func GetJWE(jweValue string) (*CustomClaims, error) {

	object, err := jwe.ParseEncrypted(jweValue)
	if err != nil {
		return nil, err
	}

	payload, err := object.Decrypt(privateKey)
	if err != nil {
		return nil, err
	}

	var c CustomClaims
	err = json.Unmarshal(payload, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
