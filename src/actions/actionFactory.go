package actions

type actionFactory interface {
	create() action
}
