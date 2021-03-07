package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerInterface interface {
	createArticleHandler(ctx *gin.Context)
	getArticleHandler(ctx *gin.Context)
	getArticleByIdHandler(ctx *gin.Context)
	deleteArticleHandler(ctx *gin.Context)
	updateArticleHandler(ctx *gin.Context)
}

type handlerFields struct {
	articleService *Service
}
type createArticleRequest struct {
	article Article
}

type createArticleResponse struct {
	article Article `json:"article"`
	err     string  `json:"err"`
}

// CreateHTTPHandlers init
func CreateHTTPHandlers(router *gin.RouterGroup, articleService *Service) {
	h := handlerFields{
		articleService: articleService,
	}

	router.POST("article", h.createArticleHandler)
	//router.GET("article", h.getArticleHandler)
	//router.GET("article/:id", h.getArticleByIdHandler)
	//router.DELETE("article/:id", h.deleteArticleHandler)
	//router.PUT("article/:id", h.updateArticleHandler)
}

func (h *handlerFields) createArticleHandler(ctx *gin.Context) {
	var req createArticleRequest

	if err := ctx.ShouldBindJSON(&req.article); err != nil {
		ctx.JSON(http.StatusInternalServerError, createArticleResponse{
			article: Article{},
			err:     "Error in data binding",
		})
		return
	}

	postResult, err := h.articleService.repo.createArticle(req.article)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, createArticleResponse{
			article: Article{},
			err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, createArticleResponse{
		article: postResult,
		err:     "",
	})
	return
}
