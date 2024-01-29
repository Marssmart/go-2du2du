package constants

const ScreenWidth = 640
const ScreenHeight = 480
const PlayerIconPath = "assets/pink_panther_32x32.png"
const GhostIconPath = "assets/ghost_32x32.png"
const HeartFullIconPath = "assets/heart_full_16x16.png"
const HeartEmptyIconPath = "assets/heart_empty_16x16.png"
const DefaultLives = 10
const DefaultIconHeight = 16
const DefaultIconWidth = 16
const BoardItemIconWidth = DefaultIconWidth * 2
const BoardItemIconHeight = DefaultIconHeight * 2
const StatusBarIconGap = 3
const ReservedRowsSpaceForStatusBar = (DefaultIconHeight * 2) - 5
const BoardItemHeightBoundary = ScreenHeight - ReservedRowsSpaceForStatusBar - BoardItemIconWidth
const BoardItemWidthBoundary = ScreenWidth - BoardItemIconWidth
const MarginX = 7
const MarginY = 10

// Rows 1 row reserved for status bar, 1 row for top bar, 1 row that starts in bounds but ends outside of bounds
const Rows = (ScreenHeight / BoardItemIconHeight) - 3
const Columns = ScreenWidth / BoardItemIconWidth
