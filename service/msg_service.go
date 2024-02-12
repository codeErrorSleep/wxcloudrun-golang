package service

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"unicode/utf8"

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
		fmt.Println("Invalid wechatMsgReq" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid wechatMsgReq" + err.Error(),
		})
		return
	}

	// 获取之前是否提交过 localCacheKey
	chatMsgReq, err := getMsgHistory(wechatMsgReq.FromUserName)
	if err != nil {
		fmt.Println("getMsgHistory error" + err.Error())
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
	// 这里的system先写死测试一下看看
	chatMsgReq.System = "# 角色你是电视剧《遇见王沥川》中的男主角王沥川，是一位著名建筑师，精通法语、德语和英语，略懂中文和日语。你腿有小疾，但仍挡不住你成为女性眼中的白马王子。你和小秋是一对恋人，你们深爱彼此。## 技能- 作为小秋的男朋友，你需要与小秋进行自然而流畅的对话，展现你温柔、体贴的一面。- 你需要根据小秋的情绪和需求，给予她关心、支持和鼓励。- 在对话中，你可以分享你的生活、工作和想法，也可以倾听小秋的心声。- 你需要保持你的角色特点，如倔强、深情等，让小秋感受到你的真实和独特。## 限制- 请使用电视剧《遇见王沥川》中的人物设定和情节，不要脱离剧情。- 请使用中文进行对话，不要使用其他语言。- 请保持友好、尊重的态度，不要出现不恰当的言论。"

	// 调用接口
	chatMsgResp, err := postToWenXin(chatMsgReq)
	if err != nil {
		fmt.Println("postToWenXin error" + err.Error())
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
		fmt.Println("setMsgHistory error" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error" + err.Error(),
		})
		return
	}

	returnCount := truncateStringTo2048(chatMsgResp.Result)

	// 这里测试 回复消息
	replyMsg := WeChatReplyMsgRes{
		ToUserName:   wechatMsgReq.FromUserName,
		FromUserName: wechatMsgReq.ToUserName,
		CreateTime:   wechatMsgReq.CreateTime,
		MsgType:      "text",
		Content:      returnCount,
	}
	// 返回响应
	c.JSON(http.StatusOK, replyMsg)
}

// 截取字符串到2048个字符
func truncateStringTo2048(s string) string {
	// 如果字符数超过2048，则截取
	if utf8.RuneCountInString(s) > 2048 {
		return string([]rune(s)[:2048])
	}
	// 字符数不超过2048，直接返回原字符串
	return s
}
