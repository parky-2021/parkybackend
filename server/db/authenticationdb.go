package db

import (
	"parky/server/models"

	"github.com/golang/glog"
	"github.com/zebresel-com/mongodm"
)

func ConnectToAuthenticationDB() *mongodm.Connection {
	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{"127.0.0.1:27017"},
		DatabaseName:  "users",
		//DatabaseUser:     "admin",
		//DatabasePassword: "admin",
		// The option `DatabaseSource` is the database used to establish
		// credentials and privileges with a MongoDB server. Defaults to the value
		// of `DatabaseName`, if that is set, or "admin" otherwise.
		DatabaseSource: "admin",
		Locals:         localmaps["en-US"],
	}

	connection, err := mongodm.Connect(dbConfig)

	if err != nil {
		glog.Error("Database connection error: %v", err)
	}
	connection.Register(&models.User{}, "users")
	return connection
}
