package condition

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type GenderCondition struct {
	gender Gender
}

func NewGenderCondition(gender Gender) *GenderCondition {
	condition := new(GenderCondition)

	condition.gender = gender

	return condition
}

func (condition *GenderCondition) verify(user User) bool {

}
