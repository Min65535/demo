package role

type Role = uint

const (
	Admin Role = 1 + iota
	Operate
)

var Roles = map[Role]string{
	Admin:   "admin",
	Operate: "operate",
}
