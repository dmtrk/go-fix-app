package main

import (
	"os"
	"github.com/quickfixgo/quickfix"
	"os/signal"
	"strings"
	"log"
)

func main() {
	fileName := "config/acceptor.cfg"
	if len(os.Args)>1 {fileName = os.Args[1]}
	log.Println("Loading file: '"+fileName+"'")

	cfg, err := os.Open(fileName)
	if err != nil { panic(err) }

	settings, err := quickfix.ParseSettings(cfg)
	if err != nil { panic(err) }

	if isInitiator(settings) {//initiator
		startInitiator(settings)
	} else { // acceptor
		startAcceptor(settings)
	}
}

func startAcceptor(settings *quickfix.Settings) {
	log.Println("Starting acceptor...")
	app:= FixApplication {}
	var storeFactory = quickfix.NewFileStoreFactory(settings)
	logFactory, err := quickfix.NewFileLogFactory(settings)
	if err != nil { panic(err) }
	//
	acceptor, err := quickfix.NewAcceptor(app, storeFactory, settings, logFactory)
	if err != nil { panic(err) }
	err = acceptor.Start()
	if err != nil { panic(err) }
	//
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
	//
	acceptor.Stop()
}

func startInitiator(settings *quickfix.Settings) {
	log.Println("Starting initiator...")
	var app FixApplication = FixApplication{}
	var storeFactory = quickfix.NewFileStoreFactory(settings)
	logFactory, err := quickfix.NewFileLogFactory(settings)
	if err != nil { panic(err) }
	//
	initiator, err := quickfix.NewInitiator(app, storeFactory, settings, logFactory)
	if err != nil { panic(err) }
	err = initiator.Start()
	if err != nil { panic(err) }
	//
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
	//
	initiator.Stop()
}

func isInitiator(settings *quickfix.Settings) bool {
	if settings.GlobalSettings().HasSetting("ConnectionType") {
		var connectionType, _ = settings.GlobalSettings().Setting("ConnectionType")
		return strings.Compare(strings.ToLower(connectionType),"initiator")==0
	}
	return false
}