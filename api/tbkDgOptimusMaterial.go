/*
@author:Joker
@file: tbkDgOptimusMaterial.go
@time: 2021/7/12/11:10 上午
@blog: https://github.com/joker-heyra
@description: https://open.taobao.com/api.htm?docId=33947&docType=2&scopeId=16518#requestParams

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgOptimusMaterial struct {
	PageSize      string `json:"page_size"`
	AdZoneId      string `json:"adzone_id"`
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

const MethodName = "taobao.tbk.dg.optimus.material"

func (t *TbkDgOptimusMaterial) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, MethodName)
	fMaps := convert([]map[string]string{bMap, pMap})

	sign := getSign(fMaps, rest.AppSecret)

	return request(fMaps, sign, rest.RestUrl)

}
