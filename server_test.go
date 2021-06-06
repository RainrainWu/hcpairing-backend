package hcpairing_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RainrainWu/hcpairing"
	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {

	params := []struct {
		prefix string
		tags   []string
	}{
		{
			"tooth",
			[]string{hcpairing.Toothache},
		},
		{
			"s",
			[]string{hcpairing.SoreMuscles, hcpairing.Stomachache},
		},
	}

	for _, param := range params {
		server := hcpairing.NewServer()
		recorder := httptest.NewRecorder()
		queryString := fmt.Sprintf("/v1/tags?prefix=%v", param.prefix)

		t.Run(
			queryString,
			func(t *testing.T) {
				req, _ := http.NewRequest("GET", queryString, nil)
				server.GetRouter().ServeHTTP(recorder, req)
				respBytes, _ := json.Marshal(
					map[string]interface{}{
						"tags": param.tags,
					},
				)
				assert.Equal(t, 200, recorder.Code)
				assert.Equal(t, string(respBytes), recorder.Body.String())
			},
		)
	}
}

func TestPostRecords(t *testing.T) {

	server := hcpairing.NewServer()
	recorder := httptest.NewRecorder()
	payload := map[string]interface{}{
		"state": "California",
		"tags":  []string{hcpairing.Toothache},
	}
	paylodBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/records", bytes.NewBuffer(paylodBytes))
	req.Header.Set("Content-Type", "application/json")
	server.GetRouter().ServeHTTP(recorder, req)

	respBytes, _ := json.Marshal(
		map[string]interface{}{
			"specialties": []string{hcpairing.Dentistry, hcpairing.ChildDentistry},
		},
	)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, string(respBytes), recorder.Body.String())
}
