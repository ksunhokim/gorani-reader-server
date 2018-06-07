package relword

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

type Edge struct {
	TargetId int
	Score    int
}

type Vertex struct {
	WordId int
	Edges  []Edge
}

type Graph struct {
	Vertexs []Vertex
	Reltype string
}

func (graph *Graph) UpsertToDB(db *gorm.DB) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = dbh.DeleteRelevantWords(tx, graph.Reltype)
	if err != nil {
		return
	}

	err = graph.addRelevantWords(tx)
	return
}

func (graph *Graph) addRelevantWords(db *gorm.DB) error {
	c := make(chan dbh.RelevantWord)
	errC := dbh.StreamAddRelevantWords(db, c)
	for _, v := range graph.Vertexs {
		for _, e := range v.Edges {
			word := dbh.RelevantWord{
				WordId:       v.WordId,
				TargetWordId: e.TargetId,
				RelType:      graph.Reltype,
				Score:        e.Score,
				VoteSum:      0,
			}
			c <- word
			select {
			case err := <-errC:
				close(c)
				return err
			default:
			}
		}
	}
	close(c)

	err := <-errC
	return err
}
