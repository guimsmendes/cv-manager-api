package mongo

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

func Connect(Host string, Username string, Password string, Database string) *mgo.Database {
	mongoDbDialInfo := &mgo.DialInfo {
		Addrs: []string{Host},
		Timeout: 60 * time.Second,
		Username : Username,
		Password : Password,
		Database : Database,
	}
	session, err := mgo.DialWithInfo(mongoDbDialInfo)

	if err != nil {
		log.Fatal(err)
	}

	return session.DB(Database)
}