package main

import (
	"fmt"
	"ideianotifier/repository"
	usecases "ideianotifier/useCases"
)

func main() {
	fmt.Println("Verificando estoque pendente")
	pendingQty := repository.GetPendingQty()
	companyName := repository.GetCompanyName()
	usecases.SendNotification(companyName, pendingQty)
}
