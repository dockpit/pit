package model

type Action int

const (
	ActionCreated = Action(1)
	ActionRemoved = Action(2)
	ActionUpdate  = Action(3)
)

type Event struct {
	Action  Action
	Subject interface{}
}

func NewEvent(a Action, s interface{}) Event {
	return Event{
		Action:  a,
		Subject: s,
	}
}
