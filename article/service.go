package article

//Service struct
type Service struct {
	repo *repository
}

//StartArticleService def
func StartArticleService(repo *repository) *Service {
	return &Service{
		repo: repo,
	}
}
