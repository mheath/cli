package commands

import (
	"cf/commands"
	"github.com/codegangsta/cli"
	testreq "testhelpers/requirements"
)

var CommandDidPassRequirements bool

func RunCommand(cmd commands.Command, ctxt *cli.Context, reqFactory *testreq.FakeReqFactory) {
	CommandDidPassRequirements = false

	reqs, err := cmd.GetRequirements(reqFactory, ctxt)
	if err != nil {
		return
	}

	for _, req := range reqs {
		success := req.Execute()
		if !success {
			return
		}
	}

	cmd.Run(ctxt)
	CommandDidPassRequirements = true

	return
}
