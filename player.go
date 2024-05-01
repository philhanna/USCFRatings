package uscfratings

import "fmt"

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

// GetPage gets the page for the specified player from the USCF website.
// It is implemented as a variable so that it can be overridden in unit
// tests.
var (
	DefaultGetPage = func(USCFID string) (string, error) {
		url := BuildURL(USCFID)
		_ = url
		return "", nil // TODO write me
	}
	GetPage = DefaultGetPage
)

const (
	USCF_WEBSITE = "https://www.uschess.org"
)

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

