package core

type Transformation struct {
	old interface{}
	new interface{}
}

type Monitor struct {
}

func (monitor *Monitor) transform(data Transformation) {

	core.eventBus.Trigger(Event{
		EventType: EVENT_EXECUTE_FUNCTION,
		EventData: nil,
	})

}
