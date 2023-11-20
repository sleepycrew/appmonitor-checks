package systemd

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/sleepycrew/appmonitor-client/pkg/check"
	"github.com/sleepycrew/appmonitor-client/pkg/data/result"
)

type Service struct {
	Name   string
	Status string
}

func (s Service) RunCheck(res chan<- check.Result) {
	ctx := context.Background()
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		res <- check.Result{
			Result: result.Unknown,
			Value:  "Could not connect to dbus",
		}
	}
	var unitNames = []string{s.Name}
	units, err := conn.ListUnitsByNamesContext(ctx, unitNames)
	if err != nil {
		res <- check.Result{
			Result: result.Unknown,
			Value:  "Could retrieve units",
		}
	}

	// TODO handle activeState and substate
	// TODO proper result
	statusMatch := units[0].SubState == s.Status
	var code result.Code

	if statusMatch {
		code = result.OK
	} else {
		code = result.Error
	}

	res <- check.Result{
		Result: code,
		Value:  fmt.Sprintf("unit %v is %v", unitNames[0], units[0].SubState),
	}
}
