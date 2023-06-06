package user

import "github.com/17660472762/miniprogramtest/pkg/codec"

type Service struct {
	wechatconfig codec.WechatConfig
}

func NewService(appid, secret string) Service {
	return Service{
		wechatconfig: codec.WechatConfig{
			GrantType: codec.GrantType,
			AppID:     appid,
			Secret:    secret,
		},
	}
}
