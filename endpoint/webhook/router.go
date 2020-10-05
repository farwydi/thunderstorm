package webhook

import (
	"github.com/farwydi/cleanwhale/tonic"
	"github.com/farwydi/thunderstorm/domain"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter(cfg domain.Config, fsm *domain.StateMachine) http.Handler {
	r := tonic.NewMix(cfg.Project.Mode, zap.L().Named("webhook"))

	v1 := r.Group("/v1", tonic.V("v1"))
	{
		v1.GET("/webhook", func(c *gin.Context) {
			c.Set("handler.url", "/webhook")

			var update tgbotapi.Update
			err := c.ShouldBindJSON(&update)
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			ctx := c.Request.Context()

			err = fsm.Tx(ctx, update)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			c.Status(http.StatusOK)
		})
	}

	return r
}
