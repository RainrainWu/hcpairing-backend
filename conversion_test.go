package hcpairing_test

import (
	"fmt"
	"testing"

	"github.com/RainrainWu/hcpairing"
	"github.com/stretchr/testify/assert"
)

func sliceCompare(a, b []string) bool {

	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDirectConversion(t *testing.T) {

	params := []struct {
		tags   []string
		result []string
	}{
		{
			[]string{hcpairing.Toothache},
			[]string{hcpairing.Dentistry, hcpairing.ChildDentistry},
		},
		{
			[]string{hcpairing.Cough, hcpairing.Pregnancy},
			[]string{
				hcpairing.Pneumology,
				hcpairing.FamilyMedicine,
				hcpairing.Cardiology,
				hcpairing.Otorhinolaryngology,
				hcpairing.Obstetrics,
			},
		},
	}

	for _, param := range params {
		t.Run(
			fmt.Sprintf("%v", param.tags),
			func(t *testing.T) {
				specialties := hcpairing.DirectConversion(param.tags, -1)
				assert.EqualValues(t, specialties, param.result)
			},
		)
	}
}
