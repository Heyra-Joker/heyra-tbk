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
	CurrentPage      string `json:"current_page"`
	PageSize         string `json:"page_size"`
	Pid              string `json:"pid"`
	Postage          string `json:"postage"`
	Status           string `json:"status"`
	TaobaoCategoryId string `json:"taobao_category_id"`
	Word             string `json:"word"`
}

func (j *JuItemsSearchRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(j)
	pMap := convertPublicMap(rest, "taobao.ju.items.search")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
