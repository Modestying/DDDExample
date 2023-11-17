package application

import "ddd_demo/domain/iface"

type BankApp struct {
	servers  []iface.IServer
	stopChan chan struct{}
}

func NewBankApp() *BankApp {
	return &BankApp{
		servers:  make([]iface.IServer, 2),
		stopChan: make(chan struct{}),
	}
}

func (app *BankApp) Run() {
	for _, server := range app.servers {
		go server.Start()
	}
	<-app.stopChan
}

func (app *BankApp) RegisterServers(servers []iface.IServer) {
	app.servers = servers
}
