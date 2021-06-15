package hcpairing_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
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
			[]string{
				hcpairing.ShortOfBreath,
				hcpairing.Sneeze,
				hcpairing.SoreEyes,
				hcpairing.SoreMuscles,
				hcpairing.SoreThroat,
				hcpairing.Stomachache,
				hcpairing.Stuffy,
			},
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
				sort.Strings(param.tags)
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
