//go:build exclude
// +build exclude

package listener

import (
	"log"
	"os"
	"testing"

	"github.com/matheuspolitano/GadgetHub/utils"
)

var webhookHandlerTest WebhookHandler

func TestMain(m *testing.M) {
	conf, err := utils.LoadConfig("../.")
	if err != nil {
		log.Fatal(err)
	}

	webhookHandlerTest = WebhookHandler{conf}

	os.Exit(m.Run())
}
