package utils

import (
	"crypto/rand"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/configs"
	"github.com/louistwiice/go/fripclose/entity"
	logger "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const otpChars = "1234567890"

// Function used to parse API response
func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    errCode,
		Message: msg,
		Data:    data,
	})
}

// Allow to cypher a given word
func HashString(password string) (string, error) {
	if password == "" {
		return "", entity.ErrInvalidPassword
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare a cyphered word and a plain word
func CheckHashedString(plain_word, hashed_word string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_word), []byte(plain_word))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password associated with identifier")
	}
	return err
}

func SendMail(to []string, subject, message string, attachments... string) error {
	conf := configs.LoadConfigEnv()
	mail := gomail.NewMessage()

	addresses := make([]string, len(to))
	for i, recipient := range to {
		addresses[i] = mail.FormatAddress(recipient, "")
	}

	mail.SetHeader("From", conf.EmailUser)
	mail.SetHeader("To", addresses...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", message)

	for _, attach := range attachments {
		mail.Attach(attach)
	}
	
	mail_server := gomail.NewDialer("smtp.gmail.com", 587, conf.EmailUser, conf.EmailPassword)
	if err := mail_server.DialAndSend(mail); err != nil {
		logger.Error().Str("error", err.Error()).Msg("SEND_MAIL")
		return err
	}
	return nil
}

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

