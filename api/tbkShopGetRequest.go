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
	Fields              string `json:"fields,omitempty"`
	Q                   string `json:"q,omitempty"`
	Sort                string `json:"sort,omitempty"`
	IsTmall             string `json:"is_tmall,omitempty"`
	StartCredit         string `json:"start_credit,omitempty"`
	EndCredit           string `json:"end_credit,omitempty"`
	StartCommissionRate string `json:"start_commission_rate,omitempty"`
	EndCommissionRate   string `json:"end_commission_rate,omitempty"`
	StartTotalAction    string `json:"start_total_action,omitempty"`
	EndTotalAction      string `json:"end_total_action,omitempty"`
	StartAuctionCount   string `json:"start_auction_count,omitempty"`
	EndAuctionCount     string `json:"end_auction_count,omitempty"`
	Platform            string `json:"platform,omitempty"`
	PageNo              string `json:"page_no,omitempty"`
	PageSize            string `json:"page_size,omitempty"`
}

func (t *TbkShopGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.shop.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
