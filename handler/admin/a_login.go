package admin

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/ops-cn/admin/app/ginplus"
	"github.com/ops-cn/common/captcha"
	"github.com/ops-cn/common/config"
	proto "github.com/ops-cn/proto/admin/login"
	"github.com/ops-cn/proto/unified"
	"net/http"
)

func GetCaptcha(w http.ResponseWriter, r *http.Request) unified.Response {

	ctx := r.Context()
	captchaID := captcha.NewLen(config.C.Captcha.Length)

	item := &proto.LoginCaptcha{
		CaptchaID: captchaID,
	}

	any, err := ptypes.MarshalAny(item)
	return unified.Response{
		TraceId: "",
		Status:  0,
		Desc:    "",
		Items:   any,
	}

	if err != nil {
		ginplus.ResError(ctx, err)
		return
	}
	ginplus.ResSuccess(ctx, item)
}
