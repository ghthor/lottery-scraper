package main

import (
	"errors"
	"io/ioutil"
	"testing"
)

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
