package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheuspolitano/GadgetHub/utils"
)

var testQuerier Querier

func TestMain(m *testing.M) {
	conf, err := utils.LoadConfig("../../..")
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	testQuerier = New(connPool)
	os.Exit(m.Run())
}
