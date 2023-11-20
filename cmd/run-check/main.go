package main

import (
	"encoding/json"
	"github.com/sleepycrew/appmonitor-checks/internal/plugins"
	"github.com/sleepycrew/appmonitor-client/pkg/check"
	"os"
	"plugin"
)

func main() {
	args := os.Args

	p, err := plugin.Open(args[1])
	if err != nil {
		panic(err)
	}

	if !IsJSON(args[2]) {
		panic("Input is not valid JSON")
	}

	printInfo(p)
	runCheck(p, args[2])
}

func printInfo(p *plugin.Plugin) {
	printInfo, err := p.Lookup("PrintPluginInfo")
	if err != nil {
		panic(err)
	}
	printInfo.(func())()
}

func runCheck(p *plugin.Plugin, jsonStr string) {
	checks, err := p.Lookup("Checks")
	if err != nil {
		panic(err)
	}
	c, err := checks.(*[1]plugins.CheckFactory)[0].BuildCheck(jsonStr)
	if err != nil {
		panic(err)
	}

	channel := make(chan check.Result)
	go c.RunCheck(channel)
	result := <-channel
	println(result.Value)
	os.Exit(int(result.Result))
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
