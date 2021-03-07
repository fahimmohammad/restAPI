package article

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// Init - initializes package

func Init(router *gin.RouterGroup, session *mgo.Session) {
	repository := startRepositoryService(session)
	service := StartArticleService(repository)
	CreateHTTPHandlers(router, service)
}
