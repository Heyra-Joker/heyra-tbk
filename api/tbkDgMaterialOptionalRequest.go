/*
@author:Joker
@file: tbkDgMaterialOptionalRequest.go
@time: 2021/7/13/10:09 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgMaterialOptionalRequest struct {
	StartDsr          string `json:"start_dsr"`
	PageSize          string `json:"page_size"`
	PageNo            string `json:"page_no"`
	Platform          string `json:"platform"`
	EndTkRate         string `json:"end_tk_rate"`
	StartTkRate       string `json:"start_tk_rate"`
	EndPrice          string `json:"end_price"`
	StartPrice        string `json:"start_price"`
	IsOverseas        string `json:"is_overseas"`
	IsTmall           string `json:"is_tmall"`
	Sort              string `json:"sort"`
	Itemloc           string `json:"itemloc"`
	Cat               string `json:"cat"`
	Q                 string `json:"q"`
	MaterialId        string `json:"material_id"`
	HasCoupon         string `json:"has_coupon"`
	Ip                string `json:"ip"`
	AdzoneId          string `json:"adzone_id"`
	NeedFreeShipment  string `json:"need_free_shipment"`
	NeedPrepay        string `json:"need_prepay"`
	IncludePayRate30  string `json:"include_pay_rate_30"`
	IncludeGoodRate   string `json:"include_good_rate"`
	IncludeRfdRate    string `json:"include_rfd_rate"`
	NpxLevel          string `json:"npx_level"`
	EndKaTkRate       string `json:"end_ka_tk_rate"`
	StartKaTkRate     string `json:"start_ka_tk_rate"`
	DeviceEncrypt     string `json:"device_encrypt"`
	DeviceValue       string `json:"device_value"`
	DeviceType        string `json:"device_type"`
	LockRateEndTime   string `json:"lock_rate_end_time"`
	LockRateStartTime string `json:"lock_rate_start_time"`
	Longitude         string `json:"longitude"`
	Latitude          string `json:"latitude"`
	CityCode          string `json:"city_code"`
	SellerIds         string `json:"seller_ids"`
	SpecialId         string `json:"special_id"`
	RelationId        string `json:"relation_id"`
	PageResultKey     string `json:"page_result_key"`
	UcrowdId          string `json:"ucrowd_id"`
	UcrowdRankItems   string `json:"ucrowd_rank_items"`
	GetTopnRate       string `json:"get_topn_rate"`
}

func (t *TbkDgMaterialOptionalRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.material.optional")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
