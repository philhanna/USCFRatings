package uscfratings

import (
	"fmt"
	"io"
	"net/http"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Player is a USCF member
type Player struct {
	USCFID string
	Name   string
	Rating float32
	NGames int
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

// GetPage gets the page for the specified player from the USCF website.
// It is implemented as a variable so that it can be overridden in unit
// tests.
var (
	DefaultGetPage = func(USCFID string) (string, error) {
		url := BuildURL(USCFID)
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			errmsg := fmt.Errorf("status code was %d, not %d", resp.StatusCode, http.StatusOK)
			return "", errmsg
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		data := string(body)
		return data, nil
	}
	GetPage = DefaultGetPage
)

const (
	USCF_WEBSITE = "https://www.uschess.org"
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// BuildURL creates the URL for the player detail page on the USCF website.
func BuildURL(USCFID string) string {
	url := fmt.Sprintf("%s/msa/MbrDtlMain.php?%s", USCF_WEBSITE, USCFID)
	return url
}

// GetPlayer returns the Player structure for the specified USCFID
func GetPlayer(USCFID string) *Player {
	// page, err := GetPage(USCFID)
	return nil // TODO write me
}
