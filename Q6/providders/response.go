package providders

type Response struct {
	HttpCode int
	Error    error
	Message  string
	Data     interface{}
	Paging   *Paging
}

type Paging struct {
	TotalSize int `json:"total_size"`
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
}

func (p *provider) NewResponse() *Response {

	return &Response{
		HttpCode: 1,
	}
}
