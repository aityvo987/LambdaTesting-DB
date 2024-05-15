package pkg

import (
	"bytes"
	"encoding/json"
	entity "lambda-test/internal/entity"
	"log"
	"time"
)

func bodyResponse(respBody entity.AppResponse) []byte {
	mapResp := make(map[string]interface{})
	mapResp["responseId"] = respBody.RequestID
	mapResp["responseTime"] = respBody.CreatedTime
	mapResp["responseCode"] = ParseError(respBody.Error).Code()
	mapResp["responseMessage"] = ParseError(respBody.Error).Message()

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
		"responseId":   respBody.RequestID,
		"responseTime": respBody.CreatedTime,
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
