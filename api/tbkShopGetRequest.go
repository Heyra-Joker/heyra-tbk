/*
@author:Joker
@file: tbkShopGetRequest.go
@time: 2021/7/13/10:03 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkShopGetRequest struct {
	Fields              string `json:"fields"`
	Q                   string `json:"q"`
	Sort                string `json:"sort"`
	IsTmall             string `json:"is_tmall"`
	StartCredit         string `json:"start_credit"`
	EndCredit           string `json:"end_credit"`
	StartCommissionRate string `json:"start_commission_rate"`
	EndCommissionRate   string `json:"end_commission_rate"`
	StartTotalAction    string `json:"start_total_action"`
	EndTotalAction      string `json:"end_total_action"`
	StartAuctionCount   string `json:"start_auction_count"`
	EndAuctionCount     string `json:"end_auction_count"`
	Platform            string `json:"platform"`
	PageNo              string `json:"page_no"`
	PageSize            string `json:"page_size"`
}

func (t *TbkShopGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.shop.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
