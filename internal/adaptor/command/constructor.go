package command

import "github.com/Goboolean/manager-cli/internal/port/in"

type CommandAdaptor struct {
	backUpService in.BackupCmdPort
	regService    in.RegCmdPort
	statusService in.StatusCmdPort
}

func New(backUpCmdPort in.BackupCmdPort, regCmdPort in.RegCmdPort, statusCmdPort in.StatusCmdPort) *CommandAdaptor {
	return &CommandAdaptor{
		backUpService: backUpCmdPort,
		regService:    regCmdPort,
		statusService: statusCmdPort,
	}
}
