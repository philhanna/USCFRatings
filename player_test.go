package uscfratings

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAGNUS_CARLSEN = "15218438"
const PHIL_HANNA = "12910923"

func LocalGetPage(USCFID string) (string, error) {
	body, err := os.ReadFile(fmt.Sprintf("testdata/%s.html", USCFID))
	return string(body), err
}

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
			USCFID: MAGNUS_CARLSEN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				GetPage = DefaultGetPage
			}()
			GetPage = LocalGetPage
			page, err := GetPage(tt.USCFID)
			assert.Nil(t, err)
			assert.Contains(t, page, "US Chess MSA - Member Details")
			assert.Contains(t, page, tt.USCFID)
		})
	}
}

func TestGetPlayer(t *testing.T) {
	tests := []struct {
		name    string
		USCFID  string
		GetPage func(ID string) (string, error)
		want    *Player
		wantErr bool
	}{
		{
			USCFID:  MAGNUS_CARLSEN,
			GetPage: LocalGetPage,
			want: &Player{
				USCFID: MAGNUS_CARLSEN,
				Name:   "MAGNUS CARLSEN",
				Rating: 2914.0,
				NGames: 25,
			},
			wantErr: false,
		},
		{
			USCFID:  PHIL_HANNA,
			GetPage: LocalGetPage,
			want: &Player{
				USCFID: PHIL_HANNA,
				Name:   "PHIL HANNA",
				Rating: 829.0,
				NGames: 22,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				GetPage = DefaultGetPage
			}()
			GetPage = LocalGetPage
			want := tt.want
			have, err := GetPlayer(tt.USCFID)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
				assert.Equal(t, want, have)
			}
		})
	}
}

func TestParsePlayerPage(t *testing.T) {
	tests := []struct {
		name    string
		page    string
		want    *Player
		wantErr bool
	}{
		{
			page: func() string {
				body, err := os.ReadFile("testdata/15218438.html")
				assert.Nil(t, err)
				return string(body)
			}(),
			want: &Player{
				USCFID: MAGNUS_CARLSEN,
				Name:   "MAGNUS CARLSEN",
				Rating: 2914.0,
				NGames: 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := ParsePlayerPage(tt.page)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
				assert.Equal(t, tt.want, have)
			}
		})
	}
}
