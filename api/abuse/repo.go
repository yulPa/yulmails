package abuse

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yulpa/yulmails/api/utils"
)

// Abuse represents an abuse address
type Abuse struct {
	Id int `json:"id,omitempty"`
	// Name is the abuse address
	Name    string `json:"name"`
	Created string `json:"created,omitempty"`
}

// AbuseRepo is the interface to implements in order to manage
// abuse address
type AbuseRepo interface {
	ListAbuse() ([]*Abuse, error)
	GetAbuse(id int) (*Abuse, error)
	DeleteAbuse(id int) error
	InsertAbuse(*Abuse) error
}

type abuseRepo struct{ d *sql.DB }

// NewAbuseRepo returns a struct that implements abuse repo
// with a database connection
func NewAbuseRepo(db *sql.DB) *abuseRepo {
	return &abuseRepo{d: db}
}

// ListAbuse will return a list of abuses from the database
func (a *abuseRepo) ListAbuse() ([]*Abuse, error) {
	query := "SELECT * from abuse;"
	abuses := make([]*Abuse, 0)
	res, err := a.d.Query(query)
	if err != nil {
		return abuses, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var (
			id      int
			name    string
			created string
		)
		if err := res.Scan(&id, &name, &created); err != nil {
			return abuses, errors.Wrap(err, "unable to extract result")
		}
		abuses = append(abuses, &Abuse{
			Id:      id,
			Name:    name,
			Created: created,
		})
	}
	return abuses, nil
}

// GetAbuse returns an entity selected from the DB with its ID
func (a *abuseRepo) GetAbuse(id int) (*Abuse, error) {
	query := fmt.Sprintf("SELECT name, created FROM abuse WHERE id = %d", id)
	var (
		name    string
		created string
	)
	err := a.d.QueryRow(query).Scan(&name, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrapf(err, "unable to select abuse adress: %d", id)
	default:
		return &Abuse{
			Name:    name,
			Created: created,
			Id:      id,
		}, nil
	}

}

// DeleteAbuse removes an entity selected from the DB with its ID
func (a *abuseRepo) DeleteAbuse(id int) error {
	// first we assert that the record exists
	ab, err := a.GetAbuse(id)
	if err != nil {
		return errors.Wrapf(err, "unable to fetch abuse address: %d", id)
	}
	if ab == nil {
		return utils.NotFound
	}
	query := fmt.Sprintf("DELETE FROM abuse WHERE id = %d", id)
	if _, err := a.d.Exec(query); err != nil {
		return errors.Wrapf(err, "unable to delete abuse: %d", id)
	}
	return nil
}

// InsertAbuse creates and returns an entity
func (a *abuseRepo) InsertAbuse(ab *Abuse) error {
	ab.Created = time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO abuse(created, name) VALUES( '%s', '%s' ) RETURNING id;", ab.Created, ab.Name)
	var id int
	if err := a.d.QueryRow(query).Scan(&id); err != nil {
		return errors.Wrapf(err, "unable to insert abuse adress: %d", id)
	}
	ab.Id = id
	return nil
}
