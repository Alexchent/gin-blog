package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	setting "gin-blog/pkg/settting"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.ERROR_NOT_EXIST_ARTICLE
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context)  {
	title := c.Query("title")
	tagId := com.StrTo(c.DefaultQuery("tagId", "-1")).MustInt()

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}

	if tagId > 0 {
		maps["tag_id"] = tagId
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// AddArticle 新建文章
func AddArticle(c *gin.Context) {

}

// EditArticle 修改文章
func EditArticle(c *gin.Context)  {

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.ERROR_NOT_EXIST_ARTICLE
	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}