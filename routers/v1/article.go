package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取单个文章
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

//获取多个文章
func GetArticles(c *gin.Context)  {

}
//新建文章
func AddArticle(c *gin.Context) {

}
//修改文章
func EditArticle(c *gin.Context)  {

}

//删除文章
func DeleteArticle(c *gin.Context) {

}