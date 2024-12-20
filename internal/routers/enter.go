package routers

import (
	"github.com/QuocHuannn/Go-to-Work/internal/routers/manager"
	"github.com/QuocHuannn/Go-to-Work/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
