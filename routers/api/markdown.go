package api

import (
	"github.com/missdeer/kelly/modules/utils"
)

func (this *ApiRouter) Markdown() {
	if this.CheckActiveRedirect() {
		return
	}

	if this.IsAjax() {
		result := map[string]interface{}{
			"success": false,
		}
		action := this.GetString("action")
		switch action {
		case "preview":
			content := this.GetString("content")
			result["preview"] = utils.Render(content)
			result["success"] = true
		}
		this.Data["json"] = result
		this.ServeJSON()
	}
}
