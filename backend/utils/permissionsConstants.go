package utils

const (
	admin_prefix = "admin"

	admin_ban_prefix    = admin_prefix + ".ban"
	PERMISSION_BAN_VIEW = admin_ban_prefix + ".view"
	PERMISSION_BAN_EDIT = admin_ban_prefix + ".edit"

	admin_report_prefix    = admin_prefix + ".report"
	PERMISSION_REPORT_VIEW = admin_report_prefix + ".view"
	PERMISSION_REPORT_EDIT = admin_report_prefix + ".edit"

	admin_texture_prefix      = admin_prefix + ".texture"
	PERMISSION_TEXTURE_VIEW   = admin_texture_prefix + ".view"
	PERMISSION_TEXTURE_EDIT   = admin_texture_prefix + ".edit"
	PERMISSION_TEXTURE_UPLOAD = admin_texture_prefix + ".upload"

	admin_map_prefix      = admin_prefix + ".map"
	PERMISSION_MAP_VIEW   = admin_map_prefix + ".view"
	PERMISSION_MAP_EDIT   = admin_map_prefix + ".edit"
	PERMISSION_MAP_UPLOAD = admin_map_prefix + ".upload"
	PERMISSION_MAP_EDITOR = admin_map_prefix + ".editor"

	admin_user_prefix            = admin_prefix + ".user"
	PERMISSION_USER_VIEW         = admin_map_prefix + ".view"
	PERMISSION_USER_EDIT         = admin_map_prefix + ".edit"
	PERMISSION_USER_INVITE       = admin_map_prefix + ".invite"
	PERMISSION_USER_IMPOSTERNATE = admin_map_prefix + ".imposternate"
)
