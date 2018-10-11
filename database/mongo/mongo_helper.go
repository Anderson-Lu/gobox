package mongo_helper 

import (
	mgo "gopkg.in/mgo.v2"
)

//connStr like mongodb://username:password@host:port/database
func  NewMongoClient(connStr string) (*mgo.Session,error) {
	session ,err := mgo.Dial(connStr)
	if err!= nil {
		session.SetMode(mgo.Monotonic, true)
	}
	return session,err
}