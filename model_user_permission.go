/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// UserPermission ユーザー権限
type UserPermission string

// List of UserPermission
const (
	USERPERMISSION_GetWebhook                UserPermission = "get_webhook"
	USERPERMISSION_CreateWebhook             UserPermission = "create_webhook"
	USERPERMISSION_EditWebhook               UserPermission = "edit_webhook"
	USERPERMISSION_DeleteWebhook             UserPermission = "delete_webhook"
	USERPERMISSION_AccessOthersWebhook       UserPermission = "access_others_webhook"
	USERPERMISSION_GetBot                    UserPermission = "get_bot"
	USERPERMISSION_CreateBot                 UserPermission = "create_bot"
	USERPERMISSION_EditBot                   UserPermission = "edit_bot"
	USERPERMISSION_DeleteBot                 UserPermission = "delete_bot"
	USERPERMISSION_AccessOthersBot           UserPermission = "access_others_bot"
	USERPERMISSION_BotActionJoinChannel      UserPermission = "bot_action_join_channel"
	USERPERMISSION_BotActionLeaveChannel     UserPermission = "bot_action_leave_channel"
	USERPERMISSION_CreateChannel             UserPermission = "create_channel"
	USERPERMISSION_GetChannel                UserPermission = "get_channel"
	USERPERMISSION_EditChannel               UserPermission = "edit_channel"
	USERPERMISSION_DeleteChannel             UserPermission = "delete_channel"
	USERPERMISSION_ChangeParentChannel       UserPermission = "change_parent_channel"
	USERPERMISSION_EditChannelTopic          UserPermission = "edit_channel_topic"
	USERPERMISSION_GetChannelStar            UserPermission = "get_channel_star"
	USERPERMISSION_EditChannelStar           UserPermission = "edit_channel_star"
	USERPERMISSION_GetMyTokens               UserPermission = "get_my_tokens"
	USERPERMISSION_RevokeMyToken             UserPermission = "revoke_my_token"
	USERPERMISSION_GetClients                UserPermission = "get_clients"
	USERPERMISSION_CreateClient              UserPermission = "create_client"
	USERPERMISSION_EditMyClient              UserPermission = "edit_my_client"
	USERPERMISSION_DeleteMyClient            UserPermission = "delete_my_client"
	USERPERMISSION_ManageOthersClient        UserPermission = "manage_others_client"
	USERPERMISSION_UploadFile                UserPermission = "upload_file"
	USERPERMISSION_DownloadFile              UserPermission = "download_file"
	USERPERMISSION_DeleteFile                UserPermission = "delete_file"
	USERPERMISSION_GetMessage                UserPermission = "get_message"
	USERPERMISSION_PostMessage               UserPermission = "post_message"
	USERPERMISSION_EditMessage               UserPermission = "edit_message"
	USERPERMISSION_DeleteMessage             UserPermission = "delete_message"
	USERPERMISSION_ReportMessage             UserPermission = "report_message"
	USERPERMISSION_GetMessageReports         UserPermission = "get_message_reports"
	USERPERMISSION_CreateMessagePin          UserPermission = "create_message_pin"
	USERPERMISSION_DeleteMessagePin          UserPermission = "delete_message_pin"
	USERPERMISSION_GetChannelSubscription    UserPermission = "get_channel_subscription"
	USERPERMISSION_EditChannelSubscription   UserPermission = "edit_channel_subscription"
	USERPERMISSION_ConnectNotificationStream UserPermission = "connect_notification_stream"
	USERPERMISSION_RegisterFCMDevice         UserPermission = "register_fcm_device"
	USERPERMISSION_GetStamp                  UserPermission = "get_stamp"
	USERPERMISSION_CreateStamp               UserPermission = "create_stamp"
	USERPERMISSION_EditStamp                 UserPermission = "edit_stamp"
	USERPERMISSION_EditStampCreatedByOthers  UserPermission = "edit_stamp_created_by_others"
	USERPERMISSION_DeleteStamp               UserPermission = "delete_stamp"
	USERPERMISSION_AddMessageStamp           UserPermission = "add_message_stamp"
	USERPERMISSION_RemoveMessageStamp        UserPermission = "remove_message_stamp"
	USERPERMISSION_GetMyStampHistory         UserPermission = "get_my_stamp_history"
	USERPERMISSION_GetStampPalette           UserPermission = "get_stamp_palette"
	USERPERMISSION_CreateStampPalette        UserPermission = "create_stamp_palette"
	USERPERMISSION_EditStampPalette          UserPermission = "edit_stamp_palette"
	USERPERMISSION_DeleteStampPalette        UserPermission = "delete_stamp_palette"
	USERPERMISSION_GetUser                   UserPermission = "get_user"
	USERPERMISSION_RegisterUser              UserPermission = "register_user"
	USERPERMISSION_GetMe                     UserPermission = "get_me"
	USERPERMISSION_EditMe                    UserPermission = "edit_me"
	USERPERMISSION_ChangeMyIcon              UserPermission = "change_my_icon"
	USERPERMISSION_ChangeMyPassword          UserPermission = "change_my_password"
	USERPERMISSION_EditOtherUsers            UserPermission = "edit_other_users"
	USERPERMISSION_GetUserQRCode             UserPermission = "get_user_qr_code"
	USERPERMISSION_GetUserTag                UserPermission = "get_user_tag"
	USERPERMISSION_EditUserTag               UserPermission = "edit_user_tag"
	USERPERMISSION_GetUserGroup              UserPermission = "get_user_group"
	USERPERMISSION_CreateUserGroup           UserPermission = "create_user_group"
	USERPERMISSION_CreateSpecialUserGroup    UserPermission = "create_special_user_group"
	USERPERMISSION_EditUserGroup             UserPermission = "edit_user_group"
	USERPERMISSION_DeleteUserGroup           UserPermission = "delete_user_group"
	USERPERMISSION_AllUserGroupsAdmin        UserPermission = "edit_others_user_group"
	USERPERMISSION_WebRTC                    UserPermission = "web_rtc"
	USERPERMISSION_GetMySessions             UserPermission = "get_my_sessions"
	USERPERMISSION_DeleteMySessions          UserPermission = "delete_my_sessions"
	USERPERMISSION_GetMyExternalAccount      UserPermission = "get_my_external_account"
	USERPERMISSION_EditMyExternalAccount     UserPermission = "edit_my_external_account"
	USERPERMISSION_GetUnread                 UserPermission = "get_unread"
	USERPERMISSION_DeleteUnread              UserPermission = "delete_unread"
	USERPERMISSION_GetClipFolder             UserPermission = "get_clip_folder"
	USERPERMISSION_CreateClipFolder          UserPermission = "create_clip_folder"
	USERPERMISSION_EditClipFolder            UserPermission = "edit_clip_folder"
	USERPERMISSION_DeleteClipFolder          UserPermission = "delete_clip_folder"
)
