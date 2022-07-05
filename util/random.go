package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//randomSring generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}

	return sb.String()
}

//function returns random owner
func RandomOwner() string {
	return RandomString(6)
}

//function returns random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

//function returns random currency type
func RandomCurrency() string {
	currencies := []string{"NGN", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
