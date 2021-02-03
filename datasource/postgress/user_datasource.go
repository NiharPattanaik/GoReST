package postgress

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBPool      *pgxpool.Pool
	databaseURL = "postgres://niharpattanaik:welcome123@localhost/mydb?sslmode=disable"
)

func init() {
	var err error
	DBPool, err = pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		panic(err)
	}

}
