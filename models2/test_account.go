package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTestAccountTable(ctx *context.Context) table.Table {

	testAccount := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testAccount.GetInfo()

	info.AddField("Id", "id", db.Int)
	info.AddField("Cluster", "cluster", db.Varchar)
	info.AddField("Business", "business", db.Varchar)
	info.AddField("User_name", "user_name", db.Varchar)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Database", "database", db.Varchar)
	info.AddField("Contacts", "contacts", db.Varchar)

	info.SetTable("test_account").SetTitle("TestAccount").SetDescription("TestAccount")

	formList := testAccount.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Cluster", "cluster", db.Varchar, form.Text)
	formList.AddField("Business", "business", db.Varchar, form.Text)
	formList.AddField("User_name", "user_name", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Database", "database", db.Varchar, form.Text)
	formList.AddField("Contacts", "contacts", db.Varchar, form.Text)

	formList.SetTable("test_account").SetTitle("TestAccount").SetDescription("TestAccount")

	return testAccount
}
