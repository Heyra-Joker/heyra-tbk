/*
@author:Joker
@file: tbkActivityInfoGetRequest.go
@time: 2021/7/13/11:02 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkActivityInfoGetRequest struct {
	AdzoneId           string `json:"adzone_id,omitempty"`
	SubPid             string `json:"sub_pid,omitempty"`
	RelationId         string `json:"relation_id,omitempty"`
	ActivityMaterialId string `json:"activity_material_id,omitempty"`
	UnionId            string `json:"union_id,omitempty"`
}

func (t *TbkActivityInfoGetRequest) GetResponse(rest Rest) (string, error) {
	bMap := convertBusinessMap(t)
	pMap := convertPublicMap(rest, "taobao.tbk.activity.info.get")
	fMaps := convert([]map[string]string{bMap, pMap})
	sign := getSign(fMaps, rest.AppSecret)
	return request(fMaps, sign, rest.RestUrl)
}
