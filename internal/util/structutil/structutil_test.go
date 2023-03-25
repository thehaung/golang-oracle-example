package structutil

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/thehaung/golang-oracle-example/domain"
	"testing"
)

func TestToJsonString(t *testing.T) {
	testCases := []struct {
		Name          string
		CheckResponse func(t *testing.T)
	}{
		{
			Name: "ConvertSuccess",
			CheckResponse: func(t *testing.T) {
				structData := domain.Staff{
					Name:         "Hau",
					Title:        "SE",
					Organization: "IT",
				}

				want := "{\n  \"id\": 0,\n  \"name\": \"Hau\",\n  \"team_name\": \"\",\n  \"organization\": \"IT\",\n  \"title\": \"SE\",\n  \"onboard_date\": \"0001-01-01T00:00:00Z\",\n  \"active\": false,\n  \"created_at\": \"0001-01-01T00:00:00Z\",\n  \"modified_at\": \"0001-01-01T00:00:00Z\"\n}"
				got := ToJsonString(structData)
				fmt.Println(got)
				require.Equal(t, want, got)
			},
		},
		{
			Name: "NullConvert",
			CheckResponse: func(t *testing.T) {
				want := "null"
				got := ToJsonString(nil)

				fmt.Println(got)
				require.Equal(t, want, got)
			},
		},
		{
			Name: "ShouldErr",
			CheckResponse: func(t *testing.T) {
				want := "{\"id\"\""
				got := ToJsonString(want)

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
