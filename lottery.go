package main

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

func ParseLotteryGames(bytes []byte) (games []LottoGame, err error) {
	return
}
