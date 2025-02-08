package manager

import "github.com/gin-gonic/gin"
type ManagerRouterGroup struct {
	UserRouter
	AdminRouter
}

func (m ManagerRouterGroup) InitUserRouter(MainGroup *gin.RouterGroup) {
	panic("unimplemented")
}
