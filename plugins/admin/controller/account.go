package controller

import (
	"fmt"
	"git.jd.com/jdolap/adminserver/pkg/cluster"
	"git.jd.com/jdolap/adminserver/pkg/controller"
	"github.com/GoAdminGroup/go-admin/context"
	models "github.com/GoAdminGroup/go-admin/models2"
	"github.com/GoAdminGroup/go-admin/modules/language"

	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"net/http"
)

func (h *Handler) NewAccountForm(ctx *context.Context) {

	fmt.Println("=============== NewAccountForm")

	param := guard.GetNewFormParam(ctx)

	// process uploading files, only support local storage
	if len(param.MultiForm.File) > 0 {
		err := file.GetFileEngine(h.config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			logger.Error("get file engine error: ", err)
			if ctx.WantJSON() {
				response.Error(ctx, err.Error())
			} else {
				h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
			}
			return
		}
	}

	dataList := param.Value()
	// validate business and cluster
	accountModel := models.FormValueToAccountModel(dataList)
	err := controller.DefaultAccountCtrl.ValidateAccount(&accountModel)
	if err != nil {
		logger.Error("validate cluster and business error: ", err)
		if ctx.WantJSON() {
			response.Error(ctx, err.Error(), map[string]interface{}{
				"token": h.authSrv().AddToken(),
			})
		} else {
			h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
		}
	}

	err = param.Panel.InsertData(dataList)
	if err != nil {
		logger.Error("insert data error: ", err)
		if ctx.WantJSON() {
			response.Error(ctx, err.Error(), map[string]interface{}{
				"token": h.authSrv().AddToken(),
			})
		} else {
			h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
		}
		return
	}

	// account to ck
	fmt.Println("=============== create account to ck")
	fmt.Printf("business:%s, account: %s, cluster info: %+v\n", accountModel.Business.Name, accountModel.Account, *accountModel.Cluster)
	err = cluster.CreateCKAccount(&accountModel)
	fmt.Println("=============== err:", err)
	if err != nil {
		logger.Error("create ck account error: ", err)
		if ctx.WantJSON() {
			response.Error(ctx, err.Error(), map[string]interface{}{
				"token": h.authSrv().AddToken(),
			})
		} else {
			h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
		}
		return
	}

	f := param.Panel.GetActualNewForm()

	if f.Responder != nil {
		f.Responder(ctx)
		return
	}

	if ctx.WantJSON() && !param.IsIframe {
		response.OkWithData(ctx, map[string]interface{}{
			"url":   param.PreviousPath,
			"token": h.authSrv().AddToken(),
		})
		return
	}

	if !param.FromList {

		if isNewUrl(param.PreviousPath, param.Prefix) {
			h.showNewForm(ctx, param.Alert, param.Prefix, param.Param.GetRouteParamStr(), true)
			return
		}

		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	if param.IsIframe {
		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>
		swal('%s', '', 'success');
		setTimeout(function(){
			$("#%s", window.parent.document).hide();
			$('.modal-backdrop.fade.in', window.parent.document).hide();
		}, 1000)
</script>`, language.Get("success"), param.IframeID))
		return
	}

	buf := h.showTable(ctx, param.Prefix, param.Param, nil)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, h.routePathWithPrefix("info", param.Prefix)+param.Param.GetRouteParamStr())
}
