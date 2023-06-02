package app

import (
	"context"
	"errors"
	"fmt"
	"go.temporal.io/sdk/workflow"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	return "Hello " + name + "!", nil
}

func GreetInSpanish(ctx context.Context, name string) (string, error) {
	base := "http://localhost:9999/get-spanish-greeting?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
