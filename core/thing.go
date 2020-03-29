package core

import (
	"container/list"
	"fmt"
)

type Property struct {
}

type State struct {
}

type Function struct {
}

type Thing struct {
	state      State
	properties list.List
	function   list.List
}

func (thing *Thing) changeState() {

	core.monitor.transform(Transformation{})

}

func (thing *Thing) execute(function Function) {
	core.monitor.transform(Transformation{})
	fmt.Println("execute some function")
}
