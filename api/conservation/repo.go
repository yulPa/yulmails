package conservation

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yulpa/yulmails/api/utils"
)

// Conservation represea conservation rule
type Conservation struct {
	ID      int    `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	// Sent is the number of day to keep the email in DB
	Sent int `json:"sent"`
	// Unsent is the number of day to keep unsent email in DB
	Unsent int `json:"unsent"`
	// KeepEmailContent will keep the body in DB
	KeepEmailContent bool `json:"keep_email_content"`
}

// ConservationRepo is the interface to implements in order to manage
// conservation address
type ConservationRepo interface {
	ListConservation() ([]*Conservation, error)
	GetConservation(id int) (*Conservation, error)
	DeleteConservation(id int) error
	InsertConservation(*Conservation) error
}

type conservationRepo struct{ d *sql.DB }

// NewConservationRepo returns a struct that implements conservation repo
// with a database connection
func NewConservationRepo(db *sql.DB) *conservationRepo {
	return &conservationRepo{d: db}
}

// ListConservation will return a list of conservations from the database
func (a *conservationRepo) ListConservation() ([]*Conservation, error) {
	query := "SELECT id, created, sent, unsent, keep_email_content from conservation;"
	conservations := make([]*Conservation, 0)
	res, err := a.d.Query(query)
	if err != nil {
		return conservations, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var (
			id               int
			created          string
			sent             int
			unsent           int
			keepEmailContent bool
		)
		if err := res.Scan(&id, &created, &sent, &unsent, &keepEmailContent); err != nil {
			return conservations, errors.Wrap(err, "unable to extract result")
		}
		conservations = append(conservations, &Conservation{
			ID:               id,
			Created:          created,
			Sent:             sent,
			Unsent:           unsent,
			KeepEmailContent: keepEmailContent,
		})
	}
	return conservations, nil
}

// GetConservation returns an entity selected from the DB with its ID
func (a *conservationRepo) GetConservation(id int) (*Conservation, error) {
	query := fmt.Sprintf("SELECT created, sent, unsent, keep_email_content FROM conservation WHERE id = %d", id)
	var (
		created          string
		sent             int
		unsent           int
		keepEmailContent bool
	)
	err := a.d.QueryRow(query).Scan(&created, &sent, &unsent, &keepEmailContent)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrapf(err, "unable to select conservation rule: %d", id)
	default:
		return &Conservation{
			Created:          created,
			ID:               id,
			Sent:             sent,
			Unsent:           unsent,
			KeepEmailContent: keepEmailContent,
		}, nil
	}

}

// DeleteConservation removes an entity selected from the DB with its ID
func (a *conservationRepo) DeleteConservation(id int) error {
	// first we assert that the record exists
	ab, err := a.GetConservation(id)
	if err != nil {
		return errors.Wrapf(err, "unable to fetch conservation rule: %d", id)
	}
	if ab == nil {
		return utils.NotFound
	}
	query := fmt.Sprintf("DELETE FROM conservation WHERE id = %d", id)
	if _, err := a.d.Exec(query); err != nil {
		return errors.Wrapf(err, "unable to delete conservation: %d", id)
	}
	return nil
}

// InsertConservation creates and returns an entity
func (a *conservationRepo) InsertConservation(ab *Conservation) error {
	ab.Created = time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO conservation(created, sent, unsent, keep_email_content) VALUES( '%s', %d, %d, %t) RETURNING id;", ab.Created, ab.Sent, ab.Unsent, ab.KeepEmailContent)
	var id int
	if err := a.d.QueryRow(query).Scan(&id); err != nil {
		return errors.Wrapf(err, "unable to insert conservation rule: %d", id)
	}
	ab.ID = id
	return nil
}
