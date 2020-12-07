package database

import "gopkg.in/mgo.v2"

//Database Session
var DbSession *mgo.Session

//Model Declare
var Users *mgo.Collection

//Initiate Database
func Init(uri, dbname string) error {
	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}

	// See https://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	// Expose session and models
	DbSession = session
	Users = session.DB(dbname).C("Users")

	return nil
}
