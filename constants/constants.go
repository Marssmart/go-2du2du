package constants

const (
	ScreenWidth                   = 640
	ScreenHeight                  = 480
	PlayerIconPath                = "assets/pink_panther_32x32.png"
	GhostIconPath                 = "assets/ghost_32x32.png"
	DevilIconPath                 = "assets/devil_32x32.png"
	HeartFullIconPath             = "assets/heart_full_16x16.png"
	HeartEmptyIconPath            = "assets/heart_empty_16x16.png"
	DefaultLives                  = 10
	DefaultIconHeight             = 16
	DefaultIconWidth              = 16
	BoardItemIconWidth            = DefaultIconWidth * 2
	BoardItemIconHeight           = DefaultIconHeight * 2
	StatusBarIconGap              = 3
	ReservedRowsSpaceForStatusBar = (DefaultIconHeight * 2) - 5
	BoardItemHeightBoundary       = ScreenHeight - ReservedRowsSpaceForStatusBar - BoardItemIconWidth
	BoardItemWidthBoundary        = ScreenWidth - BoardItemIconWidth
	MarginX                       = 7
	MarginY                       = 10

	// Rows 1 row reserved for status bar, 1 row for top bar, 1 row that starts in bounds but ends outside of bounds
	Rows    = (ScreenHeight / BoardItemIconHeight) - 3
	Columns = ScreenWidth / BoardItemIconWidth
)
