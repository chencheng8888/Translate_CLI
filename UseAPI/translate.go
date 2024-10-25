package useapi

import (
	"encoding/json"
	"fmt"
	"translate_cli/UseAPI/utils"
	"translate_cli/UseAPI/utils/authv3"
	"translate_cli/conf"
)
type AutoGenerated struct {
	TSpeakURL string `json:"tSpeakUrl"`
	RequestID string `json:"requestId"`
	Query string `json:"query"`
	IsDomainSupport bool `json:"isDomainSupport"`
	Translation []string `json:"translation"`
	MTerminalDict MTerminalDict `json:"mTerminalDict"`
	ErrorCode string `json:"errorCode"`
	Dict Dict `json:"dict"`
	Webdict Webdict `json:"webdict"`
	L string `json:"l"`
	IsWord bool `json:"isWord"`
	SpeakURL string `json:"speakUrl"`
}
type MTerminalDict struct {
	URL string `json:"url"`
}
type Dict struct {
	URL string `json:"url"`
}
type Webdict struct {
	URL string `json:"url"`
}
func Translate(txt, from, to string) ([]string, error) {
	var resp AutoGenerated
	paramsMap := createRequestParams(txt, from, to)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(conf.Global_App.AppKey, conf.Global_App.AppSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/api", header, paramsMap, "application/json")
	err := json.Unmarshal(result,&resp)
	if err!=nil {
		return nil,err
	}
	if resp.ErrorCode != "0" {
		return nil,fmt.Errorf("Request error, error code is "+ resp.ErrorCode)
	}
	return resp.Translation,nil
}

func createRequestParams(q, from, to string) map[string][]string {
	vocabId := "out_id"
	return map[string][]string{
		"q":       {q},
		"from":    {from},
		"to":      {to},
		"vocabId": {vocabId},
	}
}