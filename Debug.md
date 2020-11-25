Debug 流程：
1. 首先 MySQL 库执行 data/db
2. 根据 db 生成 tables 和 generators 并移动到 models2 ：`go install github.com/GoAdminGroup/go-admin/adm && adm generate`
3. 先在 git.jd.com/jdolap/adminserver 开发好 Controller 
4. push 到某个分支
5. goadmin/go.mod 添加分支依赖
6. 如修改添加账号功能
    1. 新开发一个文件：plugins/admin/controller/account.go
    2. plugins/admin/controller/new.go 修改 NewForm 直接调用自己的 NewAccountForm
    3. 启动，使用测试集群测试 
7. 前端页面是按照数据库结构生成好的，不需要改，该项目仅支持简单功能测试
8. 打包到开发环境，`make serve` 启动
9. 日志：adminserver/LOG 


#### 附
##### 初始化
goadmin 初始化过程：
- 初始化 engine：调用 Register注册默认 adapter，默认adapter包括：{db，gin.Context，gin.Engine}
- 初始化 engine.config：调用InitDB 获取数据库连接池（ InitDB的实现在modules/db/mysql）
    - 初始化eng.Services（map{driver:dbList{name, conn}}）和 eng.Adapter
				     以及 eng.DefaultConnection()
- AddGenerators 新建eng.PluginList对象pluginList：调用 plugins/admin/NewAdmin 新建 plugins/admin.Admin 对象，添加自定义GeneratorList 到 admin.tablelist
- Use(r)：
	1. 添加eng.Services[“token_csrf_helper”]：调用 modules/auth.InitCSRFTokenSrv 查询表goadmin_session 获取 csrf_tokens 
	2. 添加 site Config：调用 plugins/admin/(models.SiteModel).AllToMap 从表 goadmin_site 中查询 key/value 注入 Config 中，Config 添加到 eng.Services
	3. 添加 eng.Services[“ui”] ：UI组件NavJumpButtons
	4. plugins：调用 plugins/admin.InitPlugin 
		1. append 系统表到 admin.tableList
		2. new Guardian 到 admin.guardian：维护 admin.services, admin.conn, admin.tableList, admin.UI
		3. 初始化路由和控制器：plugins/admin.initRouter，维护在 ctx
			
			- 所有路由分组=config.Prefix()
			- route.Middlewares：admin.globalErrorHandler 
			- 添加登录 config.GetLoginUrl()： admin.handler.ShowLogin
			- 调用 adapter-> gin.AddHandler 转换 gin.Context -> admin.Context

开启 config api ：{"open_admin_api": true}

##### 请求路由都在这里
Goadmin 请求过程：debug 查看 handlers 列表

plugins/admin/router.go
```
// add delete modify query
	authPrefixRoute.GET("/info/:__prefix/detail", admin.handler.ShowDetail).Name("detail")
	authPrefixRoute.GET("/info/:__prefix/edit", admin.guardian.ShowForm, admin.handler.ShowForm).Name("show_edit")
	authPrefixRoute.GET("/info/:__prefix/new", admin.guardian.ShowNewForm, admin.handler.ShowNewForm).Name("show_new")
	authPrefixRoute.POST("/edit/:__prefix", admin.guardian.EditForm, admin.handler.EditForm).Name("edit")
	authPrefixRoute.POST("/new/:__prefix", admin.guardian.NewForm, admin.handler.NewForm).Name("new")
	authPrefixRoute.POST("/delete/:__prefix", admin.guardian.Delete, admin.handler.Delete).Name("delete")
	authPrefixRoute.POST("/export/:__prefix", admin.guardian.Export, admin.handler.Export).Name("export")
	authPrefixRoute.GET("/info/:__prefix", admin.handler.ShowInfo).Name("info")

	authPrefixRoute.POST("/update/:__prefix", admin.guardian.Update, admin.handler.Update).Name("update")
```

##### list 请求路径
查：http://localhost:9033/admin/api/list/business	
 
 prefix=business

调用链路：

调用 ShowInfo
- ->调用 plugins/admin/controller/common.go 
	h.table()-> h.generators[prefix](ctx) -> models2.GetBusinessTable() -> 
	
- ->调用 plugins/admin/(*controller.Handler).showTableData 
- ->调用 plugins/admin/modules/table/default.go 的 (*table.DefaultTable).GetData 方法	
	
// 获取 Table data 的地方
```
if tb.getDataFun != nil {
	data, size = tb.getDataFun(params)  // 可以自定义方法
} else if tb.sourceURL != "" {
	data, size = tb.getDataFromURL(params)
} else if tb.Info.GetDataFn != nil {
	data, size = tb.Info.GetDataFn(params)
} else if params.IsAll() {
	return tb.getAllDataFromDatabase(params)
} else {
	return tb.getDataFromDatabase(params)  //  默认走这里
}
```

最简单的：在 default.go 直接修改 GetDataFun()

或者：自己实现 table.go 的 Table  接口，参考 default.go 的实现

params := parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
	panel.GetInfo().GetSort())

过滤条件：params.Fields



