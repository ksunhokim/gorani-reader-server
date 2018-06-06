package relword_test

import "github.com/sunho/gorani-reader/server/pkg/dbh"

var (
	testSet1 []dbh.Word
	testSet2 []dbh.Word
)

func init() {
	grainPron := "G R EY N"
	brainPron := "B R EY N"
	testSet1 = []dbh.Word{
		dbh.Word{
			Id:            1,
			Word:          "grain",
			Pronunciation: &grainPron,
		},
		dbh.Word{
			Id:            2,
			Word:          "brain",
			Pronunciation: &brainPron,
		},
	}
	iPron := "AY"
	cryPron := "K R AY"
	dryPron := "D R AY"
	goPron := "G OW"
	testSet2 = []dbh.Word{
		dbh.Word{
			Id:            1,
			Word:          "i",
			Pronunciation: &iPron,
		},
		dbh.Word{
			Id:            2,
			Word:          "cry",
			Pronunciation: &cryPron,
		},
		dbh.Word{
			Id:            3,
			Word:          "dry",
			Pronunciation: &dryPron,
		},
		dbh.Word{
			Id:            4,
			Word:          "go",
			Pronunciation: &goPron,
		},
	}
}
