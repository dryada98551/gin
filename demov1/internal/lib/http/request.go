package libhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	libmodel "main/internal/lib/model"
	libreturncode "main/internal/lib/returnCode"
	// jsoniter "github.com/json-iterator/go"
)

func Request(url string, headers http.Header, data interface{}) ([]byte, libmodel.SystemReturn) {
	var sysReturn libmodel.SystemReturn
	// 將數據編碼為JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.DataToJson.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.DataToJson.Message
		return nil, sysReturn // 處理錯誤
	}

	// 創建請求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.ApiReqCreate.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.ApiReqCreate.Message
		return nil, sysReturn // 處理錯誤
	}

	// 設置請求頭
	req.Header = headers

	// 執行請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.ApiReqDo.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.ApiReqDo.Message
		return nil, sysReturn // 處理錯誤
	}
	defer resp.Body.Close()

	// 讀取回應
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.ApiResRead.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.ApiResRead.Message
		return nil, sysReturn // 處理錯誤
	}
	
	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.Other.Message
	return body, sysReturn
}