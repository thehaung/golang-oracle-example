package stringutil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildStringWithParams(t *testing.T) {

	testCases := []struct {
		Name          string
		CheckResponse func(t *testing.T)
	}{
		{
			Name: "Ok",
			CheckResponse: func(t *testing.T) {
				want := "HaHaHa"
				got := BuildStringWithParams("Ha", "Ha", "Ha")
				require.Equal(t, want, got)
			},
		},
		{
			Name: "NotOk",
			CheckResponse: func(t *testing.T) {
				want := "Ha Ha Ha"
				got := BuildStringWithParams("Ha", "Ha", "Ha")
				require.NotEqual(t, want, got)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			tc.CheckResponse(t)
		})
	}
}
