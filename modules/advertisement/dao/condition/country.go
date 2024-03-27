package condition

type CountryCondition struct {
	country string
}

func NewCountryCondition(country string) *CountryCondition {
	condition := new(CountryCondition)

	condition.country = country

	return condition
}

func (condition *CountryCondition) verify(user User) bool {

}
