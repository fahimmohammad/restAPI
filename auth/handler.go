package auth

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
	LoginService *Service
}
type loginRequest struct {
	UserAuth UserAuthentication
}

type loginResponse struct {
	AuthResponse Response
	Err          string `json:"err"`
}

// CreateHTTPHandlers init
func CreateHTTPHandlers(router *gin.RouterGroup, loginService *Service) {
	h := handlerFields{
		LoginService: loginService,
	}

	router.POST("login", h.loginHandler)
	//router.GET("article", h.getArticleHandler)
	//router.GET("article/:id", h.getArticleByIdHandler)
	//router.DELETE("article/:id", h.deleteArticleHandler)
	//router.PUT("article/:id", h.updateArticleHandler)
}

func (h *handlerFields) loginHandler(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req.UserAuth); err != nil {
		ctx.JSON(http.StatusInternalServerError, loginResponse{
			AuthResponse: Response{},
			Err:          "Error in data binding",
		})
		return
	}

	checkLogin, err := h.LoginService.checkLogin(req.UserAuth)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, loginResponse{
			AuthResponse: checkLogin,
			Err:          err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{
		AuthResponse: checkLogin,
		Err:          "",
	})
	return
}
