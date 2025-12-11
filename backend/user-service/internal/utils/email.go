package utils

import (
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/config"
)

// sendMail отправляет письмо с заданной темой и телом на указанный адрес.
// Используется библиотека gomail.v2
// для упрощения работы с SMTP.
func sendMail(subject, body, to string) error {
	host := config.AppConfig.SMTPHost
	portStr := config.AppConfig.SMTPPort
	user := config.AppConfig.SMTPUser
	pass := config.AppConfig.SMTPPass
	from := config.AppConfig.SMTPFrom

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 587
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	// отправляем HTML и текст (в простом примере — только текст)
	m.SetBody("text/plain", body)

	// Создаем диалер и отправляем письмо
	d := gomail.NewDialer(host, port, user, pass)
	return d.DialAndSend(m)
}

// SendConfirmationEmail формирует ссылку и отправляет письмо подтверждения
// на указанный email с помощью функции sendMail.
func SendConfirmationEmail(toEmail, token string) error {
	baseURL := config.AppConfig.AppBaseURL
	link := fmt.Sprintf("%s/api/user/confirm-email?token=%s", baseURL, token)
	body := "Пожалуйста, подтвердите свой email, перейдя по ссылке:\n\n" + link + "\n\nСсылка действительна 20 минут."
	return sendMail("Подтвердите email", body, toEmail)
}
