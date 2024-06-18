package repository

import (
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/lib/pq"
)

type Sample struct {
	IndexColumn   int
	VarcharColumn string
	IntColumn     int
}

var db *sql.DB

func Init(dbp *sql.DB) {
	db = dbp
}

func SampleGetAll() ([](*Sample), error) {
	result := make([](*Sample), 0)

	rows, err := db.Query(`SELECT "indexcolumn", "varcharcolumn", "intcolumn" FROM sample`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sample := new(Sample)

		err = rows.Scan(&sample.IndexColumn, &sample.VarcharColumn, &sample.IntColumn)
		if err != nil {
			return nil, err
		}

		result = append(result, sample)
	}

	return result, nil
}

func SampleGetById(id int) (*Sample, error) {
	rows, err := db.Query(`SELECT "indexcolumn", "varcharcolumn", "intcolumn" FROM sample WHERE indexcolumn = $1`, id)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		sample := new(Sample)

		err = rows.Scan(&sample.IndexColumn, &sample.VarcharColumn, &sample.IntColumn)
		if err != nil {
			return nil, err
		}

		return sample, nil
	} else {
		return nil, errors.New("Sample with id = " + strconv.Itoa(id) + " not exist")
	}
}

func SampleInsert(sample *Sample) error {
	_, err := db.Query("INSERT INTO sample (varcharcolumn, intcolumn) VALUES ($1, $2);", sample.VarcharColumn, sample.IntColumn)

	return err
}
