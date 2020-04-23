package env

import (
	"os"
	"strconv"
)

func PhoneNumber() string {
	return os.Getenv("PHONE_NUMBER")
}

func Limit() float64 {
	f, _ := strconv.ParseFloat(os.Getenv("LIMIT"), 32)
	return f
}
