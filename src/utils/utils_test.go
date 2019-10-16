package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanSize(t *testing.T) {
	cases := []struct {
		N uint64
		S string
	}{
		{0, "0B"},
		{10, "10.00B"},
		{5.5 * 1024, "5.50KB"},
		{5.5 * 1024 * 1024, "5.50MB"},
		{5.5 * 1024 * 1024 * 1024, "5.50GB"},
		{5.5 * 1024 * 1024 * 1024 * 1024, "5.50TB"},
		{5.5 * 1024 * 1024 * 1024 * 1024 * 1024, "5.50PB"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case:%d", i+1), func(t *testing.T) {
			assert.Equal(t, c.S, HumanSize(c.N))
		})
	}
}
