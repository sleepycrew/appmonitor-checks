package main

import (
	"github.com/sleepycrew/appmonitor-checks/check"
	"github.com/sleepycrew/appmonitor-checks/pkg/systemd"
)

const Name = "Systemd"

var Checks = [...]check.Factory{systemd.ServiceFactory{}}

func PrintPluginInfo() {
	println("Module: ", Name)
	for _, f := range Checks {
		println("Check:", f.GetName())
	}
}
