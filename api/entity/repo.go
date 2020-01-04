package entity

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yulpa/yulmails/api/utils"
)

// Entity represents an entity
type Entity struct {
	Id int `json:"id,omitempty"`
	// Name is the entity address
	Name           string `json:"name"`
	Created        string `json:"created,omitempty"`
	Description    string `json:"description,omitempty"`
	ConservationID int    `json:"conservation_id"`
	AbuseID        int    `json:"abuse_id"`
}

// EntityRepo is the interface to implements in order to manage
// entity address
type EntityRepo interface {
	ListEntity() ([]*Entity, error)
	GetEntity(id int) (*Entity, error)
	DeleteEntity(id int) error
	InsertEntity(*Entity) error
}

type entityRepo struct{ d *sql.DB }

// NewEntityRepo returns a struct that implements entity repo
// with a database connection
func NewEntityRepo(db *sql.DB) *entityRepo {
	return &entityRepo{d: db}
}

// ListEntity will return a list of entities from the database
func (a *entityRepo) ListEntity() ([]*Entity, error) {
	query := "SELECT id, name, created, description, conservation_id, abuse_id FROM entity;"
	entities := make([]*Entity, 0)
	res, err := a.d.Query(query)
	if err != nil {
		return entities, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var (
			id             int
			name           string
			created        string
			description    string
			conservationID int
			abuseID        int
		)
		if err := res.Scan(&id, &name, &created, &description, &conservationID, &abuseID); err != nil {
			return entities, errors.Wrap(err, "unable to extract result")
		}
		entities = append(entities, &Entity{
			Id:             id,
			Name:           name,
			Created:        created,
			Description:    description,
			ConservationID: conservationID,
			AbuseID:        abuseID,
		})
	}
	return entities, nil
}

// GetEntity returns an entity selected from the DB with its ID
func (a *entityRepo) GetEntity(id int) (*Entity, error) {
	query := fmt.Sprintf("SELECT name, created, description, conservation_id, abuse_id FROM entity WHERE id = %d", id)
	var (
		name           string
		created        string
		description    string
		conservationID int
		abuseID        int
	)
	err := a.d.QueryRow(query).Scan(&name, &created, &description, &conservationID, &abuseID)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrapf(err, "unable to select entity adress: %d", id)
	default:
		return &Entity{
			Name:           name,
			Created:        created,
			Id:             id,
			Description:    description,
			ConservationID: conservationID,
			AbuseID:        abuseID,
		}, nil
	}

}

// DeleteEntity removes an entity selected from the DB with its ID
func (a *entityRepo) DeleteEntity(id int) error {
	// first we assert that the record exists
	ab, err := a.GetEntity(id)
	if err != nil {
		return errors.Wrapf(err, "unable to fetch entity: %d", id)
	}
	if ab == nil {
		return utils.NotFound
	}
	query := fmt.Sprintf("DELETE FROM entity WHERE id = %d", id)
	if _, err := a.d.Exec(query); err != nil {
		return errors.Wrapf(err, "unable to delete entity: %d", id)
	}
	return nil
}

// InsertEntity creates and returns an entity
func (a *entityRepo) InsertEntity(ab *Entity) error {
	ab.Created = time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO entity(created, name, description, conservation_id, abuse_id) VALUES( '%s', '%s', '%s', %d, %d) RETURNING id;", ab.Created, ab.Name, ab.Description, ab.ConservationID, ab.AbuseID)
	var id int
	if err := a.d.QueryRow(query).Scan(&id); err != nil {
		return errors.Wrapf(err, "unable to insert entity adress: %d", id)
	}
	ab.Id = id
	return nil
}
