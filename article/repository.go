package article

import (
	"github.com/fahimsGit/restAPI/configuration"
	"github.com/globalsign/mgo"
)

type repository struct {
	dbSession *mgo.Session
	dbName    string
	tableName string
}

type repositoryInterface interface {
	createArticle(article Article) (Article, error)
	//getAllArticle() ([]Article, error)
	//getArticleById(articleID string) (Article, error)
	//updateArticle(articleID string, article Article) (Article, error)
	//deleteArticle(articleID string) error
}

func (r *repository) createArticle(article Article) (Article, error) {
	coll := r.dbSession.DB(r.dbName).C(r.tableName)
	err := coll.Insert(&article)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

func startRepositoryService(dbSession *mgo.Session) *repository {
	return &repository{
		dbSession: dbSession,
		dbName:    configuration.DbName,
		tableName: configuration.TargetTable,
	}
}
