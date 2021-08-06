package v1

import (
	setting "gin-blog/pkg/settting"
	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"

	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := "admin"

	//参数验证
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	valid.MaxSize(name, 100, "name").Message("name最大长度为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只有0或1")
	code := e.SUCCESS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		code = 1000
	}

	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"message":e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}