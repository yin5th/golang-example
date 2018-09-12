package v1

import (
	"dothis.top/example/gin_blog/models"
	"dothis.top/example/gin_blog/pkg/e"
	"dothis.top/example/gin_blog/pkg/logging"
	"dothis.top/example/gin_blog/pkg/setting"
	"dothis.top/example/gin_blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个文章【指定文章】
func GetArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()

	valid.Min(id, 1, "id").Message("ID必须是大于0的整数")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			code = e.SUCCESS
			data = models.GetArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var tagId int = -1
	if arg := c.PostForm("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须是大于0的整数")
	}

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只能是0或者1")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})

}

//新增文章
func AddArticle(c *gin.Context) {

}

//修改文章
func EditArticle(c *gin.Context) {

}

//删除文章
func DeleteArticle(c *gin.Context) {

}
