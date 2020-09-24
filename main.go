package main

import (
	"log"
	"net/http"
)

var quotes = Quotes{
	Quote{Id: 0, Quotation: "I got a fever, and the only prescription is MOAR COWBELL!", Author: "Will Ferrell"},
	Quote{Id: 1, Quotation: "Knowledge is power.", Author: "Francis Bacon"},
	Quote{Id: 2, Quotation: "Life is really simple, but we insist on making it complicated.", Author: "Confucius"},
	Quote{Id: 3, Quotation: "This above all, to thine own self be true.", Author: "William Shakespeare"},
	Quote{Id: 4, Quotation: "Never complain. Never explain.", Author: "Katharine Hepburn"},
	Quote{Id: 5, Quotation: "Fly me to the moon...", Author: "Frank Sinatra"},
	Quote{Id: 6, Quotation: "Seventy-eight percent of internet quotes are made up.", Author: "Abraham Lincoln"},
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}
