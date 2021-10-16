package controller

type LevelManager struct {
	currentLevelNumber int
	levels             []Level
}

type Level struct {
	Width, Height int
	MapData       string // Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
}

// NewLevelManager - Creates a level manager
func NewLevelManager() *LevelManager {
	lm := LevelManager{}

	lm.levels = []Level{
		{},
		{
			Width:  7,
			Height: 3,
			MapData: "" +
				"#######" +
				"#@ $ .#" +
				"#######",
		},
		{
			Width:  3,
			Height: 7,
			MapData: "" +
				"###" +
				"#.#" +
				"#$#" +
				"#@#" +
				"#$#" +
				"#.#" +
				"###",
		},
		/*
			{
				Width:  9,
				Height: 9,
				MapData: "" +
					"#########" +
					"#    ..*#" +
					"# # #.#.#" +
					"# #  ...#" +
					"# $$$ # #" +
					"# $@$   #" +
					"#  $$## #" +
					"#       #" +
					"#########",
			},
		*/
	}

	return &lm
}

// GetCurrentLevelNumber - Returns the current level number
func (lm *LevelManager) GetCurrentLevelNumber() int {
	return lm.currentLevelNumber
}

// GetCurrentLevel - Returns the current level
func (lm *LevelManager) GetCurrentLevel() *Level {
	return &lm.levels[lm.currentLevelNumber]
}

// HasNextLevel - Returns true if the current level is not the last
func (lm *LevelManager) HasNextLevel() bool {
	// Note: len(levels) is a safer way to achieve this, but this approach better suits the Jack OS API
	return lm.currentLevelNumber < 2
}

// ProgressToNextLevel - Increments the current level
func (lm *LevelManager) ProgressToNextLevel() {
	lm.currentLevelNumber++
}

// Reset - Resets the level manager
func (lm *LevelManager) Reset() {
	lm.currentLevelNumber = 0
}
