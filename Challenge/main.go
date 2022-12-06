package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	con, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.QuotesAnime{})

	quo := repository.NewQuote(con)

	http.HandleFunc("/insert-quotes-from-api", q.InsertQuotesFromApi)
	http.HandleFunc("/get-quotes-from-api", q.GetQuotesFromDB)
	http.HandleFunc("/get-total-quotes", q.GetTotalQuotes)
	http.HandleFunc("/insert-quote", q.InsertQuotes)

	log.Println("Listening...at port 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}
}
