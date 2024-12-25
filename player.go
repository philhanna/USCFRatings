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
// 2.3 of the algorithm.
func AgeBasedRating(age int) float64 {
	if age == 0 {
		age = 26
	}
	var rating float64
	switch {
	case age < 2:
		rating = 100
	case age >= 2 && age <= 26:
		rating = float64(age) * 50
	default:
		rating = 1300
	}
	return rating
}
