// +build go1.8

package main

import (
	_ "gitlab.33.cn/chain33/chain33/system"
	"gitlab.33.cn/chain33/chain33/util/cli"
	_ "gitlab.33.cn/chain33/plugin/plugin"
	"gitlab.33.cn/chain33/plugin/cli/buildflags"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName)

}