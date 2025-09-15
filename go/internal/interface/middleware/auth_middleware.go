package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

// AuthRequired は Authorization: Bearer <token> を検証し、正しければ Claims をコンテキストに格納します。
// 何をするか: HTTPヘッダからトークン抽出 → TokenService.Verify 実行 → 失敗時は 401 を返却
// なぜここか: HTTP固有の責務（ヘッダ処理・401返却）は interface 層に閉じるため
func AuthRequired(verifyService domainauth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authz := c.GetHeader("Authorization")
		if !strings.HasPrefix(authz, "Bearer ") {
			// Authorization ヘッダ不備はアプリ標準のエラーフォーマットで返却
			status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing or invalid bearer token"))
			c.AbortWithStatusJSON(status, payload)
			return
		}
		token := strings.TrimPrefix(authz, "Bearer ")
		claims, err := verifyService.Verify(token)
		if err != nil {
			// トークン検証失敗も標準フォーマットで返却（詳細はメッセージに載せない）
			status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "invalid token"))
			c.AbortWithStatusJSON(status, payload)
			return
		}
		// 後続ハンドラで使えるようにクレームを格納
		c.Set("claims", claims)
		c.Next()
	}
}


