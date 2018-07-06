package char

const (
	DefaultCharacterScale  float32 = 2.0 // 32px -> 64px
	DefaultCharacterSize   float32 = 80.0
	DefaultCharacterZIndex float32 = 1.0

	// These offset values center 64x64 characters in 80x80 tiles.
	DefaultCharacterOffsetX float32 = 8.0
	DefaultCharacterOffsetY float32 = 4.0
)

type Character interface {
	GetName() string

	GetX() int
	SetX(int)

	GetY() int
	SetY(int)

	GetHostility() bool
	SetHostility(bool)

	GetHitPoints() int
	SetHitPoints(int)

	GetDamage() int

	GetExperiencePoints() int

	Destroy()
}