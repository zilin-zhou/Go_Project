package template

import "log"

var Templates = make(map[string]*template)

type template struct {
	ID       string
	TempType string
	TempName string
	Content  string
}

func Register(tempName string, temp *template) {
	if temp == nil {
		log.Fatalln("短信对象不能为空！")
	}
	if _, ok := Templates[tempName]; ok {
		log.Fatalln("同一个模板不能多次添加！")
	}
	Templates[tempName] = temp
}
func init() {
	t := &template{
		ID:       "1804550",
		TempType: "普通短信",
		TempName: "验证码",
		Content:  "验证码为：{1}，您正在登录，若非本人操作，请勿泄露。",
	}
	Register(t.TempName, t)
	//可以实现多个模板注册
}
