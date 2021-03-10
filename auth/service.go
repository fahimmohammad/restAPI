package auth

import "errors"

type serviceInterface interface {
	checkLogin(UserAuthentication) (Response, error)
}

//Service struct
type Service struct {
	repo *authRepository
}

func (loginService *Service) checkLogin(authUser UserAuthentication) (Response, error) {
	getResult, err := loginService.repo.checkLogin(authUser)
	if err != nil {
		return Response{}, errors.New("cannot get")
	}
	return getResult, nil
}

//StartArticleService def
func StartArticleService(repo *authRepository) *Service {
	return &Service{
		repo: repo,
	}
}
