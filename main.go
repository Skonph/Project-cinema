package main

import (
	"fmt"

	"github.com/skonph/Project-cinema/movie"
	"github.com/skonph/Project-cinema/ticket"
)

// package movie

// package ticket

func main() {
	movieName := movie.FindName("tt4154769")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
