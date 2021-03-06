package requirements

import (
	"cf"
	"cf/api"
	"cf/net"
	"cf/terminal"
)

type BuildpackRequirement interface {
	Requirement
	GetBuildpack() cf.Buildpack
}

type buildpackApiRequirement struct {
	name          string
	ui            terminal.UI
	buildpackRepo api.BuildpackRepository
	buildpack     cf.Buildpack
}

func newBuildpackRequirement(name string, ui terminal.UI, bR api.BuildpackRepository) (req *buildpackApiRequirement) {
	req = new(buildpackApiRequirement)
	req.name = name
	req.ui = ui
	req.buildpackRepo = bR
	return
}

func (req *buildpackApiRequirement) Execute() (success bool) {
	var apiResponse net.ApiResponse
	req.buildpack, apiResponse = req.buildpackRepo.FindByName(req.name)

	if apiResponse.IsNotSuccessful() {
		req.ui.Failed(apiResponse.Message)
		return false
	}

	return true
}

func (req *buildpackApiRequirement) GetBuildpack() cf.Buildpack {
	return req.buildpack
}
