package domain

import (
	"github.com/yulPa/yulmails/pkg/mail"
	"github.com/yulPa/yulmails/pkg/mongo"
)

const (
	BUFFER_DATABASE_NAME = "buffer"
)

func GetMailToSend() ([]mail.MailEntry, error) {
	/*
	   This domain object will return mail ready to be sent
	*/
	var mg = mongo.NewSession("mongodb://workdb:27017")
	return mg.DB(BUFFER_DATABASE_NAME).GetSendableMails()
}

func GetMailToCompute() ([]mail.MailEntry, error) {
	/*
	   This domain object will return mail ready to be sent
	*/
	var mg = mongo.NewSession("mongodb://workdb:27017")
	return mg.DB(BUFFER_DATABASE_NAME).GetMailToCompute()
}
