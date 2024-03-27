package condition

type AgeCondition struct {
	age int
}

func NewAgeCondition(age int) *AgeCondition {
	condition := new(AgeCondition)

	condition.age = age

	return condition
}

func (condition *AgeCondition) verify(user User) bool {

}
