package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewModel(t *testing.T) {
	m := NewModel()
	assert.IsType(t, &Model{}, m)
	assert.IsType(t, &LevelManager{}, m.LM)
	assert.IsType(t, &Board{}, m.Board)
}

func TestUpdate(t *testing.T) {
	m := NewModel()
	m.TickAccumulator = 18
	m.Update()
	assert.Equal(t, 19, m.TickAccumulator)
	m.Update()
	assert.Equal(t, 20, m.TickAccumulator)
	m.Update()
	assert.Equal(t, 0, m.TickAccumulator)
	m.Update()
	assert.Equal(t, 1, m.TickAccumulator)
}
