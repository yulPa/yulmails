package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yulpa/yulmails/api/utils"
)

// Domain represents an domain
type Domain struct {
	Id int `json:"id,omitempty"`
	// Name is the domain address
	Name           string `json:"name"`
	Created        string `json:"created,omitempty"`
	EnvironmentID  int    `json:"environment_id"`
	ConservationID int    `json:"conservation_id"`
}

// DomainRepo is the interface to implements in order to manage
// domain address
type DomainRepo interface {
	ListDomain() ([]*Domain, error)
	GetDomain(id int) (*Domain, error)
	DeleteDomain(id int) error
	InsertDomain(*Domain) error
}

type domainRepo struct{ d *sql.DB }

// NewDomainRepo returns a struct that implements domain repo
// with a database connection
func NewDomainRepo(db *sql.DB) *domainRepo {
	return &domainRepo{d: db}
}

// ListDomain will return a list of domains from the database
func (a *domainRepo) ListDomain() ([]*Domain, error) {
	query := "SELECT id, name, created, conservation_id, environment_id FROM domain;"
	domains := make([]*Domain, 0)
	res, err := a.d.Query(query)
	if err != nil {
		return domains, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var (
			id             int
			name           string
			created        string
			conservationID int
			environmentID  int
		)
		if err := res.Scan(&id, &name, &created, &conservationID, &environmentID); err != nil {
			return domains, errors.Wrap(err, "unable to extract result")
		}
		domains = append(domains, &Domain{
			Id:             id,
			Name:           name,
			Created:        created,
			EnvironmentID:  environmentID,
			ConservationID: conservationID,
		})
	}
	return domains, nil
}

// GetDomain returns an domain selected from the DB with its ID
func (a *domainRepo) GetDomain(id int) (*Domain, error) {
	query := fmt.Sprintf("SELECT name, created, conservation_id, environment_id FROM domain WHERE id = %d", id)
	var (
		name           string
		created        string
		environmentID  int
		conservationID int
	)
	err := a.d.QueryRow(query).Scan(&name, &created, &conservationID, &environmentID)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrapf(err, "unable to select domain adress: %d", id)
	default:
		return &Domain{
			Name:           name,
			Created:        created,
			Id:             id,
			EnvironmentID:  environmentID,
			ConservationID: conservationID,
		}, nil
	}

}

// DeleteDomain removes an domain selected from the DB with its ID
func (a *domainRepo) DeleteDomain(id int) error {
	// first we assert that the record exists
	ab, err := a.GetDomain(id)
	if err != nil {
		return errors.Wrapf(err, "unable to fetch domain: %d", id)
	}
	if ab == nil {
		return utils.NotFound
	}
	query := fmt.Sprintf("DELETE FROM domain WHERE id = %d", id)
	if _, err := a.d.Exec(query); err != nil {
		return errors.Wrapf(err, "unable to delete domain: %d", id)
	}
	return nil
}

// InsertDomain creates and returns an domain
func (a *domainRepo) InsertDomain(ab *Domain) error {
	ab.Created = time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO domain(created, name, environment_id, conservation_id) VALUES( '%s', '%s', %d, %d) RETURNING id;", ab.Created, ab.Name, ab.EnvironmentID, ab.ConservationID)
	var id int
	if err := a.d.QueryRow(query).Scan(&id); err != nil {
		return errors.Wrapf(err, "unable to insert domain adress: %d", id)
	}
	ab.Id = id
	return nil
}
