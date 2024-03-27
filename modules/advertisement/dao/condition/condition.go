package condition

type Condition interface {
	verify() bool
}
