package rulengine

import (
	"catcher/core"
	"fmt"
)

type EngineConsumer struct {
	Rules []*Rule
}

// check conditions and execute actions
func (ec *EngineConsumer) Execute(e *core.Event) error {

	fmt.Println(e.EventType)
	fmt.Println(e.Data)

	// find conditions and actions eventType -> []{conditions, actions}
	return nil
}
