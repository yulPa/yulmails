package whitelist

import (
	"database/sql"

	"github.com/pkg/errors"
)

// WhitelistRepo is the interface to implements in order to manage
// whitelist address
type WhitelistRepo interface {
	ListIP() ([]string, error)
}

type whitelistRepo struct{ d *sql.DB }

// NewWhitelistRepo returns a struct that implements whitelist repo
// with a database connection
func NewWhitelistRepo(db *sql.DB) *whitelistRepo {
	return &whitelistRepo{d: db}
}

// ListIP returns whitelist from the DB
func (w *whitelistRepo) ListIP() ([]string, error) {
	query := "SELECT * FROM whitelist;"
	ips := make([]string, 0)
	res, err := w.d.Query(query)
	if err != nil {
		return ips, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var ip string
		if err := res.Scan(&ip); err != nil {
			return ips, errors.Wrap(err, "unable to extract result")
		}
		ips = append(ips, ip)
	}
	return ips, nil

}
