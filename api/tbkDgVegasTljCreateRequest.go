/*
@author:Joker
@file: tbkDgVegasTljCreateRequest.go
@time: 2021/7/13/10:47 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgVegasTljCreateRequest struct {
	AdzoneId             string `json:"adzone_id,omitempty"`
	SecurityLevel        string `json:"security_level,omitempty"`
	UseStartTime         string `json:"use_start_time,omitempty"`
	UseEndTimeMode       string `json:"use_end_time_mode,omitempty"`
	UseEndTime           string `json:"use_end_time,omitempty"`
	SendEndTime          string `json:"send_end_time,omitempty"`
	SendStartTime        string `json:"send_start_time,omitempty"`
	PerFace              string `json:"per_face,omitempty"`
	SecuritySwitch       string `json:"security_switch,omitempty"`
	UserTotalWinNumLimit string `json:"user_total_win_num_limit,omitempty"`
	Name                 string `json:"name,omitempty"`
	TotalNum             string `json:"total_num,omitempty"`
	ItemId               string `json:"item_id,omitempty"`
	CampaignType         string `json:"campaign_type,omitempty"`
}

func (t *TbkDgVegasTljCreateRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.vegas.tlj.create")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
