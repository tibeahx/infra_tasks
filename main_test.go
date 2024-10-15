package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValues(t *testing.T) {
	cases := []struct {
		name           string
		phones         int
		expectedPhones int
		wg             *sync.WaitGroup
		td             *testData
	}{
		{name: "successGenerate", phones: 100, expectedPhones: 100, wg: &sync.WaitGroup{}, td: &testData{}},
		{name: "negativePhones", phones: -1, expectedPhones: 0, wg: &sync.WaitGroup{}, td: &testData{}},
		{name: "zeroPhones", phones: 0, expectedPhones: 0, wg: &sync.WaitGroup{}, td: &testData{}},
		{name: "avgPhones", phones: 50, expectedPhones: 50, wg: &sync.WaitGroup{}, td: &testData{}},
		{name: "nilTestData", phones: 10, expectedPhones: 0, wg: &sync.WaitGroup{}, td: nil},
	}

	for _, tc := range cases {
		name := tc.name
		t.Run(name, func(t *testing.T) {
			if tc.td != nil {
				assert.NotNil(t, tc.td, "want initialized testdata, got nil")
			}
			generate(tc.wg, tc.phones, tc.td)
			if tc.wg != nil {
				tc.wg.Wait()
			}
			if tc.td != nil {
				assert.Equal(t, tc.expectedPhones, len(tc.td.phones), "want %d, got %d numbers", tc.expectedPhones, tc.td.phones)
			}
		})
	}
}
