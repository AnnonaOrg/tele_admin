package api_handler

import (
	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/umfaka/umfaka_core/handler"
	"github.com/umfaka/umfaka_core/internal/constvar"
)

// 404 Not found
func ApiNotFound(c *gin.Context) {
	handler.SendResultResponse(c, errno.Err404, constvar.APPDesc404())
}

// API Hello
func ApiHello(c *gin.Context) {
	handler.SendResultResponse(c, errno.SayHello, constvar.APPDesc())
	// indexSite := config.GetSiteConfig().RootWebURL
	// handler.SendRedirect(c, indexSite)
}

// ping
func ApiPing(c *gin.Context) {
	handler.SendResultResponse(c, errno.PONG, constvar.APPVersion())
}
