package main

import (
	"texas-poker/poker"
)

var MatchSamplesPaths = []string{
	"./match_samples/seven_cards_with_ghost.result.json",
	//"./match_samples/seven_cards_with_ghost.json",
	//"./match_samples/five_cards_with_ghost.json",
	//"./match_samples/match.json",
}

func main() {

	for _, path := range MatchSamplesPaths {
		poker.MustGetMatchesFromMatchSamples(path).PrintCompareResult()
	}

}
