package respons

type respons struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Value   string `json:"value"`
}

func ResponsValue(message string, code int, value string, data interface{}) respons {
	meta := meta{
		Message: message,
		Code:    code,
		Value:   value,
	}

	respons := respons{
		Meta: meta,
		Data: data,
	}

	return respons
}
