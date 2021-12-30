package routers

import (
	"gin-blog/middleware"
	setting "gin-blog/pkg/settting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.GET("/articles", v1.GetArticles)
		//新建标签
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定标签
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定标签
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	//使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})


	// HTML渲染 https://www.kancloud.cn/oldlei/go-gin/1333093
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	return r
}
