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
	ToUserName   string  `json:"ToUserName" form:"ToUserName" url:"ToUserName"`
	FromUserName string  `json:"FromUserName" form:"FromUserName" url:"FromUserName"`
	CreateTime   int64   `json:"CreateTime" form:"CreateTime" url:"CreateTime"`
	MsgType      MsgType `json:"MsgType" form:"MsgType" url:"MsgType"`
	Content      string  `json:"Content" form:"Content" url:"Content"`
	MsgId        int64   `json:"MsgId" form:"MsgId" url:"MsgId"`
	PicUrl       string  `json:"PicUrl" form:"PicUrl" url:"PicUrl"`
	MediaId      string  `json:"MediaId" form:"MediaId" url:"MediaId"`
	AppId        string  `json:"AppId" form:"AppId" url:"AppId"`
	PagePath     string  `json:"PagePath" form:"PagePath" url:"PagePath"`
	ThumbUrl     string  `json:"ThumbUrl" form:"ThumbUrl" url:"ThumbUrl"`
	ThumbMediaId string  `json:"ThumbMediaId" form:"ThumbMediaId" url:"ThumbMediaId"`
}

const (
	GrantType       = "client_credential"
	Text            = MsgType("text")
	Image           = MsgType("image")
	Link            = MsgType("link")
	MiniProgramPage = MsgType("miniprogrampage")
)

type MsgType string
type TextMsg struct {
	ToUser  string  `json:"touser" form:"touser" url:"touser"`
	MsgType MsgType `json:"msgtype" form:"msgtype" url:"msgtype"`
	Text    struct {
		Content string `json:"content" form:"content" url:"content"`
	} `json:"text" form:"text" url:"text"`
}

type SendResult struct {
	ErrCode int    `json:"errcode" form:"errcode" url:"errcode"`
	ErrMsg  string `json:"errmsg" form:"errmsg" url:"errmsg"`
}

type ImageMsg struct {
	ToUser  string  `json:"touser" form:"touser" url:"touser"`
	MsgType MsgType `json:"msgtype" form:"msgtype" url:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id" form:"media_id" url:"media_id"`
	} `json:"image" form:"image" url:"image"`
}

type LinkMsg struct {
	ToUser  string
	MsgType MsgType
	Link    struct {
		Title       string `json:"title" form:"title" url:"title"`
		Description string `json:"description" form:"description" url:"description"`
		URL         string `json:"url" form:"url" url:"url"`
		ThumbURL    string `json:"thumb_url" form:"thumb_url" url:"thumb_url"`
	} `json:"link" form:"link" url:"link"`
}

type MiniProgramPageMsg struct {
	ToUser          string  `json:"touser" form:"touser" url:"touser"`
	MsgType         MsgType `json:"msgtype" form:"msgtype" url:"msgtype"`
	MiniProgramPage struct {
		Title       string `json:"title" form:"title" url:"title"`
		PagePath    string `json:"pagepath" form:"pagepath" url:"pagepath"`
		TumbMediaId string `json:"thumb_media_id" form:"thumb_media_id" url:"thumb_media_id"`
	} `json:"miniprogrampage" form:"miniprogrampage" url:"miniprogrampage"`
}

type TransferMsg struct {
	ToUserName   string `json:"ToUserName" form:"ToUserName" url:"ToUserName"`
	FromUserName string `json:"FromUserName" form:"FromUserName" url:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" form:"CreateTime" url:"CreateTime"`
	MsgType      string `json:"MsgType" form:"MsgType" url:"MsgType"`
}
