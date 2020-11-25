package models

import (
	"git.jd.com/jdolap/adminserver/pkg/models"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"strconv"
)

func FormValueToAccountModel(values form2.Values) models.AccountModel {
	tmpClusterID := values.Get("cluster_id")
	clusterID, _ := strconv.ParseInt(tmpClusterID, 10, 64)

	tmpBusinessID := values.Get("business_id")
	businessID, _ := strconv.ParseInt(tmpBusinessID, 10, 64)

	tmpAccountType := values.Get("account_type")
	accounType, _ := strconv.Atoi(tmpAccountType)

	return models.AccountModel{
		ClusterID:   clusterID,
		BusinessID:  businessID,
		Account:     values.Get("account"),
		Password:    values.Get("password"),
		Databases:   values.Get("databases"),
		AccountType: accounType,
	}
}

func GetAccountTable(ctx *context.Context) table.Table {

	account := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	//info := account.GetInfo().HideFilterArea()
	info := account.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Cluster_id", "cluster_id", db.Int).
		FieldFilterable()
	info.AddField("Business_id", "business_id", db.Int).
		FieldFilterable()
	info.AddField("Account", "account", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Account_type", "account_type", db.Int).FieldXssFilter()
	info.AddField("Databases", "databases", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("account").SetTitle("Account").SetDescription("Account")

	formList := account.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Cluster_id", "cluster_id", db.Int, form.Number)
	formList.AddField("Business_id", "business_id", db.Int, form.Number)
	formList.AddField("Account", "account", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Account_type", "account_type", db.Int, form.Number)
	formList.AddField("Databases", "databases", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("account").SetTitle("Account").SetDescription("Account")

	return account
}
