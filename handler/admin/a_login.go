package admin

import (
	"github.com/ops-cn/admin/app/ginplus"
	"github.com/ops-cn/common/config"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := c.Request.Context()
	item, err := a.LoginBll.GetCaptcha(ctx, config.C.Captcha.Length)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}
