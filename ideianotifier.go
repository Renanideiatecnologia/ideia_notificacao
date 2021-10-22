package main

import (
	"ideianotifier/repository"
	usecases "ideianotifier/useCases"
)

func main() {
	pendingQty := repository.GetPendingQty()
	companyName := repository.GetCompanyName()
	usecases.SendNotification(companyName, pendingQty)
}
