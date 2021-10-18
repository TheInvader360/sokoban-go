package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelManager(t *testing.T) {
	lm := NewLevelManager(true)
	assert.IsType(t, &LevelManager{}, lm)
	assert.Equal(t, 0, lm.GetCurrentLevelNumber())
	assert.Equal(t, 0, lm.GetCurrentLevel().Width)
	assert.Equal(t, 0, lm.GetCurrentLevel().Height)
	assert.Equal(t, "", lm.GetCurrentLevel().MapData)

	lm.ProgressToNextLevel()
	assert.Equal(t, 1, lm.GetCurrentLevelNumber())
	assert.Equal(t, 7, lm.GetCurrentLevel().Width)
	assert.Equal(t, 3, lm.GetCurrentLevel().Height)
	assert.Equal(t, "########@ $ .########", lm.GetCurrentLevel().MapData)

	lm.ProgressToNextLevel()
	assert.Equal(t, 2, lm.GetCurrentLevelNumber())
	assert.Equal(t, 3, lm.GetCurrentLevel().Width)
	assert.Equal(t, 7, lm.GetCurrentLevel().Height)
	assert.Equal(t, "####.##$##@##$##.####", lm.GetCurrentLevel().MapData)

	lm.Reset()
	assert.Equal(t, 0, lm.GetCurrentLevelNumber())
	assert.Equal(t, 0, lm.GetCurrentLevel().Width)
	assert.Equal(t, 0, lm.GetCurrentLevel().Height)
	assert.Equal(t, "", lm.GetCurrentLevel().MapData)
}
