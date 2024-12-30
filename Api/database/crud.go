package database

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5"
)

type Thought struct {
	Thought string `json:"thought"`
	ThoughtType string `json:"thoughtType"`
	TimePosted time.Time `json:"timePosted"`
}

func (d *DB) InsertRow (thought string, thoughtType string) error {
	query := `INSERT INTO thoughts (thought, thoughttype, timeposted) VALUES (@thought, @thoughttype, @time)`
	timestamp := time.Now()

	args := pgx.NamedArgs{
		"thought": thought,
		"thoughttype": thoughtType,
		"time" : timestamp,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	
	return nil
}

func (d *DB) GetRows() ([]Thought, error) {
	query := `SELECT * FROM thoughts`
	rows, err := d.conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to get rows: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Thought])
}