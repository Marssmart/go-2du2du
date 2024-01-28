package constants

const ScreenWidth = 640
const ScreenHeight = 480
const PlayerIconPath = "assets/cricket_32x32.png"
const HeartFullIconPath = "assets/heart_full_16x16.png"
const HeartEmptyIconPath = "assets/heart_empty_16x16.png"
const DefaultLives = 10
const DefaultIconHeight = 16
const DefaultIconWidth = 16
const PaddingWidth = DefaultIconWidth
const PaddingHeight = DefaultIconHeight
const PlayerIconWidth = DefaultIconWidth * 2
const PlayerIconHeight = DefaultIconHeight * 2
const StatusBarIconGap = 3
const ReservedRowsStatusBar = PlayerIconHeight

// keep space for status bar
const PlayerHeightBoundary = ScreenHeight - ReservedRowsStatusBar - PlayerIconWidth
const PlayerWidthBoundary = ScreenWidth - PlayerIconWidth

// 1 row reserved for status bar
const Rows = ((ScreenWidth - PaddingWidth) / PlayerIconWidth) - 1
const Columns = (ScreenHeight - PaddingHeight) / PlayerIconHeight
