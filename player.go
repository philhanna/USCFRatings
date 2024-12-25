package uscfratings

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Player is structure representing a USCF member and their rating. It
// is also a work area for the ratings calculation.
type Player struct {
	USCFID    string
	Name      string
	Rating    float64
	NGames    int
	EffNGames int
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// AgeBasedRating returns the provisional rating described in Section
// 2.3 of the algorithm.  The full algorithm handles the case where the
// player is not rated in the basic USCF system but *is* rated in one
// of the other systems, like FIDE or Canada.  But this version does not
// implement those cases, using instead only the player's age.
func AgeBasedRating(age int) float64 {
	var rating float64

	// If the age is not specified (i.e., age == 0) then assume the age
	// is 26.  We do not currently support the case of the unrated player
	// being a youth.
	if age == 0 {
		age = 26
	}

	// Now calculate the rating based on which slot the age fits in.
	switch {
	case age < 2:
		rating = 100.0
	case age <= 26:
		rating = float64(age) * 50
	default:
		rating = 1300.0
	}

	return rating
}
