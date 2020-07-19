package core

import (
	"catcher/utils"
	"path"
	"plugin"
)

func Start() {
	monitor := &Monitor{
		things: []*Thing{},
	}
	serviceRegistry := &ServiceRegistry{
		services: make(map[string][]string),
	}
	eventBus := &EventBus{
		bus: make(map[*Event][]*Service),
	}
	core := &Core{
		monitor:         monitor,
		serviceRegistry: serviceRegistry,
		eventBus:        eventBus,
	}
	loadComponents(core)
}

func loadComponents(core *Core) {
	// TODO 绝对路径
	dir := "E:\\Project\\go_workspace\\src\\catcher\\utils"
	com := "E:\\Project\\go_workspace\\src\\catcher\\components"
	components := utils.LoadComponents(dir)
	for k, v := range components {
		// TODO 组件相互依赖，加载有先后顺序
		// 异步加载组件
		go func() {
			p := path.Join(com, k)
			plugin, err := plugin.Open(p)
			if err != nil {
				panic(err)
			}
			f, err := plugin.Lookup("setup")
			if err != err {
				panic(err)
			}
			f.(func(interface{}))(v)
		}()
	}

}
