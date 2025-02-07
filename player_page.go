package uscfratings

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

// GetPage gets the page for the specified player from the USCF website.
// It is implemented as a variable so that it can be overridden in unit
// tests.  The default function definition is DefaultGetPage, which gets
// the page from the real website over the network.
var DefaultGetPage = func(USCFID string) (string, error) {
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

var GetPage = DefaultGetPage

const USCF_WEBSITE = "https://www.uschess.org"

// BuildURL creates the URL for the player detail page on the USCF website.
func BuildURL(USCFID string) string {
	url := fmt.Sprintf("%s/msa/MbrDtlMain.php?%s", USCF_WEBSITE, USCFID)
	return url
}

// GetPlayer returns the Player structure for the specified USCFID
func GetPlayer(USCFID string) (*Player, error) {

	// Get the HTML for this player's page on the USCF website
	page, err := GetPage(USCFID)
	if err != nil {
		return nil, err
	}

	// Parse it into a Player structure
	p, err := ParsePlayerPage(page)
	return p, err
}

// ParsePlayerPage returns the Player structure for the specified USCFID
func ParsePlayerPage(page string) (*Player, error) {
	var (
		err       error
		p         = new(Player)
		reName    = regexp.MustCompile(`<font size=\+1><b>(\d+): ([A-Z ]+)</b></font>`)
		reRating  = regexp.MustCompile(`(\d+)`)
		reBasedOn = regexp.MustCompile(`\(Based on (\d+) games\)`)
	)

	type State uint8

	const (
		INIT State = iota + 1
		LOOKING_FOR_RATING_HEADER
		LOOKING_FOR_RATING
		DONE
	)

	state := INIT
	p.NGames = 25 // Default number of games if not specified otherwise

	reader := strings.NewReader(page)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if state == DONE {
			break
		}
		line := scanner.Text()
		switch state {

		// Looking for the line with the ID and player name

		case INIT:
			m := reName.FindStringSubmatch(line)
			if m != nil {
				p.USCFID = m[1]
				p.Name = m[2]
				state = LOOKING_FOR_RATING_HEADER
			}

		// Looking for <td valign=top>Regular Rating</td>

		case LOOKING_FOR_RATING_HEADER:
			if line == "Regular Rating" {
				state = LOOKING_FOR_RATING
			}

		// Looking for the rating (and possibly number of games)
		case LOOKING_FOR_RATING:
			m := reRating.FindString(line)
			if m != "" {
				p.Rating, _ = strconv.ParseFloat(m, 64)
				b := reBasedOn.FindStringSubmatch(line)
				if b != nil {
					p.NGames, _ = strconv.Atoi(b[1])
				}
				state = DONE
			}
		}
	}
	return p, err
}
