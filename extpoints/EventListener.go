package extpoints

type EventData interface {

}
type Event struct {
	target 	string
	data 	EventData;
}

type EventListener interface {
	Notify(event Event) error
}