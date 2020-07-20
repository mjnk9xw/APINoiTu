package controllers

import (
	"APINoiTu/models"
	"encoding/json"
)

// @Title API Check
// @Description API check
// @Param	body		body 	models.BodyCheck	true	"body param"
// @Success 200 {object} models.BoolResponse
// @router /check [post]
func (o *Controller) Check() {

	var data *models.BodyCheck
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &data)
	if err != nil {
		o.Data["json"] = models.BoolResponse{
			Data:      false,
			ErrorMess: "Yêu cầu gửi lên sai cấu trúc",
		}
		o.ServeJSON()
		return
	}

	modelsIf := models.GetModels()
	r, e := modelsIf.CheckTu(data.PreSeq, data.Seq)
	o.Data["json"] = models.BoolResponse{
		Data:      r,
		ErrorMess: e.Error(),
	}
	o.ServeJSON()
	return

}
