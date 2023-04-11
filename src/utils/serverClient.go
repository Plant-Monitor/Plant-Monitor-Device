package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (client *ServerClient) WriteSnapshot(snapshot interface{}) (statusCode int, err error) {
	return client.write(
		snapshot,
		fmt.Sprintf("%s/snapshots", client.hostUri),
	)
}

func (client *ServerClient) CreateAction(action interface{}) error {
	_, err := client.write(
		action,
		fmt.Sprintf("%s/actions/create", client.hostUri),
	)
	return err
}

func (client *ServerClient) ResolveAction(resolvedActionDto interface{}) error {
	_, err := client.update(
		resolvedActionDto,
		fmt.Sprintf("%s/actions/resolve", client.hostUri),
	)
	return err
}

func (client *ServerClient) update(dto interface{}, endpoint string) (statusCode int, err error) {
	snapshotJSON, _ := json.Marshal(dto)
	requestBody := bytes.NewBuffer(snapshotJSON)

	req, err := http.NewRequest(
		"PUT",
		endpoint,
		requestBody,
	)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	statusCode = resp.StatusCode

	return statusCode, err
}

func (client *ServerClient) write(dto interface{}, endpoint string) (statusCode int, err error) {
	snapshotJSON, _ := json.Marshal(dto)
	requestBody := bytes.NewBuffer(snapshotJSON)

	resp, err := http.Post(
		endpoint,
		"application/json",
		requestBody,
	)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	statusCode = resp.StatusCode

	return statusCode, err
}
