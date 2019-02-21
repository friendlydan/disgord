package disgord

// Warning: This file has been automatically generated by generate/restbuilders/main.go
// DO NOT EDIT! This file is overwritten at "go generate"
// This file holds all the basic RESTBuilder methods a builder is expected to.

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *guildAuditLogsBuilder) IgnoreCache() *guildAuditLogsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *guildAuditLogsBuilder) CancelOnRatelimit() *guildAuditLogsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *guildAuditLogsBuilder) URLParam(name string, v interface{}) *guildAuditLogsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *guildAuditLogsBuilder) Set(name string, v interface{}) *guildAuditLogsBuilder {
	b.r.body[name] = v
	return b
}

func (b *guildAuditLogsBuilder) SetUserID(userID Snowflake) *guildAuditLogsBuilder {
	b.r.addPrereq(userID.Empty(), "userID can not be 0")
	b.r.param("user_id", userID)
	return b
}

func (b *guildAuditLogsBuilder) SetActionType(actionType uint) *guildAuditLogsBuilder {
	b.r.param("action_type", actionType)
	return b
}

func (b *guildAuditLogsBuilder) SetBefore(before Snowflake) *guildAuditLogsBuilder {
	b.r.addPrereq(before.Empty(), "before can not be 0")
	b.r.param("before", before)
	return b
}

func (b *guildAuditLogsBuilder) SetLimit(limit int) *guildAuditLogsBuilder {
	b.r.param("limit", limit)
	return b
}

func (b *guildAuditLogsBuilder) Execute() (log *AuditLog, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*AuditLog), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createGuildEmojiBuilder) IgnoreCache() *createGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createGuildEmojiBuilder) CancelOnRatelimit() *createGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createGuildEmojiBuilder) URLParam(name string, v interface{}) *createGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createGuildEmojiBuilder) Set(name string, v interface{}) *createGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

func (b *createGuildEmojiBuilder) SetRoles(roles []Snowflake) *createGuildEmojiBuilder {
	b.r.param("roles", roles)
	return b
}

func (b *createGuildEmojiBuilder) Execute() (emoji *Emoji, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Emoji), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getGuildEmojiBuilder) IgnoreCache() *getGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getGuildEmojiBuilder) CancelOnRatelimit() *getGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getGuildEmojiBuilder) URLParam(name string, v interface{}) *getGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getGuildEmojiBuilder) Set(name string, v interface{}) *getGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

func (b *getGuildEmojiBuilder) Execute() (emoji *Emoji, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Emoji), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getGuildEmojisBuilder) IgnoreCache() *getGuildEmojisBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getGuildEmojisBuilder) CancelOnRatelimit() *getGuildEmojisBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getGuildEmojisBuilder) URLParam(name string, v interface{}) *getGuildEmojisBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getGuildEmojisBuilder) Set(name string, v interface{}) *getGuildEmojisBuilder {
	b.r.body[name] = v
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *modifyGuildEmojiBuilder) IgnoreCache() *modifyGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *modifyGuildEmojiBuilder) CancelOnRatelimit() *modifyGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *modifyGuildEmojiBuilder) URLParam(name string, v interface{}) *modifyGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *modifyGuildEmojiBuilder) Set(name string, v interface{}) *modifyGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

func (b *modifyGuildEmojiBuilder) SetName(name string) *modifyGuildEmojiBuilder {
	b.r.param("name", name)
	return b
}

func (b *modifyGuildEmojiBuilder) SetRoles(roles []Snowflake) *modifyGuildEmojiBuilder {
	b.r.param("roles", roles)
	return b
}

func (b *modifyGuildEmojiBuilder) Execute() (emoji *Emoji, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Emoji), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *deleteInviteBuilder) IgnoreCache() *deleteInviteBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *deleteInviteBuilder) CancelOnRatelimit() *deleteInviteBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *deleteInviteBuilder) URLParam(name string, v interface{}) *deleteInviteBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *deleteInviteBuilder) Set(name string, v interface{}) *deleteInviteBuilder {
	b.r.body[name] = v
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getInviteBuilder) IgnoreCache() *getInviteBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getInviteBuilder) CancelOnRatelimit() *getInviteBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getInviteBuilder) URLParam(name string, v interface{}) *getInviteBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getInviteBuilder) Set(name string, v interface{}) *getInviteBuilder {
	b.r.body[name] = v
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *basicBuilder) IgnoreCache() *basicBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *basicBuilder) CancelOnRatelimit() *basicBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *basicBuilder) URLParam(name string, v interface{}) *basicBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *basicBuilder) Set(name string, v interface{}) *basicBuilder {
	b.r.body[name] = v
	return b
}

func (b *basicBuilder) Execute() (err error) {
	_, err = b.r.execute()
	return
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *modifyGuildRoleBuilder) IgnoreCache() *modifyGuildRoleBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *modifyGuildRoleBuilder) CancelOnRatelimit() *modifyGuildRoleBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *modifyGuildRoleBuilder) URLParam(name string, v interface{}) *modifyGuildRoleBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *modifyGuildRoleBuilder) Set(name string, v interface{}) *modifyGuildRoleBuilder {
	b.r.body[name] = v
	return b
}

