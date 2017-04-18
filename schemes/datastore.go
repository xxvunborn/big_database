package schemes

import (
	"database/sql"
	"time"

	runner "gopkg.in/mgutz/dat.v1/sqlx-runner"
)

// Datastore sote data in db using postgresql as a backend
type Datastore struct {
	*runner.DB
}

// NewDatastore returns a new datastore instance or an error if
// a datasore cannot be returned
func NewDatastore(dsn string) (*Datastore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	runner.MustPing(db)

	r := runner.NewDB(db, "postgres")
	runner.LogQueriesThreshold = 1 * time.Millisecond
	return &Datastore{DB: r}, nil
}
