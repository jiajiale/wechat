package context

import (
	"github.com/jiajiale/wechat/credential"
	"github.com/jiajiale/wechat/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
