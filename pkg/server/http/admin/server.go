package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})
	v1 := r.Group("/v1")
	{
		// word
		v1.GET("/dict/word", WordList)
		v1.GET("/dict/word/:id", WordInfo)
		v1.GET("/dict/words/:ids", WordListByIds)
		v1.POST("/dict/word", WordCreate)
		v1.PUT("/dict/word/:id", WordUpdate)
		v1.DELETE("/dict/word/:id", WordDelete)
		v1.DELETE("/dict/word", WordMultipleDelete)
		// weight
		v1.GET("/dict/weight", WeightList)
		v1.POST("/dict/weight", WeightCreate)
		v1.PUT("/dict/weight/:id", WeightUpdate)
		v1.DELETE("/dict/weight/:id", WeightDelete)
		v1.DELETE("/dict/weight", WeightMultipleDelete)
		// banned
		v1.GET("/dict/banned", BannedList)
		v1.POST("/dict/banned", BannedCreate)
		v1.PUT("/dict/banned/:id", BannedUpdate)
		v1.DELETE("/dict/banned/:id", BannedDelete)
		v1.DELETE("/dict/banned", BannedMultipleDelete)
		// stop
		v1.GET("/dict/stop", StopList)
		v1.POST("/dict/stop", StopCreate)
		v1.PUT("/dict/stop/:id", StopUpdate)
		v1.DELETE("/dict/stop/:id", StopDelete)
		v1.DELETE("/dict/stop", StopMultipleDelete)
		// festival
		v1.GET("/dict/festival", FestivalList)
		v1.POST("/dict/festival", FestivalCreate)
		v1.PUT("/dict/festival/:id", FestivalUpdate)
		v1.DELETE("/dict/festival/:id", FestivalDelete)
		v1.DELETE("/dict/festival", FestivalMultipleDelete)
		// high_frequency
		v1.GET("/dict/high_frequency", HighFrequencyList)
		v1.POST("/dict/high_frequency", HighFrequencyCreate)
		v1.PUT("/dict/high_frequency/:id", HighFrequencyUpdate)
		v1.DELETE("/dict/high_frequency/:id", HighFrequencyDelete)
		v1.DELETE("/dict/high_frequency", HighFrequencyMultipleDelete)
		// hyponym
		v1.GET("/dict/hyponym", HyponymList)
		v1.POST("/dict/hyponym", HyponymCreate)
		v1.PUT("/dict/hyponym/:id", HyponymUpdate)
		v1.DELETE("/dict/hyponym/:id", HyponymDelete)
		v1.DELETE("/dict/hyponym", HyponymMultipleDelete)
		// project
		v1.GET("/dict/synonym", SynonymList)
		v1.POST("/dict/synonym", SynonymCreate)
		v1.PUT("/dict/synonym/:id", SynonymUpdate)
		v1.DELETE("/dict/synonym/:id", SynonymDelete)
		v1.DELETE("/dict/synonym", SynonymMultipleDelete)
		// project
		v1.GET("/dict/project", ProjectList)
		v1.POST("/dict/project", ProjectCreate)
		v1.PUT("/dict/project/:id", ProjectUpdate)
		v1.DELETE("/dict/project/:id", ProjectDelete)
		v1.DELETE("/dict/project", ProjectMultipleDelete)
	}
}
