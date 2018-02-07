package admin

type AdminDashboardRouter struct {
	BaseAdminRouter
}

func (this *AdminDashboardRouter) Get() {
	this.Data["consoleAdmin"] = true
	this.TplName = "admin/dashboard.html"
}
