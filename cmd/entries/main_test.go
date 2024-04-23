package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	doc := GetRoot()
	assert.NotNil(t, doc)
}
