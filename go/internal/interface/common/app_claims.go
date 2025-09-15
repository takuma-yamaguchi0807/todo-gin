package common

import (
	"github.com/gin-gonic/gin"
	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
)

// ClaimsFromContext はミドルウェアで格納されたドメインの Claims を取り出します。
func ClaimsFromContext(c *gin.Context) (domainauth.Claims, bool) {
    v, ok := c.Get("claims")
    if !ok { return domainauth.Claims{}, false }
    claims, ok := v.(domainauth.Claims)
    if !ok { return domainauth.Claims{}, false }
    return claims, true
}