func (b *modifyGuildRoleBuilder) SetName(name string) *modifyGuildRoleBuilder {
	b.r.param("name", name)
	return b
}

func (b *modifyGuildRoleBuilder) SetPermissions(permissions uint64) *modifyGuildRoleBuilder {
	b.r.param("permissions", permissions)
	return b
}

func (b *modifyGuildRoleBuilder) SetColor(color uint) *modifyGuildRoleBuilder {
	b.r.param("color", color)
	return b
}

func (b *modifyGuildRoleBuilder) SetHoist(hoist bool) *modifyGuildRoleBuilder {
	b.r.param("hoist", hoist)
	return b
}

func (b *modifyGuildRoleBuilder) SetMentionable(mentionable bool) *modifyGuildRoleBuilder {
	b.r.param("mentionable", mentionable)
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createDMBuilder) IgnoreCache() *createDMBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createDMBuilder) CancelOnRatelimit() *createDMBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createDMBuilder) URLParam(name string, v interface{}) *createDMBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createDMBuilder) Set(name string, v interface{}) *createDMBuilder {
	b.r.body[name] = v
	return b
}

func (b *createDMBuilder) Execute() (channel *Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Channel), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createGroupDMBuilder) IgnoreCache() *createGroupDMBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createGroupDMBuilder) CancelOnRatelimit() *createGroupDMBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createGroupDMBuilder) URLParam(name string, v interface{}) *createGroupDMBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createGroupDMBuilder) Set(name string, v interface{}) *createGroupDMBuilder {
	b.r.body[name] = v
	return b
}

func (b *createGroupDMBuilder) Execute() (channel *Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Channel), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getCurrentUserGuildsBuilder) IgnoreCache() *getCurrentUserGuildsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getCurrentUserGuildsBuilder) CancelOnRatelimit() *getCurrentUserGuildsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getCurrentUserGuildsBuilder) URLParam(name string, v interface{}) *getCurrentUserGuildsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getCurrentUserGuildsBuilder) Set(name string, v interface{}) *getCurrentUserGuildsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getCurrentUserGuildsBuilder) SetBefore(before Snowflake) *getCurrentUserGuildsBuilder {
	b.r.addPrereq(before.Empty(), "before can not be 0")
	b.r.param("before", before)
	return b
}

func (b *getCurrentUserGuildsBuilder) SetAfter(after Snowflake) *getCurrentUserGuildsBuilder {
	b.r.addPrereq(after.Empty(), "after can not be 0")
	b.r.param("after", after)
	return b
}

func (b *getCurrentUserGuildsBuilder) SetLimit(limit int) *getCurrentUserGuildsBuilder {
	b.r.param("limit", limit)
	return b
}

func (b *getCurrentUserGuildsBuilder) Execute() (guilds []*Guild, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*Guild)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserBuilder) IgnoreCache() *getUserBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserBuilder) CancelOnRatelimit() *getUserBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserBuilder) URLParam(name string, v interface{}) *getUserBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserBuilder) Set(name string, v interface{}) *getUserBuilder {
	b.r.body[name] = v
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserConnectionsBuilder) IgnoreCache() *getUserConnectionsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserConnectionsBuilder) CancelOnRatelimit() *getUserConnectionsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserConnectionsBuilder) URLParam(name string, v interface{}) *getUserConnectionsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserConnectionsBuilder) Set(name string, v interface{}) *getUserConnectionsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getUserConnectionsBuilder) Execute() (cons []*UserConnection, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*UserConnection)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserDMsBuilder) IgnoreCache() *getUserDMsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserDMsBuilder) CancelOnRatelimit() *getUserDMsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserDMsBuilder) URLParam(name string, v interface{}) *getUserDMsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserDMsBuilder) Set(name string, v interface{}) *getUserDMsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getUserDMsBuilder) Execute() (channels []*Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*Channel)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *modifyCurrentUserBuilder) IgnoreCache() *modifyCurrentUserBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *modifyCurrentUserBuilder) CancelOnRatelimit() *modifyCurrentUserBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *modifyCurrentUserBuilder) URLParam(name string, v interface{}) *modifyCurrentUserBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *modifyCurrentUserBuilder) Set(name string, v interface{}) *modifyCurrentUserBuilder {
	b.r.body[name] = v
	return b
}

func (b *modifyCurrentUserBuilder) SetUsername(username string) *modifyCurrentUserBuilder {
	b.r.param("username", username)
	return b
}

func (b *modifyCurrentUserBuilder) SetAvatar(avatar string) *modifyCurrentUserBuilder {
	b.r.param("avatar", avatar)
	return b
}

func (b *modifyCurrentUserBuilder) Execute() (user *User, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*User), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *listVoiceRegionsBuilder) IgnoreCache() *listVoiceRegionsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *listVoiceRegionsBuilder) CancelOnRatelimit() *listVoiceRegionsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *listVoiceRegionsBuilder) URLParam(name string, v interface{}) *listVoiceRegionsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *listVoiceRegionsBuilder) Set(name string, v interface{}) *listVoiceRegionsBuilder {
	b.r.body[name] = v
	return b
}
