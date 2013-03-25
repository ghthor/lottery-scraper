package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type LottoGame struct {
	Name   string
	Cost   int
	Prizes []Prize
}

type Prize struct {
	Value            int
	TotalTickets     int
	UnclaimedTickets int
}

func ParseLotteryGames(raw []byte) ([]LottoGame, error) {
	var games []LottoGame
	buf := bytes.NewBuffer(raw)

	for {
		game, err := ParseLotteryGame(buf)
		if err != nil {
			return nil, err
		}

		games = append(games, game)

		// Consume the blankline
		_, err = buf.ReadString('\n')

		// Check for EOF
		switch err {
		case io.EOF:
			return games, nil
		default:
			return nil, err
		case nil:
			continue
		}
	}

	panic("never reached")
}

func ParseLotteryGame(b *bytes.Buffer) (game LottoGame, err error) {

	// Consume Name
	name, err := b.ReadString('\n')
	if err != nil {
		return game, err
	}

	var cost, numPrizes int

	// Consume Game Cost and Total number of Prize Values
	n, err := fmt.Fscanln(b, &cost, &numPrizes)
	if n != 2 || err != nil {
		return game, err
	}

	prizes := make([]Prize, numPrizes)
	for i := 0; i < numPrizes; i++ {
		var value, totalTickets, unclaimedTickets int

		n, err := fmt.Fscanln(b, &value, &totalTickets, &unclaimedTickets)
		if n != 3 || err != nil {
			return game, err
		}

		prizes[i] = Prize{value, totalTickets, unclaimedTickets}
	}

	game = LottoGame{
		strings.TrimSpace(name),
		cost,
		prizes,
	}

	return game, nil
}
