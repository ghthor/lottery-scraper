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
}
