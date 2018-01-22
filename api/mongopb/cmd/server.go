package main

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/yulPa/yulmails/api/mongopb/proto"
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"
)

var log = logger.GetLogger()

type mailServer struct {
	mongo.Session
}

func newMailServer(mgo mongo.Session) pb.MailServer {
	return mailServer{mgo}
}

func (s mailServer) ReadMails(req *pb.MailRequest, res pb.Mail_ReadMailsServer) error {

	sess := s.Copy()
	db := sess.DB("mails")
	mails, err := db.ReadMails(req.GetEntity(), req.GetEnvironment())

	if err != nil {
		log.Errorln(err)
	}

	for _, mail := range mails {
		if err = res.Send(&mail); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	l, _ := net.Listen("tcp", ":9090")
	session := mongo.NewSession("mongodb://database:27017")

	s := grpc.NewServer()
	pb.RegisterMailServer(s, newMailServer(session))

	s.Serve(l)
}
