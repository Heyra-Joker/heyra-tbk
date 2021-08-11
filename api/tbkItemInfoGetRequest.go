/*
@author:Joker
@file: tbkItemInfoGetRequest.go
@time: 2021/7/13/11:07 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkItemInfoGetRequest struct {
	NumIids  string `json:"num_iids,omitempty"`
	Platform string `json:"platform,omitempty"`
	Ip       string `json:"ip,omitempty"`
}

func (t *TbkItemInfoGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.item.info.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
