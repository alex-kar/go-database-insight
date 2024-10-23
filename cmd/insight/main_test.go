package insight

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
	"time"
)

const (
	user = "postgres"
	pass = "postgres"
	db   = "postgres"
)

func Test(t *testing.T) {

	_ = pq.Driver{}
	_, err := Insight{}.Wrap()
	if err != nil {
		return
	}

	ctx := context.Background()
	ctr, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(db),
		postgres.WithUsername(user),
		postgres.WithPassword(pass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(ctr); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable")

	db, err := sql.Open("postgres", connStr)
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)
}
