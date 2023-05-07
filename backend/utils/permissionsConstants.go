package utils

const (
	admin_prefix = "admin"

	admin_announcement_prefix    = admin_prefix + ".announcement"
	PERMISSION_ANNOUNCEMENT_VIEW = admin_announcement_prefix + ".view"
	PERMISSION_ANNOUNCEMENT_EDIT = admin_announcement_prefix + ".edit"

	admin_ban_prefix      = admin_prefix + ".ban"
	PERMISSION_BAN_VIEW   = admin_ban_prefix + ".view"
	PERMISSION_BAN_CREATE = admin_ban_prefix + ".create"
	PERMISSION_BAN_DELETE = admin_ban_prefix + ".delete"

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
	PERMISSION_USER_VIEW         = admin_user_prefix + ".view"
	PERMISSION_USER_EDIT         = admin_user_prefix + ".edit"
	PERMISSION_USER_INVITE       = admin_user_prefix + ".invite"
	PERMISSION_USER_IMPOSTERNATE = admin_user_prefix + ".imposternate"

	admin_group_prefix    = admin_prefix + ".group"
	PERMISSION_GROUP_EDIT = admin_group_prefix + ".edit"
)
