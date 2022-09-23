package utils

import "fmt"


func ActivationMail(first_name, last_name, otp_code string) string {
	message := fmt.Sprintf("<p>Bonjour %s %s</p> <p>Ton code OTP est le <b>%s</b>. Connectes-toi vite pour activer ton compte avant qu'il n'expire.</p> <b>Il n'est valable que 5mn</b>", first_name, last_name, otp_code)

	return message
}