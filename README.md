# U. S. Chess Federation rating system revised September 2020

## Overview

This application implements the adjustment to a chess player's
USCF rating after a tournament.  The algorithm used is documented
on the USCF website at
https://new.uschess.org/sites/default/files/media/documents/the-us-chess-rating-system-revised-september-2020.pdf

The algorithm employs five steps:
- The first step sets temporary initial ratings for unrated players.
- The second step calculates an "effective" number of games played by each player.
- The third step calculates temporary estimates of ratings for certain unrated players
only to be used when rating their opponents on the subsequent step.
- The fourth step then calculates intermediate ratings for all players.
- The fifth step uses these intermediate ratings from the previous step as estimates of
opponents' strengths to calculate final post-event ratings.

