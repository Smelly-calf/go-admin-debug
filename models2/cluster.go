package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetClusterTable(ctx *context.Context) table.Table {

	cluster := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := cluster.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Alias", "alias", db.Varchar)
	info.AddField("Admin_password", "admin_password", db.Varchar)
	info.AddField("Http_addr", "http_addr", db.Varchar)
	info.AddField("Zk_addrs", "zk_addrs", db.Varchar)

	info.SetTable("cluster").SetTitle("Cluster").SetDescription("Cluster")

	formList := cluster.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Alias", "alias", db.Varchar, form.Text)
	formList.AddField("Admin_password", "admin_password", db.Varchar, form.Text)
	formList.AddField("Http_addr", "http_addr", db.Varchar, form.Text)
	formList.AddField("Zk_addrs", "zk_addrs", db.Varchar, form.Text)

	formList.SetTable("cluster").SetTitle("Cluster").SetDescription("Cluster")

	return cluster
}
