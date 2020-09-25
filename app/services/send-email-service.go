package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

// SendEmail : sends a email to destinatary
func SendEmail(destinataryEmail string, username string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: " + err.Error())
	}

	emailAddress := os.Getenv("EMAIL_ACCOUNNT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	protocol := os.Getenv("EMAIL_PROTOCOL")
	protocolPort := os.Getenv("EMAIL_PROTOCOL_PORT")

	e := email.NewEmail()
	e.From = "Arreglapp <" + emailAddress + ">"
	e.To = []string{destinataryEmail}
	e.Subject = "Activa tu cuenta"
	// e.Text = []byte("Activa tu cuenta")
	e.HTML = []byte("<p>Haz click en el siguiente link para activar tu cuenta: </p> <a href='http://192.168.0.11:8080/user-profile/" + username + "/activate'>ACTIVAR</a>")
	e.Send(protocol+":"+protocolPort, smtp.PlainAuth("", emailAddress, emailPassword, protocol))
}
