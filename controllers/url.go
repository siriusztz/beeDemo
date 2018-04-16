package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UrlController struct {
	beego.Controller
}

func (c *UrlController) Get(){
	//得到url后面接上的id参数,参数名id必须写在router里面
	id:=c.Ctx.Input.Param(":id")
	//获得?后的parameter,GetString实则是调用的c.Ctx.Input.Query(key)
	name := c.GetString("name")
	password := c.GetString("password")

	//session
	v := c.GetSession("1")
    if v == nil {
        c.SetSession("1", "sessionWord")
	} else{
		next:=fmt.Sprintf("%s%s",v,"+")
		c.SetSession("1", next)
	}
	
	value,_:=v.(string)
	c.Ctx.WriteString(id+name+password+value)
}
//自动路由方法
// url/login   调用 UrlController 中的 Login 方法
// 除了前缀两个 /:controller/:method 的匹配之外，剩下的 url,beego会帮你自动化解析为参数，保存在this.Ctx.Input.Params当中：
func (c *UrlController) Login(){
	params:=c.Ctx.Input.Params()
	c.Ctx.WriteString(params["0"]+params["1"])
}

//构建url
func (c *UrlController) GetUrl() {
    c.Ctx.Output.Body([]byte(c.URLFor(".GetUrl")))
}