package main

import (
	"database/sql"
	"log"

	pb "github.com/yulpa/yulmails/services/conservation/v1beta1"

	_ "github.com/lib/pq"
)

// DaoService is the struct in order to
// perform sql request against db
type Dao struct{ db *sql.DB }

// GetConservations returns a list of
// conservation rules from db
func (d *Dao) GetConservations() ([]pb.Conservation, error) {
	var conservations []pb.Conservation
	request := "SELECT * FROM conservation;"
	log.Printf("executing: %s", request)
	rows, err := d.db.Query(request)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conservation pb.Conservation
		if err := rows.Scan(
			&conservation.ID,
			&conservation.Created,
			&conservation.Sent,
			&conservation.Unsent,
			&conservation.KeepEmailContent,
		); err != nil {
			return nil, err
		}
		conservations = append(conservations, conservation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return conservations, nil
}

// CreateConservation add a conservation
// into the db
func (d *Dao) CreateConservation(conservation *pb.Conservation) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	insertConservation, err := tx.Prepare(
		"INSERT INTO conservation(created,sent,unsent,keep_email_content) " +
			"VALUES ($1, $2, $3, $4) RETURNING id;",
	)
	if err != nil {
		return err
	}
	if err := tx.Stmt(insertConservation).QueryRow(
		conservation.GetCreated(),
		conservation.GetSent(),
		conservation.GetUnsent(),
		conservation.GetKeepEmailContent(),
	).Scan(&conservation.ID); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
