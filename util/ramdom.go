package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Constants for different generators
const (
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
	GOLANG     = "go"
	PYTHON     = "py"
	JAVASCRIPT = "js"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(max int32) int32 {
	return rand.Int31n(max)
}

// RandomFloat generates a random float between 0.0 and 1.0
func RandomFloat() float64 {
	return rand.Float64()
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomClassName generates a random owner name
func RandomClassName() string {
	return RandomString(18)
}

// RandomTestCaseName generates a random owner name
func RandomTestCaseName() string {
	return RandomString(24)
}

// RandomLineNumber generates a random amount of money
func RandomLineNumber() int32 {
	return RandomInt(1000)
}

// RandomDuration generates a random currency code
func RandomDuration() float64 {
	return RandomFloat()
}

// RandomFileName generates a random filename
func RandomFileName() string {
	extensions := []string{GOLANG, JAVASCRIPT, PYTHON}
	n := len(extensions)
	return fmt.Sprintf("%s.%s", RandomString(6), extensions[rand.Intn(n)])
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}