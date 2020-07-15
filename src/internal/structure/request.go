package structure

type InputContent struct {
	Opt 		int `json:"opt"`
	Content 	string `json:"content"`
}

type RespContent struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            OutputContent     `json:"body"`
}

type RespHeader struct {
	ContentType   string `json:"Content-Type"`
}

type OutputContent struct {
	Origin     	string `json:"origin"`
	Base64     	string `json:"base64"`
	Html       	string `json:"html"`
	Unicode    	string `json:"unicode"`
	Url  		string `json:"url"`
	Utf8ToGbk   string `json:"utf8ToGbk"`
}