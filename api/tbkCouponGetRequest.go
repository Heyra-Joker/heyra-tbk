/*
@author:Joker
@file: tbkCouponGetRequest.go
@time: 2021/7/13/11:09 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkCouponGetRequest struct {
	Me         string `json:"me,omitempty"`
	ItemId     string `json:"item_id,omitempty"`
	ActivityId string `json:"activity_id,omitempty"`
}

func (t *TbkCouponGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.coupon.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
