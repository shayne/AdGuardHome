package home

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSubnetSet_default(t *testing.T) {
	set := parseSubnetSet(nil)

	testCases := []struct {
		name string
		addr netip.Addr
		want bool
	}{{
		name: "cgnat",
		addr: netip.MustParseAddr("100.64.0.1"),
		want: true,
	}, {
		name: "public",
		addr: netip.MustParseAddr("8.8.8.8"),
		want: false,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, set.Contains(tc.addr))
		})
	}
}
