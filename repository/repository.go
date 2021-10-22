package repository

import (
	"fmt"
	"ideianotifier/libs"
)

func GetPendingQty() int {
	var qty int

	db := libs.GetConnection()
	err := db.QueryRow("select count(*) as qtd from estoquelancamento e where coalesce(e.flagprocessado,0) <> 1").Scan(&qty)
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Close()
	return qty
}

func GetCompanyName() string {
	var name string
	db := libs.GetConnection()
	err := db.QueryRow("select e.empresa_nome from empresa e where coalesce(e.flagexcluido,0) = 0 order by e.empresa_codigo limit 1").Scan(&name)
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Close()
	return name
}
