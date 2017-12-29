package main

import (
	"github.com/quickfixgo/quickfix"
	"log"
)

//implement quickfix.Application interface
type FixApplication struct {
}

// OnCreate
func (app FixApplication) OnCreate(sessionID quickfix.SessionID) {
	log.Println("OnCreate("+sessionID.String()+")")
}

// OnLogon
func (app FixApplication) OnLogon(sessionID quickfix.SessionID) {
	log.Println("OnLogon("+sessionID.String()+")")
}

// OnLogout
func (app FixApplication) OnLogout(sessionID quickfix.SessionID) {
	log.Println("OnLogout("+sessionID.String()+")")
}

// FromAdmin
func (app FixApplication) FromAdmin(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	log.Println("FromAdmin("+sessionID.String()+") "+message.String())
	//
	return nil
}

// ToAdmin
func (app FixApplication) ToAdmin(message *quickfix.Message, sessionID quickfix.SessionID) {
	log.Println("ToAdmin("+sessionID.String()+") "+message.String())
}

// FromApp
func (app FixApplication) FromApp(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	log.Println("FromApp("+sessionID.String()+") "+message.String())
	//
	return nil
}

// ToApp
func (app FixApplication) ToApp(message *quickfix.Message, sessionID quickfix.SessionID) error {
	log.Println("ToApp("+sessionID.String()+") "+message.String())

	return nil
}


