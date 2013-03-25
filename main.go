package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var flag_profit int

func init() {
	flag.IntVar(&flag_profit, "profit", -1, "a min amount that you want to win")
}

func parseGames() <-chan []LottoGame {
	parsedGames := make(chan []LottoGame)

	go func() {
		bytes, err := ioutil.ReadFile("scratcher.txt")
		if err != nil {
			panic(err)
		}

		games, err := ParseLotteryGames(bytes)
		if err != nil {
			panic(err)
		}
		parsedGames <- games
	}()

	return parsedGames
}

func main() {
	flag.Parse()

	// Async load the LottoGames from the datafile
	gamesCh := parseGames()

	var profit int
	if flag_profit == -1 {
		// Profit parameter wasn't passed on the commandline
		// Prompt the user for the profit using stdin
		fmt.Println("Please input your required profit value in dollars and [Press Enter]")
		fmt.Print(">>> $")

		n, err := fmt.Scan(&profit)
		if n != 1 && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		profit = flag_profit
	}

	// Retrieve the parsed games
	games := <-gamesCh

	// Output the odds to stdout
	outputOdds(games, profit)
}
