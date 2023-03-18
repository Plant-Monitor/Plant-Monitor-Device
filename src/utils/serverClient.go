package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pcs/models"
	"pcs/models/dto"
	"sync"
)

type ServerClient struct {
	hostUri    ServerHostUri
	httpClient *http.Client
}

var serverClientInstance *ServerClient
var serverClientLock *sync.Mutex = &sync.Mutex{}

type ServerHostUri string

func newServerClient(initParams ...any) *ServerClient {
	return &ServerClient{
		ServerHostUri(os.Getenv("SERVER_URI")),
		&http.Client{},
	}
}

func GetServerClientInstance() *ServerClient {

	return GetSingletonInstance(
		serverClientInstance,
		serverClientLock,
		newServerClient,
		nil,
	)
}

func (client *ServerClient) WriteSnapshot(snapshot *models.Snapshot) (statusCode int, err error) {
	snapshotJSON, _ := json.Marshal(snapshot)
	requestBody := bytes.NewBuffer(snapshotJSON)

	resp, err := http.Post(
		string(fmt.Sprintf("%s/snapshots", client.hostUri)),
		"application/json",
		requestBody,
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	return statusCode, err
}

func (client *ServerClient) WriteAction(dto dto.ActionDto) (statusCode int, err error) {
	snapshotJSON, _ := json.Marshal(dto)
	requestBody := bytes.NewBuffer(snapshotJSON)

	resp, err := http.Post(
		string(fmt.Sprintf("%s/actions/create", client.hostUri)),
		"application/json",
		requestBody,
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	return statusCode, err
}
