package systemd

import (
	"fmt"
	"context"
	"errors"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/sleepycrew/appmonitor-checks/check"
)

type Service struct {
	Name   string
	Status string
}

func (s Service) RunCheck(result chan<- check.Result) {
	ctx := context.Background()
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		result <- check.Result{
			Result: 1,
			Value:  "Could not connect to dbus",
		}
	}
	var unitNames = []string{s.Name}
	units, err := conn.ListUnitsByNamesContext(ctx, unitNames)
	if err != nil {
		result <- check.Result{
			Result: 1,
			Value:  "Could retrieve units",
		}
	}

	// TODO handle activeState and substate
	// TODO proper result
	statusMatch := units[0].SubState == s.Status
	var code int8

	if statusMatch {
		code = 0
	} else {
		code = 3
	}

	result <- check.Result{
		Result: code,
		Value: fmt.Sprintf("unit %v is %v", unitNames[0], units[0].SubState),
	}
}

type ServiceFactory struct {
}

func (s ServiceFactory) GetName() string {
	return "SystemdService"
}

func (s ServiceFactory) BuildCheck(input map[string]string) (check.Check, error) {
	name, ok := input["Name"]
	if !ok {
		return nil, errors.New("variable name not specified in input")
	}
	status, ok := input["Status"]
	if !ok {
		return nil, errors.New("variable status not specified in input")
	}

	serviceCheck := Service{
		Name:   name,
		Status: status,
	}
	return serviceCheck, nil
}
