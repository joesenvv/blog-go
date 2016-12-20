package Response

import (
	"fmt"
	"encoding/json"
	"net/http"
)

const (
	WARNING = "warning"
	SUCCESS = "success"
	ALERT = "alert"
	INFO = "info"
)

type Response struct {
	Code int
	Info string
	Data interface{}
	Err  Error
}
type Error struct {
	Level string
	Msg   string
}
type Success struct {
	Level string
	Msg   string
}

func NewResponse() *Response {
	return &Response{Code: 1}
}
func (resp *Response) Tips(msg, level string) {
	resp.Info = msg
	resp.Err = Error{Level:level, Msg:msg}
}
func (resp *Response) WriteJson(w http.ResponseWriter) {
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("helper.go line:55", err)
		w.Write([]byte(`{Status:-1,Err:Error{Level:"alert",Msg:"code=-1|序列化失败！"}}`))
	} else {
		w.Write(b)
	}
}
//成功
func (resp *Response) Success(msg string) {
	if msg == "" {
		msg = "操作成功"
	}
	resp.Code = 1
	resp.Info = msg
	resp.Data = Success{Level: SUCCESS, Msg: msg}
}
//错误信息
func (resp *Response) Error(msg string) {
	if msg == "" {
		msg = "操作失败"
	}
	resp.Code = 0
	resp.Info = msg
	resp.Err = Error{Msg:msg}
}
//扩展数据
func (resp *Response) SetData(Data interface{}) {
	resp.Data = Data
}