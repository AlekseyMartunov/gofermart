package tokenmanager

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTokenController(t *testing.T) {
	tk := New(time.Hour, []byte("secretKey"))

	userID := "some info"

	token, err := tk.CreateToken(userID)
	assert.NoError(t, err, "Ошибка создания токена")

	id, err := tk.GetUserUUID(token)

	assert.NoError(t, err, "Ошибка при парсинге токена")
	assert.Equal(t, userID, id, "Неправильное значение userID, полученное из токена")

	wrongToken := token[:len(token)-5] + "wrong info"

	wrongID, err := tk.GetUserUUID(wrongToken)

	assert.Equal(t, err, InvalidToken)
	assert.Equal(t, "", wrongID)

}
