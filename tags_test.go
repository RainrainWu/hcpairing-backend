package hcpairing_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/RainrainWu/hcpairing"
	"github.com/stretchr/testify/assert"
)

func TestSearchTags(t *testing.T) {

	params := []struct {
		expect []string
		prefix string
	}{
		{
			[]string{hcpairing.Toothache},
			"to",
		},
		{
			[]string{
				hcpairing.ShortOfBreath,
				hcpairing.Sneeze,
				hcpairing.SoreEyes,
				hcpairing.SoreMuscles,
				hcpairing.SoreThroat,
				hcpairing.Stomachache,
				hcpairing.Stuffy,
			},
			"s",
		},
	}

	for _, param := range params {
		t.Run(
			fmt.Sprintf("%v", param.prefix),
			func(t *testing.T) {
				sort.Strings(param.expect)
				assert.Equal(t, param.expect, hcpairing.SearchTags(param.prefix))
			},
		)
	}
}
