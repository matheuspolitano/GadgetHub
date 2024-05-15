package utils

import (
	"math/rand"
	"strings"
	"time"
)

var rdn *rand.Rand

const alphabetic = "abcdefghijklmnopqrstuvwxyz"
const numbers = "012345678"

func init() {
	rsource := rand.NewSource(int64(time.Now().Nanosecond()))
	rdn = rand.New(rsource)
}

// retrieve a random number with min and max
func RandNumber(min int64, max int64) int64 {
	return min + rdn.Int63n(max) + 1
}

// retrieve a random string
func RandString(length int64) string {
	var sBuilder strings.Builder
	alphabeticLength := int64(len(alphabetic))
	for i := int64(0); i < length; i++ {
		sBuilder.WriteByte(alphabetic[RandNumber(0, alphabeticLength)-1])
	}
	return sBuilder.String()
}

func RandStringNumber(length int64) string {
	var sBuilder strings.Builder
	numbersLength := int64(len(numbers))
	for i := int64(0); i < length; i++ {
		sBuilder.WriteByte(numbers[RandNumber(0, numbersLength)-1])
	}
	return sBuilder.String()
}

func RandEmail() string {
	var sBuilder strings.Builder
	sBuilder.WriteString(RandString(30))
	sBuilder.WriteString("@gmail.com")
	return sBuilder.String()
}

func RandPhone() string {
	var sBuilder strings.Builder
	sBuilder.WriteString(RandStringNumber(12))
	return sBuilder.String()
}
