package codec

type WechatConfig struct {
	GrantType string `json:"grant_type" form:"grant_type" url:"grant_type"`
	AppID     string `json:"appid" form:"appid" url:"appid"`
	Secret    string `json:"secret" form:"secret" url:"secret"`
}

type AccessToken struct {
	AccessToken string `json:"access_token" form:"access_token" url:"access_token"`
	ExpiresIn   int64  `json:"expires_in" form:"expires_in" url:"expires_in"`
}

type ReceiveTextMsg struct {
	ToUserName   string `json:"ToUserName" form:"ToUserName" url:"ToUserName"`
	FromUserName string `json:"FromUserName" form:"FromUserName" url:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" form:"CreateTime" url:"CreateTime"`
	MsgType      string `json:"MsgType" form:"MsgType" url:"MsgType"`
	Content      string `json:"Content" form:"Content" url:"Content"`
	MsgId        int64  `json:"MsgId" form:"MsgId" url:"MsgId"`
}
