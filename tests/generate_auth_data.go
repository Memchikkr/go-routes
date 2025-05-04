package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Модели данных
type TelegramAuthData struct {
	AuthDate int    `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	ID        int    `json:"id"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url,omitempty"`
}

type EmailAuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	telegramData := generateTelegramAuthData("AAAA:123123123123")
	fmt.Println("Telegram Auth Data:")
	printJSON(telegramData)
}

// Генерация валидных данных Telegram WebApp
func generateTelegramAuthData(botToken string) TelegramAuthData {
	authDate := int(time.Now().Unix())
	id := rand.Intn(900000) + 100000 // 6-значный ID
	firstName := randomString(5) + " " + randomString(8)
	username := "user_" + randomString(6)

	data := TelegramAuthData{
		AuthDate: authDate,
		FirstName: firstName,
		ID:        id,
		Username:  username,
		PhotoURL:  fmt.Sprintf("https://t.me/i/userpic/%s.jpg", randomString(10)),
	}

	// Генерация хеша
	data.Hash = calculateTelegramHash(data, botToken)

	return data
}

// Расчет хеша для Telegram данных
func calculateTelegramHash(data TelegramAuthData, botToken string) string {
	secretKey := sha256.Sum256([]byte(botToken))
	var fields []string

	// Добавляем все поля кроме hash
	fields = append(fields, fmt.Sprintf("auth_date=%d", data.AuthDate))
	fields = append(fields, fmt.Sprintf("first_name=%s", data.FirstName))
	fields = append(fields, fmt.Sprintf("id=%d", data.ID))
	if data.Username != "" {
		fields = append(fields, fmt.Sprintf("username=%s", data.Username))
	}
	if data.PhotoURL != "" {
		fields = append(fields, fmt.Sprintf("photo_url=%s", data.PhotoURL))
	}

	// Сортируем и объединяем
	sort.Strings(fields)
	dataCheckString := strings.Join(fields, "\n")

	// Вычисляем HMAC
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(dataCheckString))
	return hex.EncodeToString(h.Sum(nil))
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func printJSON(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}
