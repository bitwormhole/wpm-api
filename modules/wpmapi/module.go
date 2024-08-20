package wpmapi

import (
	wpmapi "github.com/bitwormhole/wpm-api"
	"github.com/bitwormhole/wpm-api/gen/main4wpmapi"
	"github.com/bitwormhole/wpm-api/gen/test4wpmapi"
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

// Module  ...
func Module() application.Module {
	mb := wpmapi.NewMainModule()
	mb.Components(main4wpmapi.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := wpmapi.NewTestModule()
	mb.Components(test4wpmapi.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
