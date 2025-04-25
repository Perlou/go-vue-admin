package main

import (
	"github.com/perlou/go-vue-admin/server/core"
	"github.com/perlou/go-vue-admin/server/global"
	_ "go.uber.org/automaxprocs"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
}
