package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func PrintMessage(message string) error {
	jsonValue, err := json.Marshal(struct {
		Text string `json:"text"`
	}{
		Text: message,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.WithStack(fmt.Errorf("http request to slack %s error: %d", url, resp.StatusCode))
	}

	return nil
}

var url string

func Setup(destination string) error {
	url = destination
	return PrintMessage("ReefStatus Alerts Started")
}
