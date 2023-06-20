package config

import (
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	httpClient := &http.Client{}

	httpClient.Timeout = 30 * time.Second

	return httpClient
}
