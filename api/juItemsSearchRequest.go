/*
@author:Joker
@file: juItemsSearchRequest.go
@time: 2021/7/13/10:56 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type JuItemsSearchRequest struct {
	CurrentPage      string `json:"current_page,omitempty"`
	PageSize         string `json:"page_size,omitempty"`
	Pid              string `json:"pid,omitempty"`
	Postage          string `json:"postage,omitempty"`
	Status           string `json:"status,omitempty"`
	TaobaoCategoryId string `json:"taobao_category_id,omitempty"`
	Word             string `json:"word,omitempty"`
}

func (j *JuItemsSearchRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(j)
	pMap := convertPublicMap(rest, "taobao.ju.items.search")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
