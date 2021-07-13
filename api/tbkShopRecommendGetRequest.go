/*
@author:Joker
@file: tbkShopRecommendGetRequest.go
@time: 2021/7/13/9:49 上午
@blog: https://github.com/joker-heyra
@description: https://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=16292

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkShopRecommendGetRequest struct {
	Fields   string `json:"fields"`
	UserId   string `json:"user_id"`
	Count    string `json:"count"`
	Platform string `json:"platform"`
}

func (t *TbkShopRecommendGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.shop.recommend.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
