package main

import (
	"log"
	"net/http"
)

var quotes = Quotes{
	Quote{Id: 0, Quotation: "I got a fever, and the only prescription is MOAR COWBELL!", Author: "Christopher Walken"},
	Quote{Id: 1, Quotation: "Knowledge is power.", Author: "Francis Bacon"},
	Quote{Id: 2, Quotation: "Life is really simple, but we insist on making it complicated.", Author: "Confucius"},
	Quote{Id: 3, Quotation: "This above all, to thine own self be true.", Author: "William Shakespeare"},
	Quote{Id: 4, Quotation: "Never complain, never explain.", Author: "Katharine Hepburn"},
	Quote{Id: 5, Quotation: "I did it my way", Author: "Frank Sinatra"},
	Quote{Id: 6, Quotation: "Yeah, well, that's just like ... your opinion ... man ...", Author: "The Dude"},
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}
