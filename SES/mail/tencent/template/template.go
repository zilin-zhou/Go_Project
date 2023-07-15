package template

import "log"

var Templates map[string]*template

type template struct {
	ID       uint64
	TempName string
}

func Register(tempName string, temp *template) {
	if temp == nil {
		log.Fatalln("邮件对象不能为空！")
	}
	if _, ok := Templates[tempName]; ok {
		log.Fatalln("同一个模板不能多次添加！")
	}
	Templates[tempName] = temp
}
func init() {
	Templates = make(map[string]*template)
	t := &template{
		ID:       1804550,
		TempName: "邮件通知",
	}
	Register(t.TempName, t)
	//可以实现多个模板注册
}
