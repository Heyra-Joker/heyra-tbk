/*
@author:Joker
@file: api.go
@time: 2021/7/12/11:10 上午
@blog: https://github.com/joker-heyra
@description: tbk api 公共参数 https://open.taobao.com/doc.htm?docId=101617&docType=1

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/
package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type businessMap map[string]string
type publicMap map[string]string
type finallyMap map[string]string

type Rest struct {
	Simplify  bool
	RestUrl   string
	AppKey    string
	AppSecret string
	Session   string
	V         string
	Other     map[string]string // 其他公共参数，使用 key-value 形式
}

func getSign(fMaps finallyMap, appSecret string) string {
	sMaps := []string{}
	for k, _ := range fMaps {
		sMaps = append(sMaps, k)
	}

	signatureStr := appSecret
	for _, k := range sMaps {
		signatureStr += k + fMaps[k]
	}
	signatureStr += appSecret

	sign := getMD5(signatureStr)
	return sign
}

func convertBusinessMap(v interface{}) businessMap {
	var dataMap = make(businessMap)
	var bMap = make(businessMap)

	data, err := json.Marshal(v) // Convert to a json string

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &dataMap) // Convert to a map
	if err != nil {
		panic(err)
	}

	for k, v := range dataMap {
		if v != "" {
			bMap[k] = v
		}
	}
	return bMap
}

func convertPublicMap(rest Rest, method string) publicMap {
	var pMap = make(publicMap)
	session := rest.Session
	v := rest.V
	if v == "" {
		v = "2.0"
	}
	pMap["format"] = "json"
	pMap["app_key"] = rest.AppKey
	pMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	pMap["v"] = v
	pMap["sign_method"] = "md5"
	pMap["method"] = method
	if session != "" {
		pMap["session"] = rest.Session
	}

	if len(rest.Other) != 0 {
		for k, v := range rest.Other {
			pMap[k] = v
		}
	}

	return pMap
}

func convert(maps []map[string]string) finallyMap {
	fMaps := map[string]string{}

	for _, map_ := range maps {
		for k, v := range map_ {
			fMaps[k] = v
		}
	}
	return fMaps
}

func getMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func request(fMaps finallyMap, sign, restUrl string) (string, error) {
	urlParams := url.Values{}

	for k, v := range fMaps {
		urlParams.Add(k, v)
	}

	urlParams.Add("sign", sign)

	resp, err := http.Post(restUrl, "application/x-www-form-urlencoded", strings.NewReader(urlParams.Encode()))

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	return body, nil
}
