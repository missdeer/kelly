package post

import (
	"fmt"
	"strings"

	"github.com/missdeer/kelly/routers/base"
)

type SearchRouter struct {
	base.BaseRouter
}

func (this *SearchRouter) Get() {
	q := strings.TrimSpace(this.GetString("q"))
	fmt.Println("search ", q)
	this.Redirect(fmt.Sprintf("https://s.yii.li/search?newwindow=1&safe=strict&q=%s+site%3Ayii.li&oq=%s+site%3Ayii.li", q, q), 302)
}
