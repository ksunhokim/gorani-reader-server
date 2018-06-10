package relcal

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

func (graph *Graph) UpsertToDB(db *gorm.DB) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	reltype, err := dbh.GetRelevantWordTypeByName(tx, graph.RelType)
	if err != nil {
		reltype.Name = graph.RelType
		err = dbh.AddRelevantWordType(tx, &reltype)
		if err != nil {
			return
		}
	}

	err = dbh.DeleteRelevantWords(tx, reltype)
	if err != nil {
		return
	}

	err = graph.addRelevantWords(tx, reltype)
	return
}

func (graph *Graph) addRelevantWords(db *gorm.DB, reltype dbh.RelevantWordType) error {
	c := make(chan dbh.RelevantWord)
	errC := dbh.StreamAddRelevantWords(db, c)

	for _, v := range graph.Vertexs {
		for _, e := range v.Edges {
			word := dbh.RelevantWord{
				WordId:       v.WordId,
				TargetWordId: e.TargetId,
				TypeCode:     reltype.Code,
				Score:        e.Score,
				VoteSum:      0,
			}

			c <- word

			// check if there was an error
			select {
			case err := <-errC:
				close(c)
				return err
			default:
			}
		}
	}
	// in order to flush remaining buffer
	close(c)

	err := <-errC
	return err
}
