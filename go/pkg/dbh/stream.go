package dbh

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

var (
	ErrZeroColumn    = errors.New("dbh: Zero column to insert")
	ErrValColNoMatch = errors.New("dbh: Number of columns and number of values don't match")
)

func batchInsertRows(db *gorm.DB, table string, columns []string, values [][]interface{}) error {
	args := make([]interface{}, 0, len(columns)*len(values))
	placeholders := make([]string, 0, len(values))

	placeholder := strings.Repeat("?, ", len(columns))
	if placeholder == "" {
		return ErrZeroColumn
	}
	placeholder = placeholder[:len(placeholder)-2]
	placeholder = "(" + placeholder + ")"

	for _, val := range values {
		placeholders = append(placeholders, placeholder)
		if len(val) != len(columns) {
			return ErrValColNoMatch
		}

		args = append(args, val...)
	}
	stmt := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s;`,
		table, strings.Join(columns, ","), strings.Join(placeholders, ","))
	err := db.Exec(stmt, args...).Error
	return err
}

func streamAddRows(db *gorm.DB, table string, columns []string, c chan []interface{}) <-chan error {
	errC := make(chan error)

	go func() {
		// maximun # of mysql placeholder is 65536
		// 5(# of columns of relevant_word) * 10000 = 50000 50000 < 65536
		bufferSize := 10000
		buffer := make([][]interface{}, 0, bufferSize)

	loop:
		for {
			select {
			case val, more := <-c:
				if !more {
					break loop
				}
				buffer = append(buffer, val)

				if len(buffer) == bufferSize {
					err := batchInsertRows(db, table, columns, buffer)
					if err != nil {
						errC <- err
						return
					}

					buffer = make([][]interface{}, 0, bufferSize)
				}
			}
		}

		// flush remaining buffer
		if len(buffer) != 0 {
			err := batchInsertRows(db, table, columns, buffer)
			if err != nil {
				errC <- err
				return
			}
		}
		errC <- nil
	}()

	return errC
}
