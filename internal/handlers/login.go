package handlers

import (
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/schemas"
	"AuthService/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to login.")
	// Декодируем входные данные
	var input schemas.LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем пользователя из базы данных
	filters := make(map[string]interface{})
	filters["email"] = input.Email
	user, err := user_repo.GetUser(&filters)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Проверяем, соответствует ли предоставленный пароль хешу пароля
	isValid := utils.CheckPasswordHash(input.Password, user.Password)
	if !isValid {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Генерируем токен доступа
	token, err := utils.GenerateToken(user)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Возвращаем токен в ответе
	err = HandleJsonResponse(w, token)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
