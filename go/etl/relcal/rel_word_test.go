package relcal_test

import (
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

var (
	testSet1 []dbh.Word
	testSet2 []dbh.Word
)

func init() {
	testSet1 = []dbh.Word{
		dbh.Word{
			Id:            1,
			Word:          "grain",
			Pronunciation: util.NewString("G R EY N"),
		},
		dbh.Word{
			Id:            2,
			Word:          "brain",
			Pronunciation: util.NewString("B R EY N"),
		},
	}
	testSet2 = []dbh.Word{
		dbh.Word{
			Id:            1,
			Word:          "o",
			Pronunciation: util.NewString("OW"),
		},
		dbh.Word{
			Id:            2,
			Word:          "cry",
			Pronunciation: util.NewString("K R AY"),
		},
		dbh.Word{
			Id:            3,
			Word:          "dry",
			Pronunciation: util.NewString("D R AY"),
		},
		dbh.Word{
			Id:            4,
			Word:          "go",
			Pronunciation: util.NewString("G OW"),
		},
	}
}
