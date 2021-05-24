package domain

import (
	portsServer "ports/internal/pkg/pb/server"
)

type DomainApp struct {
	PortsServer portsServer.PortsServiceServer
}

func (app DomainApp) Run() {
	portsServer.Start(portsServer.NewPortsServiceServer())
}
