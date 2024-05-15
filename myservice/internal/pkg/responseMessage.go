package pkg

type ResponseCode string

var ResponseMessage = map[string]string{
	"00": "successfully",
	"01": "record not found",
	"02": "database error",
	"03": "request error",
	"04": "response error",
	"05": "system error",
}

func (rc ResponseCode) Code() string {
	return string(rc)
}

func (rc ResponseCode) Message() string {
	if value, ok := ResponseMessage[rc.Code()]; ok {
		return value
	}
	return ""
}

func ParseError(err error) ResponseCode {
	return ResponseCode(err.Error())
}

const (
	Success   ResponseCode = "00"
	RecNoF    ResponseCode = "01"
	DbError   ResponseCode = "02"
	ReqError  ResponseCode = "03"
	RespError ResponseCode = "04"
	SystemErr ResponseCode = "05"
)
