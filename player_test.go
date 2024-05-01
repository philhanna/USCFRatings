package uscfratings

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAGNUS_CARLSEN = "15218438"

func TestBuildURL(t *testing.T) {
	want := "https://www.uschess.org/msa/MbrDtlMain.php?15218438"
	have := BuildURL(MAGNUS_CARLSEN)
	assert.Equal(t, want, have)
}

func TestGetPage(t *testing.T) {
	tests := []struct {
		name   string
		USCFID string
	}{
		{
			USCFID : MAGNUS_CARLSEN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				GetPage = DefaultGetPage
			}()
			GetPage = func(id string) (string, error) {
				body, err := os.ReadFile("testdata/magnus.html")
				assert.Nil(t, err)
				return string(body), nil
			}
			page, err := GetPage(tt.USCFID)
			assert.Nil(t, err)
			assert.Contains(t, page, "US Chess MSA - Member Details")
			assert.Contains(t, page, tt.USCFID)
		})
	}
}
