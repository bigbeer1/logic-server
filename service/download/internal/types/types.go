// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type Request struct {
	FileType string `path:"file_type"`
	File     string `path:"file"`
}

type RequestWav struct {
	File string `path:"file"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseByte struct {
	Data []byte `json:"data"`
}
