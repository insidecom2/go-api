package response

import "demoecho/pkg/constants"

type ResponseOK struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseErr struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

type ResponseReqErr struct {
	Error string `json:"error"`
	Message []string `json:"message"`
}

func ResponseFail(message string ) ResponseErr{
	 m := ResponseErr{
		Error: constants.ErrMsg["INTERVAL_ERROR"],
		Message: message,
	 }

	 return m
}

func ResponseUnAuth(message string ) ResponseErr{
	 m := ResponseErr{
		Error: constants.ErrMsg["UNAUTHORIZE"],
		Message: message,
	 }

	 return m
}

func ResponseReqFail(message []string ) ResponseReqErr{
	 m := ResponseReqErr{
		Error: constants.ErrMsg["INVALID_REQUEST"],
		Message: message,
	 }

	 return m
}

func ResponseSuccess(message string, data interface{}) ResponseOK{
	 m := ResponseOK{
		Message: message,
		Data: data,
	 }
	
	 return m
}