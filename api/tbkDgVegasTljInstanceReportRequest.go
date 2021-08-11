/*
@author:Joker
@file: tbkDgVegasTljInstanceReportRequest.go
@time: 2021/7/13/10:53 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgVegasTljInstanceReportRequest struct {
	RightsId string `json:"rights_id,omitempty"`
}

func (t *TbkDgVegasTljInstanceReportRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.vegas.tlj.instance.report")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
