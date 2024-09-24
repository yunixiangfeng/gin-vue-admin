package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pkgTest2"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	PkgTestServiceGroup  pkgTest.ServiceGroup
	PkgTest2ServiceGroup pkgTest2.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
