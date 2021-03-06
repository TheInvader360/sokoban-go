package controller

import (
	"testing"

	"github.com/TheInvader360/sokoban-go/model"
	"github.com/faiface/pixel/pixelgl"
	"github.com/stretchr/testify/assert"
)

func TestStartNewGameAndRestartLevel(t *testing.T) {
	m := model.Model{LM: model.NewLevelManager(false)}
	c := NewController(&m)
	assert.IsType(t, &Controller{}, c)

	// start new game - level 1, player at start position
	c.StartNewGame()
	assert.Equal(t, 1, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, "  ###     #.#     # #######$ $.##. $@#######$#     #.#     ###  ", m.LM.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.Y == 4)

	// move player away from start position
	c.HandleInput(pixelgl.KeyUp)
	assert.False(t, m.Board.Player.Y == 4)

	// restart level - level 1, player back at start position
	c.HandleInput(pixelgl.KeyR)
	assert.Equal(t, "  ###     #.#     # #######$ $.##. $@#######$#     #.#     ###  ", m.LM.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.Y == 4)
}

func TestPlayerMovementAndWallCollisions(t *testing.T) {
	mapData := "" +
		"####" +
		"#  #" +
		"# @#" +
		"####"
	b := model.NewBoard(mapData, 4, 4)
	m := model.Model{Board: b}
	c := Controller{m: &m}

	// start position
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// move up (first attempt succeeds, second attempt fails)
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move left (first attempt succeeds, second attempt fails)
	c.HandleInput(pixelgl.KeyLeft)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)
	c.HandleInput(pixelgl.KeyLeft)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move down (first attempt succeeds, second attempt fails)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// move right (first attempt succeeds, second attempt fails)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
}

func TestPlayerAndBoxMovementAndCollisions(t *testing.T) {
	mapData := "" +
		"#######" +
		"#.  ..#" +
		"#$@$ $#" +
		"#     #" +
		"#######"
	b := model.NewBoard(mapData, 7, 5)
	m := model.Model{Board: b}
	c := Controller{m: &m}

	// start position
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// try move left (fail: can't push box into wall)
	c.HandleInput(pixelgl.KeyLeft)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// try move right (success: box pushed to the right)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, 3, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	assert.False(t, b.Get(3, 2).HasBox)
	assert.True(t, b.Get(4, 2).HasBox)

	// try move right (fail: can't push box into other box)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, 3, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
}

func TestBoardCompletion(t *testing.T) {
	mapData := "" +
		"####" +
		"#..#" +
		"#$$#" +
		"#@ #" +
		"####"
	b := model.NewBoard(mapData, 4, 5)
	m := model.Model{Board: b}
	c := Controller{m: &m}

	// start position
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 3, b.Player.Y)
	assert.False(t, c.m.Board.IsComplete())

	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	assert.False(t, c.m.Board.IsComplete())

	c.HandleInput(pixelgl.KeyUp)
	assert.True(t, c.m.Board.IsComplete())
}

func TestStateLevelComplete(t *testing.T) {
	m := model.Model{LM: model.NewLevelManager(true)}
	m.State = model.StateLevelComplete
	c := Controller{m: &m}
	m.LM.ProgressToNextLevel()

	// input other than the space key has no effect
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, 1, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StateLevelComplete, m.State)

	// press the space key to start the next level
	c.HandleInput(pixelgl.KeySpace)
	assert.Equal(t, 2, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
}

func TestStateGameComplete(t *testing.T) {
	m := model.Model{LM: model.NewLevelManager(true)}
	m.State = model.StateLevelComplete
	c := Controller{m: &m}
	m.LM.ProgressToNextLevel()
	m.LM.ProgressToNextLevel()
	m.LM.ProgressToNextLevel()

	// simulate completion of the last level
	c.tryStartNextLevel()
	assert.Equal(t, model.StateGameComplete, m.State)

	// input other than the space key has no effect
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, 3, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StateGameComplete, m.State)

	// press the space key to start a new game
	c.HandleInput(pixelgl.KeySpace)
	assert.Equal(t, 1, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
}

func TestWalkthrough(t *testing.T) {
	m := model.NewModel()
	c := NewController(m)
	c.StartNewGame()

	assert.Equal(t, 1, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 2, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 3, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 4, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 5, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 6, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 7, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 8, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 9, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, 10, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyLeft)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyDown)
	c.HandleInput(pixelgl.KeyRight)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	c.HandleInput(pixelgl.KeyUp)
	assert.Equal(t, model.StateLevelComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)

	assert.Equal(t, model.StateGameComplete, m.State)
	c.HandleInput(pixelgl.KeySpace)
	assert.Equal(t, 1, m.LM.GetCurrentLevelNumber())
	assert.Equal(t, model.StatePlaying, m.State)
}
