package rpg

const (
	// SpaceComponent dimensions of characters.
	charScale float32 = 2.0 // 32 -> 64px
	charSizeX float32 = 80
	charSizeY float32 = 80

	// Character art is 64x64 when scaled, and needs to be slightly offset in order
	// to be centered in an 80x80 tile.
	charOffsetX float32 = 8
	charOffsetY float32 = 4
)

// A Character is a generic entity which can occupy space on a tile, including
// player characters and non-player characters.
type Character interface {
	GetX() int
	GetY() int
	SetX(int)
	SetY(int)

	GetName() string
	GetHostility() string
	SetHostility(string)

	GetHitPoints() int
	ModifyHitPoints(int)
	GetDamage() int
	GetXPBonus() int

	Destroy()
}
