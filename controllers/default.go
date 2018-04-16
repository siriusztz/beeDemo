package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type DownloadControllor struct{
	beego.Controller
}


type user struct {
    Email  string `form:"email"`
	Age   string  `form:"age"`
}

func (c *MainController) Get() {
	// c.Data["Website"] = beego.AppConfig.String("mysqluser")
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "about.html"
	
	//首先获取app.conf中的section,
	data:=make(map[string]string)
	data,_=beego.AppConfig.GetSection("database")
	//通过“模式::配置参数名”或得特定模式下的参数
	devPort:=beego.AppConfig.String("dev::httpport")
	c.Ctx.WriteString(beego.AppConfig.String("postgresql")+","+data["mysqluser"]+","+devPort)
}

//普通的post请求
func (c *MainController) Post() {
	//对于放于POST的HTTP包体中的表单也使用GetString或得
	name:=c.GetString("name")
	pw:=c.GetString("password")
	//除了上面的方法外，我们还可以使用struct解析,字段后面的名字必须是表单字段的名字
	u := user{}
	c.ParseForm(&u)
	//也可以获取表单里到文件
	_,header,_:=c.GetFile("homework")
	c.Ctx.WriteString(name+";"+pw+"\n"+u.Email+u.Age+"\n"+header.Filename)
}

func (c *MainController) JSON(){
	type Mystruct struct{
		Name string
		Age string
	}
	mystruct := Mystruct{Name:"ztz",Age:"23"}
    c.Data["json"] = &mystruct
	c.ServeJSON()
}

//下载
func (c *DownloadControllor) Get(){
	c.Ctx.Output.Download("download1/12345.txt") 
}
