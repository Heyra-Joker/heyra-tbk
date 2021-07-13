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
	AdzoneId             string `json:"adzone_id"`
	SecurityLevel        string `json:"security_level"`
	UseStartTime         string `json:"use_start_time"`
	UseEndTimeMode       string `json:"use_end_time_mode"`
	UseEndTime           string `json:"use_end_time"`
	SendEndTime          string `json:"send_end_time"`
	SendStartTime        string `json:"send_start_time"`
	PerFace              string `json:"per_face"`
	SecuritySwitch       string `json:"security_switch"`
	UserTotalWinNumLimit string `json:"user_total_win_num_limit"`
	Name                 string `json:"name"`
	TotalNum             string `json:"total_num"`
	ItemId               string `json:"item_id"`
	CampaignType         string `json:"campaign_type"`
}

func (t *TbkDgVegasTljCreateRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.vegas.tlj.create")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
