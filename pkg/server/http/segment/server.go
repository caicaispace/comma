package segment

import (
	"comma/pkg/library/core/e"
	"comma/pkg/library/core/t"

	httpServer "github.com/caicaispace/gohelper/server/http"

	service "comma/pkg/service/segment"

	"github.com/gin-gonic/gin"
)

type ParticipleForm struct {
	Title string `form:"words" valid:"Required;MaxSize(10000)"`
	Tags  string `form:"words" valid:"Required;MaxSize(10000)"`
}

// Segment
// @Summary Segment
// @Produce  json
// @Param words
// @Success 200 {object} app.ResponseData
// @Failure 500 {object} app.ResponseData
// @Router /v1/api/es-participle [post]
func Segment(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form ParticipleForm
	)

	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}

	s := service.GetInstance()
	wordMap, synMap, hypMap, orderArray := s.GetSegmentData(form.Title, form.Tags)
	// fmt.Println(*wordMap)
	// fmt.Println(*synMap)
	// fmt.Println(*hypMap)
	// fmt.Println(*orderArray)

	rspData := t.Map{
		"word":  wordMap,
		"syn":   synMap,
		"hyp":   hypMap,
		"order": orderArray,
	}
	ctx.Success(rspData, nil)
}

func httpServerStart() {
	// //if *loadWordType == "remote" {
	// core.Setup(
	// 	// router
	// 	func() *gin.Engine {
	// 		r := gin.New()
	// 		r.Use(gin.Logger())
	// 		r.Use(gin.Recovery())
	// 		apiV1 := r.Group("/v1/api")
	// 		{
	// 			apiV1.POST("/es-participle", controller.Segment)
	// 		}
	// 		return r
	// 	}(),
	// 	// before start handle
	// 	func(env string) {
	// 		if env == "dev" && strings.Contains(util.GetRootPath(), "participle") == false {
	// 			fmt.Println("------------------ 请在 app 目录下启动软件 ------------------")
	// 			os.Exit(0)
	// 		}
	// 	})
	// //}
	// //videoData, err := utils.QueryVideoData()
	// //if err != nil {
	// //	return
	// //}
	// //fmt.Println(videoData)
}
