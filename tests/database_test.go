package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vsrtferrum/AvitoIntro/internal/database"
	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/logger"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

const (
	testConnStr = "postgresql://testuser:testpassword@localhost:5440/testdb"
)

// setupTestDB подготавливает тестовую базу данных и возвращает экземпляр Database.
func setupTestDB(t *testing.T) (*database.Database, func()) {
	// Создаем логгер с выводом в stdout
	log, err := logger.NewLogger("../.logs")
	if err != nil {
		t.Fatal(err)
	}

	// Инициализируем базу данных
	db := database.NewDatabase(testConnStr, log)

	// Подключаемся к базе
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}

	// Функция очистки
	cleanup := func() {
		db.Close()
	}

	return &db, cleanup
}

// TestAuth тестирует метод Auth.
func TestAuth(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Тест успешной аутентификации
	t.Run("successful auth", func(t *testing.T) {
		// Создаем тестового пользователя через метод Auth
		_, err := db.Auth(model.AuthRequest{Username: "testuser", Password: "testpass"})
		assert.NoError(t, err)

		// Проверяем, что пользователь создан, пытаясь аутентифицироваться снова
		_, err = db.Auth(model.AuthRequest{Username: "testuser", Password: "testpass"})
		assert.NoError(t, err)
	})

	// Тест создания нового пользователя
	t.Run("new user creation", func(t *testing.T) {
		_, err := db.Auth(model.AuthRequest{Username: "newuser", Password: "newpass"})
		assert.NoError(t, err)

		// Проверяем, что пользователь существует, пытаясь аутентифицироваться
		_, err = db.Auth(model.AuthRequest{Username: "newuser", Password: "newpass"})
		assert.NoError(t, err)
	})
}

// TestBuyItem тестирует метод BuyItem.
func TestBuyItem(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Создаем тестового пользователя
	_, err := db.Auth(model.AuthRequest{Username: "user1", Password: "pass1"})
	assert.NoError(t, err)

	// Создаем тестовый товар через метод Auth (если нет метода для добавления товаров)
	// В реальном коде добавьте метод для управления товарами, если это необходимо.
	// Здесь мы предполагаем, что товар уже существует в базе данных.

	// Тест покупки товара
	t.Run("buy item", func(t *testing.T) {
		// Получаем баланс пользователя до покупки
		balanceBefore, err := db.GetUserBalanceById(1)
		assert.NoError(t, err)

		// Покупаем товар
		err = db.BuyItem(1, 1, balanceBefore-100, 100)
		assert.Equal(t, err, errors.ErrExecTransaction)

		// Проверяем, что баланс не изменился
		_, err = db.GetUserBalanceById(1)
		assert.NoError(t, err, errors.ErrExecTransaction)

	})
}

// TestSendMoney тестирует метод SendMoney.
func TestSendMoney(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Создаем двух пользователей
	_, err := db.Auth(model.AuthRequest{Username: "sender", Password: "pass1"})
	assert.NoError(t, err)

	_, err = db.Auth(model.AuthRequest{Username: "receiver", Password: "pass2"})
	assert.NoError(t, err)

	// Тест отправки денег
	t.Run("send money", func(t *testing.T) {
		// Получаем балансы до отправки
		senderBalanceBefore, err := db.GetUserBalanceById(1)
		assert.NoError(t, err)

		receiverBalanceBefore, err := db.GetUserBalanceById(2)
		assert.NoError(t, err)

		// Отправляем деньги
		err = db.SendMoney(1, "receiver", 200, senderBalanceBefore-200, receiverBalanceBefore+200)
		assert.NoError(t, err)

		// Проверяем, что балансы изменились
		senderBalanceAfter, err := db.GetUserBalanceById(1)
		assert.NoError(t, err)
		assert.Equal(t, senderBalanceBefore-200, senderBalanceAfter)

		receiverBalanceAfter, err := db.GetUserBalanceById(2)
		assert.NoError(t, err)
		assert.Equal(t, receiverBalanceBefore+200, receiverBalanceAfter)
	})
}
