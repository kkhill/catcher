package core

type Transformation struct {
	old interface{}
	new interface{}
}

/*
core do not control thing directly, but monitor thing through thing proxy,
it 's a decoupling strategy for core and manager, that control the lifecycle of thing
*/
type ThingProxy struct {
}

type Monitor struct {
}

func (monitor *Monitor) transform(data Transformation) {

	core.eventBus.Trigger(Event{
		EventType: EVENT_EXECUTE_FUNCTION,
		EventData: nil,
	})

}
