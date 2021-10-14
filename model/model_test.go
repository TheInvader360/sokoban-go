package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	assert.IsType(t, &Model{}, m)
}
