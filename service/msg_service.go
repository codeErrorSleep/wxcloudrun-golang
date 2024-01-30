package service

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

// start_ai_generated

type WeChatSendMsgReq struct {
	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime" json:"CreateTime"`
	MsgType      string `xml:"MsgType" json:"MsgType"`
	Content      string `xml:"Content" json:"Content"`
	MsgId        int64  `xml:"MsgId" json:"MsgId"`
	MsgDataId    string `xml:"MsgDataId" json:"MsgDataId"`
	Idx          string `xml:"Idx" json:"Idx"`
}

type WeChatReplyMsgRes struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	FromUserName string   `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime" json:"CreateTime"`
	MsgType      string   `xml:"MsgType" json:"MsgType"`
	Content      string   `xml:"Content" json:"Content"`
}

// end_ai_generated

// SendMsgHandler 消息发送管理
func SendMsgHandler(c *gin.Context) {
	// 接收请求
	// 读取请求体
	var wechatMsgReq WeChatSendMsgReq
	if err := c.ShouldBindJSON(&wechatMsgReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid wechatMsgReq" + err.Error(),
		})
		return
	}

	// 获取之前是否提交过 localCacheKey
	chatMsgReq, err := getMsgHistory(wechatMsgReq.FromUserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error" + err.Error(),
		})
		return
	}
	if chatMsgReq.Messages == nil {
		chatMsgReq = createChatMsgReq(wechatMsgReq)
	} else {
		chatMsgReq = createChatMsgReqWithHistory(wechatMsgReq, chatMsgReq)
	}

	// 调用接口
	chatMsgResp, err := postToWenXin(chatMsgReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error" + err.Error(),
		})
		return
	}

	// 保存查询的结果
	chatMsgReq.Messages = append(chatMsgReq.Messages, Msg{
		Role:    "assistant",
		Content: chatMsgResp.Result,
		UserID:  wechatMsgReq.FromUserName,
	})
	err = setMsgHistory(wechatMsgReq.FromUserName, chatMsgReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error" + err.Error(),
		})
		return
	}

	// 这里测试 回复消息
	replyMsg := WeChatReplyMsgRes{
		ToUserName:   wechatMsgReq.FromUserName,
		FromUserName: wechatMsgReq.ToUserName,
		CreateTime:   wechatMsgReq.CreateTime,
		MsgType:      "text",
		Content:      chatMsgResp.Result,
	}
	// 返回响应
	c.JSON(http.StatusOK, replyMsg)
}
