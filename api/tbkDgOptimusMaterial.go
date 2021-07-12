/*
@author:Joker
@file: tbkDgOptimusMaterial.go
@time: 2021/7/12/11:10 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

type TbkDgOptimusMaterial struct {
	Rest
	PageSize      int
	AdZoneId      int
	PageNo        int
	MaterialId    int
	DeviceValue   string
	DeviceEncrypt string
	DeviceType    string
	ContentId     int
	ContentSource string
	ItemId        string
	FavoritesId   int
}

func (*TbkDgOptimusMaterial) GetResponse() {

}
