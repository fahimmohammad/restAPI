package article

import (
	"errors"

	"github.com/google/uuid"
)

type serviceInterface interface {
	PostTodo(article Article) (Article, error)
	GetArticle(articleID string) (Article, error)
}

//Service struct
type Service struct {
	repo *repository
}

// PostTodo defination
func (articleService *Service) PostTodo(article Article) (Article, error) {
	article.ID = generateUUID()

	postResult, err := articleService.repo.createArticle(article)
	if err != nil {
		return Article{}, errors.New("cannot create")
	}
	return postResult, nil
}

// GetArticle definition
func (articleService *Service) GetArticle(articleID string) (Article, error) {

	getResult, err := articleService.repo.readSingleArticle(articleID)
	if err != nil {
		return Article{}, errors.New("cannot get")
	}
	return getResult, nil
}

//StartArticleService def
func StartArticleService(repo *repository) *Service {
	return &Service{
		repo: repo,
	}
}
func generateUUID() string {
	return uuid.New().String()
}
