package utils

type Response struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string {
	return r.Info
}

type FinalResponse struct {
	Status string      `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}
