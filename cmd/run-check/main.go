package main

import (
	"github.com/sleepycrew/appmonitor-checks/check"
	"os"
	"plugin"
	"encoding/json"
)

func main() {
	args := os.Args

	p, err := plugin.Open(args[1])
	if err != nil {
		panic(err)
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
	input := map[string]string{}
    	json.Unmarshal([]byte(jsonStr), &input)
	c, err := checks.(*[1]check.Factory)[0].BuildCheck(input)
	if err != nil {
		panic(err)
	}

	channel := make(chan check.Result)
	go c.RunCheck(channel)
	result := <-channel
	println(result.Value)
	os.Exit(int(result.Result))
}
