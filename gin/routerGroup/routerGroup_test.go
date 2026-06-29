package routerGroup

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandleRouterGroup(t *testing.T) {
	router := gin.Default()
	HandleRouterGroup(router)
}
