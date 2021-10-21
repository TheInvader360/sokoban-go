package model

type LevelManager struct {
	currentLevelNumber int
	levels             []Level
}

type Level struct {
	Width, Height int
	MapData       string // Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
}

// NewLevelManager - Creates a level manager
func NewLevelManager(testMode bool) *LevelManager {
	lm := LevelManager{}

	if testMode {
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
			{
				Width:  8,
				Height: 4,
				MapData: "" +
					"########" +
					"#.$  $.#" +
					"#*.@$  #" +
					"########",
			},
		}
	} else {
		lm.levels = []Level{
			{},
			/*
				{
					Width:  22,
					Height: 14,
					MapData: "" +
						"######################" +
						"#                 @$.#" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"#                    #" +
						"######################",
				},
			*/
			{
				// Level 1
				Width:  8,
				Height: 8,
				MapData: "" +
					"  ###   " +
					"  #.#   " +
					"  # ####" +
					"###$ $.#" +
					"#. $@###" +
					"####$#  " +
					"   #.#  " +
					"   ###  ",
			},
			{
				// Level 2
				Width:  9,
				Height: 9,
				MapData: "" +
					"#####    " +
					"#   #    " +
					"# $@# ###" +
					"# $$# #.#" +
					"### ###.#" +
					" ##    .#" +
					" #   #  #" +
					" #   ####" +
					" #####   ",
			},
			{
				// Level 3
				Width:  6,
				Height: 8,
				MapData: "" +
					" #### " +
					"##  # " +
					"# @$# " +
					"##$ ##" +
					"## $ #" +
					"#.$  #" +
					"#..*.#" +
					"######",
			},
			{
				// Level 4
				Width:  8,
				Height: 8,
				MapData: "" +
					" ####   " +
					" #@ ### " +
					" # $  # " +
					"### # ##" +
					"#.# #  #" +
					"#.$  # #" +
					"#.   $ #" +
					"########",
			},
			{
				// Level 5
				Width:  8,
				Height: 7,
				MapData: "" +
					"  ######" +
					"  #    #" +
					"###$$$ #" +
					"#@ $.. #" +
					"# $...##" +
					"####  # " +
					"   #### ",
			},
			{
				// Level 6
				Width:  8,
				Height: 7,
				MapData: "" +
					"  ##### " +
					"###  @# " +
					"#  $. ##" +
					"#  .$. #" +
					"### *$ #" +
					"  #   ##" +
					"  ##### ",
			},
			{
				// Level 7
				Width:  8,
				Height: 8,
				MapData: "" +
					"  ####  " +
					"  #..#  " +
					" ## .## " +
					" #  $.# " +
					"## $  ##" +
					"#  #$$ #" +
					"#  @   #" +
					"########",
			},
			{
				// Level 8
				Width:  8,
				Height: 7,
				MapData: "" +
					"########" +
					"#  #   #" +
					"#@$..$ #" +
					"# $.* ##" +
					"# $..$ #" +
					"#  #   #" +
					"########",
			},
			{
				// Level 9
				Width:  9,
				Height: 7,
				MapData: "" +
					"######   " +
					"#    #   " +
					"# $$$##  " +
					"#  #..###" +
					"##  ..$ #" +
					" # @    #" +
					" ########",
			},
			{
				// Level 10
				Width:  7,
				Height: 8,
				MapData: "" +
					"#######" +
					"#..$..#" +
					"#..#..#" +
					"# $$$ #" +
					"#  $  #" +
					"# $$$ #" +
					"#  #@ #" +
					"#######",
			},
			/*
				{
					// Level 11
					Width:  8,
					Height: 8,
					MapData: "" +
						"########" +
						"#@   $.#" +
						"#      #" +
						"#      #" +
						"#      #" +
						"#      #" +
						"#      #" +
						"########",
				},
				{
					// Level 12
					Width:  6,
					Height: 9,
					MapData: "" +
						"######" +
						"#@ $.#" +
						"#    #" +
						"#    #" +
						"#    #" +
						"#    #" +
						"#    #" +
						"#    #" +
						"######",
				},
				{
					// Level 13
					Width:  7,
					Height: 9,
					MapData: "" +
						"#######" +
						"#@  $.#" +
						"#     #" +
						"#     #" +
						"#     #" +
						"#     #" +
						"#     #" +
						"#     #" +
						"#######",
				},
				{
					// Level 14
					Width:  9,
					Height: 8,
					MapData: "" +
						"#########" +
						"#@    $.#" +
						"#       #" +
						"#       #" +
						"#       #" +
						"#       #" +
						"#       #" +
						"#########",
				},
				{
					// Level 15
					Width:  8,
					Height: 7,
					MapData: "" +
						"########" +
						"#@   $.#" +
						"#      #" +
						"#      #" +
						"#      #" +
						"#      #" +
						"########",
				},
			*/
		}
	}

	return &lm
}

// GetCurrentLevelNumber - Returns the current level number
func (lm *LevelManager) GetCurrentLevelNumber() int {
	return lm.currentLevelNumber
}

// GetFinalLevelNumber - Returns the final level number
func (lm *LevelManager) GetFinalLevelNumber() int {
	// Note: len(levels) is a safer way to achieve this, but the hard coded approach better suits the Jack OS API
	return len(lm.levels) - 1 //15
}

// GetCurrentLevel - Returns the current level
func (lm *LevelManager) GetCurrentLevel() *Level {
	return &lm.levels[lm.currentLevelNumber]
}

// HasNextLevel - Returns true if the current level is not the last
func (lm *LevelManager) HasNextLevel() bool {
	// Note: len(levels) is a safer way to achieve this, but the hard coded approach better suits the Jack OS API
	return lm.currentLevelNumber < lm.GetFinalLevelNumber()
}

// ProgressToNextLevel - Increments the current level
func (lm *LevelManager) ProgressToNextLevel() {
	lm.currentLevelNumber++
}

// Reset - Resets the level manager
func (lm *LevelManager) Reset() {
	lm.currentLevelNumber = 0
}
