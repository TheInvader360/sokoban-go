package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelManager(t *testing.T) {
	lm := NewLevelManager()
	assert.IsType(t, &LevelManager{}, lm)
	assert.Equal(t, 0, lm.CurrentLevelNumber)
	assert.Equal(t, 0, lm.GetCurrentLevel().Width)
	assert.Equal(t, 0, lm.GetCurrentLevel().Height)
	assert.Equal(t, "", lm.GetCurrentLevel().MapData)

	lm.NextLevel()
	assert.Equal(t, 1, lm.CurrentLevelNumber)
	assert.Equal(t, 7, lm.GetCurrentLevel().Width)
	assert.Equal(t, 3, lm.GetCurrentLevel().Height)
	assert.Equal(t, "########@ $ .########", lm.GetCurrentLevel().MapData)

	lm.NextLevel()
	assert.Equal(t, 2, lm.CurrentLevelNumber)
	assert.Equal(t, 9, lm.GetCurrentLevel().Width)
	assert.Equal(t, 9, lm.GetCurrentLevel().Height)
	assert.Equal(t, "##########    ..*## # #.#.## #  ...## $$$ # ## $@$   ##  $$## ##       ##########", lm.GetCurrentLevel().MapData)

	lm.NextLevel()
	assert.Equal(t, 1, lm.CurrentLevelNumber)
	assert.Equal(t, 7, lm.GetCurrentLevel().Width)
	assert.Equal(t, 3, lm.GetCurrentLevel().Height)
	assert.Equal(t, "########@ $ .########", lm.GetCurrentLevel().MapData)
}