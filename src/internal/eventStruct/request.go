package eventStruct

type RequestContent struct {
	Opt 		int `json:"opt"`
	Content 	string `json:"content"`
}

type RespContent struct {
	Origin     	string `json:"origin"`
	Base64     	string `json:"base64"`
	Html       	string `json:"html"`
	Unicode    	string `json:"unicode"`
	Url  		string `json:"url"`
	Utf8ToGbk   string `json:"utf8ToGbk"`
}