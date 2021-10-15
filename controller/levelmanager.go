package controller

type LevelManager struct {
	CurrentLevelNumber int
	Levels             []Level
}

type Level struct {
	Width, Height int
	MapData       string // Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
}

// NewLevelManager - Creates a level manager
func NewLevelManager() *LevelManager {
	lm := LevelManager{}

	lm.Levels = []Level{
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
	}

	return &lm
}

func (lm *LevelManager) NextLevel() {
	lm.CurrentLevelNumber++
	// Loop back round to the first level
	// Note: "> len(Levels)" is a safer way to achieve this, but this approach better suits the Jack OS API
	if lm.CurrentLevelNumber > 2 {
		lm.CurrentLevelNumber = 1
	}
}

func (lm *LevelManager) GetCurrentLevel() *Level {
	return &lm.Levels[lm.CurrentLevelNumber]
}
