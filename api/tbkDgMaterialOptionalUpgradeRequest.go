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

type TbkDgMaterialOptionalUpgradeRequest struct {
	StartDsr          string `json:"start_dsr,omitempty"`
	PageSize          string `json:"page_size,omitempty"`
	PageNo            string `json:"page_no,omitempty"`
	Platform          string `json:"platform,omitempty"`
	EndTkRate         string `json:"end_tk_rate,omitempty"`
	StartTkRate       string `json:"start_tk_rate,omitempty"`
	EndPrice          string `json:"end_price,omitempty"`
	StartPrice        string `json:"start_price,omitempty"`
	IsOverseas        string `json:"is_overseas,omitempty"`
	IsTmall           string `json:"is_tmall,omitempty"`
	Sort              string `json:"sort,omitempty"`
	Itemloc           string `json:"itemloc,omitempty"`
	Cat               string `json:"cat,omitempty"`
	Q                 string `json:"q,omitempty"`
	MaterialId        string `json:"material_id,omitempty"`
	HasCoupon         string `json:"has_coupon,omitempty"`
	Ip                string `json:"ip,omitempty"`
	AdzoneId          string `json:"adzone_id,omitempty"`
	NeedFreeShipment  string `json:"need_free_shipment,omitempty"`
	NeedPrepay        string `json:"need_prepay,omitempty"`
	IncludePayRate30  string `json:"include_pay_rate_30,omitempty"`
	IncludeGoodRate   string `json:"include_good_rate,omitempty"`
	IncludeRfdRate    string `json:"include_rfd_rate,omitempty"`
	NpxLevel          string `json:"npx_level,omitempty"`
	EndKaTkRate       string `json:"end_ka_tk_rate,omitempty"`
	StartKaTkRate     string `json:"start_ka_tk_rate,omitempty"`
	DeviceEncrypt     string `json:"device_encrypt,omitempty"`
	DeviceValue       string `json:"device_value,omitempty"`
	DeviceType        string `json:"device_type,omitempty"`
	LockRateEndTime   string `json:"lock_rate_end_time,omitempty"`
	LockRateStartTime string `json:"lock_rate_start_time,omitempty"`
	Longitude         string `json:"longitude,omitempty"`
	Latitude          string `json:"latitude,omitempty"`
	CityCode          string `json:"city_code,omitempty"`
	SellerIds         string `json:"seller_ids,omitempty"`
	SpecialId         string `json:"special_id,omitempty"`
	RelationId        string `json:"relation_id,omitempty"`
	PageResultKey     string `json:"page_result_key,omitempty"`
	UcrowdId          string `json:"ucrowd_id,omitempty"`
	UcrowdRankItems   string `json:"ucrowd_rank_items,omitempty"`
	GetTopnRate       string `json:"get_topn_rate,omitempty"`
}

func (t *TbkDgMaterialOptionalUpgradeRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.material.optional.upgrade")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
