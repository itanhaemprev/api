package config

import (
	"fmt"
	"os"
)

//SecretAuthKey auth key is a key for auth
var SecretAuthKey string

func init() {
	SecretAuthKey = os.Getenv("API_SECRET_AUTH")
	fmt.Print(SecretAuthKey)
}
