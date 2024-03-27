package condition

type Platform string

const (
	android Platform = "android"
	ios     Platform = "ios"
	web     Platform = "web"
)

type PlatformCondition struct {
	platform Platform
}

func NewPlatformCondition(platform Platform) *PlatformCondition {
	condition := new(PlatformCondition)

	condition.platform = platform

	return condition
}

func (condition PlatformCondition) verify(user User) bool {

}
