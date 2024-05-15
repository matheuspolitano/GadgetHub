//go:build !feature1
// +build !feature1

package chat

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/utils"
)

var templateTest *Template
var testQuerier db.Store

func TestMain(m *testing.M) {
	temp, err := loadChatTemplate("../.")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := utils.LoadConfig("../.")
	if err != nil {
		log.Fatal(err)
	}
	connPool, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	testQuerier = db.NewStore(connPool)
	templateTest = temp
	os.Exit(m.Run())
}
