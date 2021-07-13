/*
@author:Joker
@file: tbkDgOptimusMaterialRequest.go
@time: 2021/7/13/11:13 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgOptimusMaterialRequest struct {
	PageSize      string `json:"page_size"`
	AdzoneId      string `json:"adzone_id"`
	PageNo        string `json:"page_no"`
	MaterialId    string `json:"material_id"`
	DeviceValue   string `json:"device_value"`
	DeviceEncrypt string `json:"device_encrypt"`
	DeviceType    string `json:"device_type"`
	ContentId     string `json:"content_id"`
	ContentSource string `json:"content_source"`
	ItemId        string `json:"item_id"`
	FavoritesId   string `json:"favorites_id"`
}

func (t *TbkDgOptimusMaterialRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.optimus.material")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
