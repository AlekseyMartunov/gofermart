package hashencoder

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var symbolsRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const testLen = 1000

// len of user login
var max = 50
var min = 5

func generateRandomString(size int) string {
	output := make([]rune, size)
	for i := range output {
		output[i] = symbolsRunes[rand.Intn(len(symbolsRunes))]
	}
	return string(output)
}

func TestHashingManager_Encode_Different_Strings(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	hm := New()
	hashMap := make(map[string]string)

	for i := 0; i < testLen; i++ {
		text := generateRandomString(rand.Intn(max-min) + min)
		hash := hm.Encode(text)

		_, ok := hashMap[hash]

		assert.False(t, ok, "collision or incorrect work of the hash manager")
		hashMap[hash] = text
	}
}

func TestHashingManager_Encode_SameStrings(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	hm := New()
	randomString := generateRandomString(rand.Intn(max-min) + min)

	hash := hm.Encode(randomString)

	for i := 0; i < testLen; i++ {
		assert.Equal(t, hash, hm.Encode(randomString),
			"Hash manager return different string for same input")
	}
}
