package char

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
