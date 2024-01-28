package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SendMsgHandler 消息发送管理
func SendMsgHandler(w http.ResponseWriter, r *http.Request) {
	// 读取内容
	content, err := io.ReadAll(r.Body)
	if err != nil {
		// 处理错误
		fmt.Println("读取错误:", err)
		return
	}

	res := &JsonResult{}
	res.Code = 0
	res.Data = "成功接收到消息" + string(content)
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Println("接收到消息:", string(content))
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
