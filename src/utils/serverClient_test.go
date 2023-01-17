package utils_test

import (
	"pcs/models"
	"pcs/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerClient(t *testing.T) {

	utils.LoadEnv()

	t.Run(
		"Client receives a successful response",
		func(t *testing.T) {
			client := utils.GetServerClientInstance()
			snapshot := models.BuildSnapshot()
			statusCode, err := client.WriteSnapshot(snapshot)

			assert.Nil(t, err)
			assert.Equal(t, statusCode, 200)
		},
	)
}
