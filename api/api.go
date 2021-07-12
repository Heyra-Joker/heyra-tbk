/*
@author:Joker
@file: api.go
@time: 2021/7/12/11:10 上午
@blog: https://github.com/joker-heyra
@description: tbk api 公共参数 https://open.taobao.com/doc.htm?docId=101617&docType=1

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type Rest struct {
	RestUrl    string
	Method     string
	AppKey     string
	Session    string
	Timestamp  string
	V          string
	SignMethod string
	Sign       string
	Format     string
	Simplify   bool
}
