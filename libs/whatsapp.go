package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func SendMsgWhatsApp(companyName string, pendingQtd int) {
	LoadENV()
	json_data, err := json.Marshal(
		formatBodyJSON(companyName, pendingQtd, os.Getenv("CHAT_ID")),
	)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(os.Getenv("URL_CHAT_API"), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	println("status code send email", resp.StatusCode)
}

func formatBodyJSON(companyName string, pendingQtd int, chatID string) map[string]string {
	values := map[string]string{
		"body":   fmt.Sprintf("*Analise de estoque pendente*\r\n\nEmpresa: %s \r\nQuantidade pendente: %d \r\n", companyName, pendingQtd),
		"chatId": chatID,
	}
	return values
}
