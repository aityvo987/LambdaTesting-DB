package pkg

import (
	"bytes"
	"encoding/json"
	entity "lambda-test/internal/entity"
	"log"
	"time"
)

type HttpResponse struct {
	ChannelID string
	Uuid      string
	Err       error
	Time      time.Time
	Data      interface{}
}

func bodyResponse(respBody entity.AppResponse) []byte {
	mapResp := make(map[string]interface{})
	mapResp["responseId"] = respBody.RequestID
	mapResp["responseTime"] = respBody.CreatedTime
	mapResp["responseError"] = respBody.Error
	// switch bodyInput.ChannelID {
	// // case string(common.ChannelC2c):
	// // 	mapResp["responseCode"] = common.ParseError(bodyInput.Err).Code()
	// // 	mapResp["responseMessage"] = common.ParseError(bodyInput.Err).MessageC2c()
	// // default:
	// // 	mapResp["responseCode"] = common.ParseError(bodyInput.Err).Code()
	// // 	mapResp["responseMessage"] = common.ParseError(bodyInput.Err).Message()
	// // }

	body, errMarshal := json.Marshal(mapResp)
	if errMarshal != nil {
		log.Default().Println("marshal response err", errMarshal)
	}
	return body
}

// build response data, cast error code, message resposne
func Response(respBody entity.AppResponse) string {
	respBody.CreatedTime = time.Now()
	var buf bytes.Buffer
	if respBody.Error != nil {
		body := bodyResponse(entity.AppResponse{
			ChannelID:   respBody.ChannelID,
			RequestID:   respBody.RequestID,
			Error:       respBody.Error,
			CreatedTime: respBody.CreatedTime,
		})
		json.HTMLEscape(&buf, body)
		return buf.String()
	}
	mapRes := map[string]interface{}{
		"responseId":      respBody.RequestID,
		"responseCode":    "00",
		"responseMessage": "successfully",
		"responseTime":    respBody.CreatedTime,
	}
	if respBody.Data != nil {
		mapRes["data"] = respBody.Data
	}
	body, errMarshal := json.Marshal(mapRes)
	if errMarshal != nil {
		log.Default().Println("marshal response err", errMarshal)
	}
	json.HTMLEscape(&buf, body)
	return buf.String()
}
