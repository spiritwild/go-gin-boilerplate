package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lincare/moria/config"
	"github.com/lincare/moria/controllers"
	"github.com/spf13/viper"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&UserSuite{})

type UserSuite struct {
	config *viper.Viper
	router *gin.Engine
}

func (s *MoriaSuite) SetUpTest(c *C) {
	config.Init("test")
	s.config = config.GetConfig()
	s.router = SetupRouter()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	health := new(controllers.HealthController)
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/", user.Signup)
			userGroup.DELETE("/:id", user.Delete)
			userGroup.PUT("/:id", user.Update)
		}
	}
	return router
}
