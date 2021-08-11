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
	PageSize      string `json:"page_size,omitempty"`
	AdZoneId      string `json:"adzone_id,omitempty"`
	PageNo        string `json:"page_no,omitempty"`
	MaterialId    string `json:"material_id,omitempty"`
	DeviceValue   string `json:"device_value,omitempty"`
	DeviceEncrypt string `json:"device_encrypt,omitempty"`
	DeviceType    string `json:"device_type,omitempty"`
	ContentId     string `json:"content_id,omitempty"`
	ContentSource string `json:"content_source,omitempty"`
	ItemId        string `json:"item_id,omitempty"`
	FavoritesId   string `json:"favorites_id,omitempty"`
}

func (t *TbkDgOptimusMaterial) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.dg.optimus.material")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)

}
