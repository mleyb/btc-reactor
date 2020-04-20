package env

import (
	"os"
)

func PhoneNumber() string {
	return os.Getenv("PHONE_NUMBER")
}
