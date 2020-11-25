package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBusinessTable(ctx *context.Context) table.Table {

	business := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := business.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Contacts", "contacts", db.Varchar)

	info.SetTable("business").SetTitle("Business").SetDescription("Business")

	formList := business.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Contacts", "contacts", db.Varchar, form.Text)

	formList.SetTable("business").SetTitle("Business").SetDescription("Business")

	return business
}


