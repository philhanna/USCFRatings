package uscfratings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAGNUS_CARLSEN = "15218438"

func TestBuildURL(t *testing.T) {
	want := "https://www.uschess.org/msa/MbrDtlMain.php?15218438"
	have := BuildURL(MAGNUS_CARLSEN)
	assert.Equal(t, want, have)
}
