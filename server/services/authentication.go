package services

import (
	"context"
	parkingpb "parky/proto"
	"parky/server/db"
	"parky/server/models"

	"github.com/zebresel-com/mongodm"
	"gopkg.in/mgo.v2/bson"
)

type AuthenticationServer struct {
	parkingpb.UnimplementedAuthenticationServer
}

func (s *AuthenticationServer) AuthenticateUser(ctx context.Context, in *parkingpb.AuthenticateUserRequest) (*parkingpb.AuthenticateUserResponse, error) {
	aunthenticationdb := db.ConnectToAuthenticationDB()
	defer aunthenticationdb.Close()
	Users := aunthenticationdb.Model("User")
	user := &models.User{}
	err := Users.FindOne(bson.M{"username": in.Username}).Exec(user)

	if _, ok := err.(*mongodm.NotFoundError); ok {
		//no records were found
		Users.New(user) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

		user.UserName = in.GetUsername()

		err1 := user.Save()
		if err1 != nil {
			return nil, err1
		}
	} else if err != nil {
		return nil, err
	}
	return &parkingpb.AuthenticateUserResponse{UserId: user.Id.Hex()}, nil
}
