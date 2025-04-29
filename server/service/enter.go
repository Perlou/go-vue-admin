package service

import (
	"github.com/perlou/go-vue-admin/server/service/example"
	"github.com/perlou/go-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
