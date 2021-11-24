package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Failure struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func APIResponse(code int, message string, data interface{}) Response {
	var response = Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return response
}

func APIFailure(code int, message string, err interface{}) Failure {
	var failure = Failure{
		Code:    code,
		Message: message,
		Error:   err,
	}
	return failure
}
