package controllers

import (
	gvalid "ThingsPanel-Go/initialize/validate"
	"ThingsPanel-Go/models"
	"ThingsPanel-Go/services"
	"ThingsPanel-Go/utils"
	valid "ThingsPanel-Go/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type HdlTasteController struct {
	beego.Controller
}

// 列表
func (c *HdlTasteController) List() {
	reqData := valid.HdlTastePaginationValidate{}
	if err := valid.ParseAndValidate(&c.Ctx.Input.RequestBody, &reqData); err != nil {
		utils.SuccessWithMessage(1000, err.Error(), (*context2.Context)(c.Ctx))
		return
	}
	var HdlTasteService services.HdlTasteService
	d, t, err := HdlTasteService.GetHdlTasteList(reqData)
	if err != nil {
		utils.SuccessWithMessage(1000, err.Error(), (*context2.Context)(c.Ctx))
		return
	}
	dd := valid.RspHdlTastePaginationValidate{
		CurrentPage: reqData.CurrentPage,
		Data:        d,
		Total:       t,
		PerPage:     reqData.PerPage,
	}
	utils.SuccessWithDetailed(200, "success", dd, map[string]string{}, (*context2.Context)(c.Ctx))
}

// 编辑
func (c *HdlTasteController) Edit() {
	reqData := valid.EditHdlTasteValidate{}
	if err := valid.ParseAndValidate(&c.Ctx.Input.RequestBody, &reqData); err != nil {
		utils.SuccessWithMessage(1000, err.Error(), (*context2.Context)(c.Ctx))
		return
	}
	var HdlTasteService services.HdlTasteService
	err := HdlTasteService.EditHdlTaste(reqData)
	if err == nil {
		d := HdlTasteService.GetHdlTasteDetail(reqData.Id)
		utils.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(c.Ctx))
	} else {
		utils.SuccessWithMessage(400, err.Error(), (*context2.Context)(c.Ctx))
	}
}

// 新增
func (c *HdlTasteController) Add() {
	reqData := valid.AddHdlTasteValidate{}
	if err := valid.ParseAndValidate(&c.Ctx.Input.RequestBody, &reqData); err != nil {
		utils.SuccessWithMessage(1000, err.Error(), (*context2.Context)(c.Ctx))
		return
	}
	var HdlTasteService services.HdlTasteService
	d, rsp_err := HdlTasteService.AddHdlTaste(reqData)
	if rsp_err == nil {
		utils.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(c.Ctx))
	} else {
		utils.SuccessWithMessage(400, rsp_err.Error(), (*context2.Context)(c.Ctx))
	}
}

// 删除
func (HdlTasteController *HdlTasteController) Delete() {
	HdlTasteIdValidate := valid.HdlTasteIdValidate{}
	err := json.Unmarshal(HdlTasteController.Ctx.Input.RequestBody, &HdlTasteIdValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(HdlTasteIdValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(HdlTasteIdValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			utils.SuccessWithMessage(1000, message, (*context2.Context)(HdlTasteController.Ctx))
			break
		}
		return
	}
	if HdlTasteIdValidate.Id == "" {
		utils.SuccessWithMessage(1000, "id不能为空", (*context2.Context)(HdlTasteController.Ctx))
	}
	var HdlTasteService services.HdlTasteService
	HdlTaste := models.HdlTaste{
		Id: HdlTasteIdValidate.Id,
	}
	req_err := HdlTasteService.DeleteHdlTaste(HdlTaste)
	if req_err == nil {
		utils.SuccessWithMessage(200, "success", (*context2.Context)(HdlTasteController.Ctx))
	} else {
		utils.SuccessWithMessage(400, "删除失败", (*context2.Context)(HdlTasteController.Ctx))
	}
}
