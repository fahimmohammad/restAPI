package article

import (
	"fmt"
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
	ArticleService *Service
}
type createArticleRequest struct {
	Article Article
}

type createArticleResponse struct {
	Article Article `json:"article"`
	Err     string  `json:"err"`
}

// CreateHTTPHandlers init
func CreateHTTPHandlers(router *gin.RouterGroup, articleService *Service) {

	h := handlerFields{
		ArticleService: articleService,
	}

	router.POST("article", h.createArticleHandler)
	router.GET("article/:id", h.getArticleHandler)
	//router.GET("article/:id", h.getArticleByIdHandler)
	//router.DELETE("article/:id", h.deleteArticleHandler)
	//router.PUT("article/:id", h.updateArticleHandler)
}

type readSingleTodoResponse struct {
	Article Article `json:"article"`
	Err     string  `json:"err"`
}

func (h *handlerFields) getArticleHandler(ctx *gin.Context) {

	articleID := ctx.Param("id")

	getResult, err := h.ArticleService.repo.readSingleArticle(articleID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, readSingleTodoResponse{
			Article: Article{},
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, readSingleTodoResponse{
		Article: getResult,
		Err:     "",
	})
	return

}
func (h *handlerFields) createArticleHandler(ctx *gin.Context) {
	fmt.Println("Inside Hwndler")
	var req createArticleRequest
	if err := ctx.ShouldBindJSON(&req.Article); err != nil {
		ctx.JSON(http.StatusInternalServerError, createArticleResponse{
			Article: Article{},
			Err:     "Error in data binding",
		})
		return
	}

	postResult, err := h.ArticleService.repo.createArticle(req.Article)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, createArticleResponse{
			Article: Article{},
			Err:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, createArticleResponse{
		Article: postResult,
		Err:     "",
	})
	return
}
