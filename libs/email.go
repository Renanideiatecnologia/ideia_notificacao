package libs

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(from string, password string, to []string, smtpHost string, smtpPort string, message []byte) {
	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		log.Fatal(err)
	}
}

func FormatMsgEMAIL(from, to, subject, name string, qtd int) string {
	return fmt.Sprintf("From: %s \r\nTo: %s \r\nSubject: %s \r\n\r\nAnalise de estoque pendente\r\n\nEmpresa: %s \r\nQuantidade pendente: %d \r\n", from, to, subject, name, qtd)
}
