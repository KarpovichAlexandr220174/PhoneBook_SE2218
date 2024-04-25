package main

import (
	"PhoneBook_SE2218/pkg/models"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "home.tmpl", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// Генерируем случайный пароль
		randomPassword := generateRandomPassword()

		// Получаем почту и ник пользователя из запроса
		email := r.FormValue("email")
		username := r.FormValue("username")

		// Сохраняем данные пользователя в базе данных
		user := models.User{Username: username, Email: email, Password: randomPassword}
		if err := models.SaveUser(user); err != nil {
			http.Error(w, "Ошибка при сохранении данных пользователя", http.StatusInternalServerError)
			return
		}

		subject := "Регистрация в системе"
		body := fmt.Sprintf("Добро пожаловать, %s! Ваш единоразовый пароль: %s", username, randomPassword)

		err := sendMail(email, subject, body)
		if err != nil {
			http.Error(w, "Ошибка отправки письма", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// Показываем страницу с сообщением об успешной регистрации
	tmpl.ExecuteTemplate(w, "register.tmpl", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем почту или ник пользователя и пароль из формы входа
		emailOrUsername := r.FormValue("emailOrUsername")
		password := r.FormValue("password")

		// Получаем пользователя из базы данных по его почте или нику
		user, err := models.GetUser(emailOrUsername)
		if err != nil {
			http.Error(w, "Неверные данные для входа", http.StatusUnauthorized)
			return
		}

		// Проверяем, совпадает ли введенный пароль с паролем пользователя
		if user.Password != password {
			http.Error(w, "Неверные данные для входа", http.StatusUnauthorized)
			return
		}

		// Если пароль верный, выполняем необходимые действия для входа пользователя
		// Например, установим куки для сессии и выполним перенаправление на другую страницу
		http.Redirect(w, r, "/application", http.StatusSeeOther)
	} else {
		// Если метод запроса не POST, показываем страницу входа
		tmpl.ExecuteTemplate(w, "login.tmpl", nil)
	}
}

func sendMail(to, subject, body string) error {
	// Настройки SMTP-сервера Google
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	email := "bdauren06@gmail.com"
	password := "lrhcvmwjvdkbjrkc"

	// Авторизация на SMTP-сервере
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// Формирование письма
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	// Отправка письма
	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, email, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
