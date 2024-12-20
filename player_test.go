package uscfratings

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAGNUS_CARLSEN = "15218438"
const PHIL_HANNA = "12910923"

func localGetPage(USCFID string) (string, error) {
	body, err := os.ReadFile(fmt.Sprintf("testdata/%s.html", USCFID))
	return string(body), err
}

func TestAgeBasedRating(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{age: 10, want: 500},
		{age: 1, want: 100},
		{age: 26, want: 1300},
		{age: 70, want: 1300},
		{age: 0, want: 1300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := AgeBasedRating(tt.age)
			assert.Equal(t, want, have)
		})
	}
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
			GetPage = localGetPage
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
			GetPage: localGetPage,
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
			GetPage: localGetPage,
			want: &Player{
				USCFID: PHIL_HANNA,
				Name:   "PHIL HANNA",
				Rating: 829.0,
				NGames: 22,
			},
			wantErr: false,
		},
		{
			USCFID:  "BOGUS",
			GetPage: localGetPage,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				GetPage = DefaultGetPage
			}()
			GetPage = localGetPage
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
