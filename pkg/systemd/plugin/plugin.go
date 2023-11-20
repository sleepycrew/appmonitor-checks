package main

import (
	"encoding/json"
	"github.com/sleepycrew/appmonitor-checks/internal/plugins"
	"github.com/sleepycrew/appmonitor-checks/pkg/systemd"
	"github.com/sleepycrew/appmonitor-client/pkg/check"
)

const Name = "Systemd"

var Checks = [...]plugins.CheckFactory{ServiceFactory{}}

func PrintPluginInfo() {
	println("Module: ", Name)
	for _, f := range Checks {
		println("Check:", f.GetName())
	}
}

type ServiceFactory struct {
}

func (s ServiceFactory) GetName() string {
	return "SystemdService"
}

func (s ServiceFactory) BuildCheck(jsonInput string) (check.Check, error) {
	var serviceCheck systemd.Service

	if err := json.Unmarshal([]byte(jsonInput), &serviceCheck); err != nil {
		return nil, err
	}

	return serviceCheck, nil
}
