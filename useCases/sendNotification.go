package usecases

import (
	"ideianotifier/libs"
	"log"
	"os"
	"strconv"
)

func SendNotification(companyName string, pendingQty int) {

	limitPendingQty := getPendingQty()

	if pendingQty > limitPendingQty {
		libs.LoadENV()
		msg := []byte(libs.FormatMsgEMAIL(os.Getenv("FROM"), os.Getenv("TO"), "Analise de estoque pendente", companyName, pendingQty))

		libs.SendEmail(os.Getenv("FROM"), os.Getenv("PASSWORD"), []string{os.Getenv("TO")}, os.Getenv("SMPT_HOST"), os.Getenv("SMPT_PORT"), msg)

		libs.SendMsgWhatsApp(companyName, pendingQty)
	}
}

func getPendingQty() int {
	limitPendingQty := 1000

	configLimit, notEmpty := os.LookupEnv("LIMIT_PENDING_QTY")

	if notEmpty {
		value, err := strconv.Atoi(configLimit)
		if err != nil {
			log.Fatal("Error read config limit", err)
		}
		limitPendingQty = value
	}
	return limitPendingQty
}
