package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

func gamesAreEqual(a, b LottoGame) bool {
	if a.Name != b.Name {
		return false
	}

	if a.Cost != b.Cost {
		return false
	}

	if len(a.Prizes) != len(b.Prizes) {
		return false
	}

	for i := 0; i < len(a.Prizes); i++ {
		if a.Prizes[i] != b.Prizes[i] {
			return false
		}
	}

	return true
}

func TestDataParser(t *testing.T) {
	bytes, err := ioutil.ReadFile("scratcher.txt")

	if err != nil {
		t.Error(err)
	}

	games, err := ParseLotteryGames(bytes)

	if err != nil {
		t.Error(err)
	}

	if len(games) != 3 {
		t.Error(errors.New("Invalid number of games parsed"))
	}

	expectedGames := []LottoGame{{
		"SMALL BEANS",
		1,
		[]Prize{{
			10000, 10, 10,
		}, {
			5000, 50, 8,
		}, {
			100, 100, 6,
		}, {
			10, 1000, 4,
		}, {
			1, 10000, 2,
		}},
	}, {
		"PIRATE'S BOOTY, ARRR",
		10,
		[]Prize{{
			50000, 20, 0,
		}, {
			10000, 100, 0,
		}, {
			1000, 500, 0,
		}, {
			100, 2000, 738,
		}, {
			10, 7500, 2945,
		}, {
			1, 10000, 4476,
		}},
	}, {
		"BIG MONEY HU$TLA$",
		20,
		[]Prize{{
			1000000, 10, 7,
		}, {
			500000, 50, 29,
		}, {
			10000, 100, 78,
		}, {
			1000, 500, 396,
		}, {
			100, 2000, 1439,
		}, {
			20, 5000, 3218,
		}, {
			10, 10000, 6210,
		}},
	}}

	for i := 0; i < len(games); i++ {
		if !gamesAreEqual(games[i], expectedGames[i]) {
			t.Error(errors.New(fmt.Sprintf("     got: %v\nexpected: %v", games[i], expectedGames[i])))
		}
	}
}
