package plugins

import "github.com/sleepycrew/appmonitor-client/pkg/check"

type CheckFactory interface {
	GetName() string
	BuildCheck(jsonInput string) (check.Check, error)
}
