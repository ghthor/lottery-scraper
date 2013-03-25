package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
		// TODO: Ask the user for the profit using stdin
		panic("invalid profit")
	} else {
		profit = flag_profit
	}

	games := <-gamesCh

	for _, game := range games {
		fmt.Println(game.OddsOfWinning(profit))
	}
}
