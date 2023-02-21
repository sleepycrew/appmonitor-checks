package main

import (
	"github.com/sleepycrew/appmonitor-checks/check"
	"os"
	"plugin"
)

func main() {
	args := os.Args

	p, err := plugin.Open(args[1])
	if err != nil {
		panic(err)
	}
	printInfo(p)
	runCheck(p)
}

func printInfo(p *plugin.Plugin) {
	printInfo, err := p.Lookup("PrintPluginInfo")
	if err != nil {
		panic(err)
	}
	printInfo.(func())()
}

func runCheck(p *plugin.Plugin) {
	checks, err := p.Lookup("Checks")
	if err != nil {
		panic(err)
	}
	input := map[string]string{"name": "nginx", "status": "running"}
	c, err := checks.([]check.Factory)[0].BuildCheck(input)
	channel := make(chan check.Result)
	go c.RunCheck(channel)
	result := <-channel
	println(result.Value)
	println(result.Value)
}
