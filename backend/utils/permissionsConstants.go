package utils

const (
	adminPrefix = "admin"

	adminAnnouncementPrefix    = adminPrefix + ".announcement"
	PermissionAnnouncementView = adminAnnouncementPrefix + ".view"
	PermissionAnnouncementEdit = adminAnnouncementPrefix + ".edit"

	adminBanPrefix      = adminPrefix + ".ban"
	PermissionBanView   = adminBanPrefix + ".view"
	PermissionBanCreate = adminBanPrefix + ".create"
	PermissionBanDelete = adminBanPrefix + ".delete"

	adminReportPrefix    = adminPrefix + ".report"
	PermissionReportView = adminReportPrefix + ".view"
	PermissionReportEdit = adminReportPrefix + ".edit"

	adminTexturePrefix      = adminPrefix + ".texture"
	PermissionTextureView   = adminTexturePrefix + ".view"
	PermissionTextureEdit   = adminTexturePrefix + ".edit"
	PermissionTextureUpload = adminTexturePrefix + ".upload"

	adminMapPrefix      = adminPrefix + ".map"
	PermissionMapView   = adminMapPrefix + ".view"
	PermissionMapEdit   = adminMapPrefix + ".edit"
	PermissionMapUpload = adminMapPrefix + ".upload"
	PermissionMapEditor = adminMapPrefix + ".editor"

	adminUserPrefix            = adminPrefix + ".user"
	PermissionUserView         = adminUserPrefix + ".view"
	PermissionUserEdit         = adminUserPrefix + ".edit"
	PermissionUserInvite       = adminUserPrefix + ".invite"
	PermissionUserImposternate = adminUserPrefix + ".imposternate"

	adminGroupPrefix    = adminPrefix + ".group"
	PermissionGroupEdit = adminGroupPrefix + ".edit"
)
