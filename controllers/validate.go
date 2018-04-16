package controllers

import (
	"log"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValiController struct {
	beego.Controller
}

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名
type Account struct {
    Name   string `valid:"Required;" form:"name"` // Name 不能为空
    Age    int    `valid:"Range(1, 100)" form:"age"` // 1 <= Age <= 100，超出此范围即为不合法
    Email  string `valid:"Email; MaxSize(100)" form:"email"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
    Mobile string `valid:"Mobile" form:"mobile"` // Mobile 必须为正确的手机号
}

//函数验证
// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *Account) Valid(v *validation.Validation) {
    if strings.Index(u.Name, "admin") != -1 {	//查找是否有amdin这个字符串
        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
        v.SetError("Name", "名称里不能含有admin")
    }
}


func (c *ValiController) Post(){
	valid:=validation.Validation{}
	//取得参数
	account:=Account{}
	c.ParseForm(&account)
	//验证
	b,_:=valid.Valid(&account)
	var result string
	if b!=true{
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			result=result+err.Key+":"+err.Message+" "
        }
	}
	c.Ctx.Output.Body([]byte(result))
}