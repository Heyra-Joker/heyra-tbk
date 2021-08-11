/*
@author:Joker
@file: tbkDgOptimusPromotionRequest.go
@time: 2021/7/13/11:23 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgOptimusPromotionRequest struct {
	PageSize    string `json:"page_size,omitempty"`
	PageNum     string `json:"page_num,omitempty"`
	AdzoneId    string `json:"adzone_id,omitempty"`
	PromotionId string `json:"promotion_id,omitempty"`
}

func (t *TbkDgOptimusPromotionRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.optimus.promotion")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)

}
