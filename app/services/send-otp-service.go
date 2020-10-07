package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

// SendOTPEmail : sends a email to destinatary
func SendOTPEmail(destinataryEmail string, otp string) {
	fmt.Println("6")
	emailAddress := os.Getenv("EMAIL_ACCOUNNT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	protocol := os.Getenv("EMAIL_PROTOCOL")
	protocolPort := os.Getenv("EMAIL_PROTOCOL_PORT")

	e := email.NewEmail()
	e.From = "Arreglapp <" + emailAddress + ">"
	e.To = []string{destinataryEmail}
	e.Subject = "Activa tu cuenta"
	// e.Text = []byte("Activa tu cuenta")
	e.HTML = []byte("<p>Utiliza el siguente numero para activar tu cuenta:  " + otp + " </p>")
	e.Send(protocol+":"+protocolPort, smtp.PlainAuth("", emailAddress, emailPassword, protocol))
	fmt.Println("7")
}
