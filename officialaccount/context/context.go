package context

import (
	"github.com/jiajiale/wechat/credential"
	"github.com/jiajiale/wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
