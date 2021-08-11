/*
@author:Joker
@file: tbkTpwdCreateRequest.go
@time: 2021/7/13/11:04 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkTpwdCreateRequest struct {
	Url    string `json:"url,omitempty"`
	Text   string `json:"text,omitempty"`
	Logo   string `json:"logo,omitempty"`
	Ext    string `json:"ext,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

func (t *TbkTpwdCreateRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.tpwd.create")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
