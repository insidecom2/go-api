package response

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseFail(message string) Response{
	 m := Response{
		Status: "false",
		Message: message,
		Data: nil,
	 }

	 return m
}


func ResponseSuccess(message string, data interface{}) Response{
	 m := Response{
		Status: "true",
		Message: message,
		Data: data,
	 }
	
	 return m
}