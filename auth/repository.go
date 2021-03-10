package auth

import (
	"fmt"

	"github.com/fahimsGit/restAPI/configuration"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type authRepository struct {
	dbSession *mgo.Session
	dbName    string
	tableName string
}
type repositoryInterface interface {
	checklogin(userAuth UserAuthentication) (Response, error)
}

func (r *authRepository) checkLogin(userAuth UserAuthentication) (Response, error) {
	var authRes Response
	fmt.Println(userAuth.UserName)
	fmt.Println(r.dbName)
	fmt.Println(r.tableName)

	coll := r.dbSession.DB(r.dbName).C(r.tableName)

	err := coll.Find(bson.M{"userName": userAuth.UserName, "pass": userAuth.Pass}).One(&userAuth)
	if err != nil {

		authRes.IsLog = false
		authRes.Token = "Not Found"
		return authRes, err
	}

	authRes.IsLog = true
	authRes.Token = "ActiveToken"
	return authRes, nil
}

func startAuthRepositoryService(dbSession *mgo.Session) *authRepository {
	return &authRepository{
		dbSession: dbSession,
		dbName:    configuration.DbName,
		tableName: configuration.AuthTable,
	}
}
