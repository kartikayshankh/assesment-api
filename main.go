package main

import (
	handler "api/handler"
	Service "api/service"

	database "api/database"
)

// var itemsMap = make(map[string]*Item)

func main() {
	// e := echo.New()
	// db := new(ent.Client)
	db, err := database.NewEntClient("employee")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	service := Service.NewHandler(db)
	e := handler.MakeHTTPHandler(service)
	e.Logger.Fatal(e.Start(":8081"))
}
