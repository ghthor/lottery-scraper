package main

import (
	"fmt"
	"strings"
)

func outputOdds(games []LottoGame, profit int) {
	// See http://golang.org/pkg/fmt/ for format string documentation

	// Outputs as 3 columns
	// Left justified 30 width column
	// Right justified 8 width column
	// Left justified 20 width column
	fmtStr := "%-30s %5s  %-20s\n"

	fmt.Printf(fmtStr, "Game", "Cost", "Odds")
	fmt.Println(strings.Repeat("-", 30+8+20))

	for _, game := range games {
		// The `Cost` column needs an "$" appended to the left side of it
		// The width of the column for the integer is therefor 5-1 = 4
		fmt.Printf(fmtStr, game.Name, fmt.Sprintf("$%4d", game.Cost), game.OddsOfWinning(profit))
	}
}
