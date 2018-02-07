package post

import (
	"fmt"
	"github.com/missdeer/kelly/routers/base"
)

type ForwarderRouter struct {
	base.BaseRouter
}

func (this *ForwarderRouter) TaobaoItem() {
	id := this.GetString(":id")
	url := fmt.Sprintf("http://item.taobao.com/item.htm?id=%s", id)
	this.Redirect(url, 302)
}
