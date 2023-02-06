package actions

type action interface {
	execute()
}

type automatedAction interface {
	action
}

type userAction interface {
	action
}
