// Code generated by generate/interfaces; DO NOT EDIT.

package disgord

func (g *GetInviteParams) URLQueryString() string {
	params := make(paramHolder)

	if !(g.WithMemberCount == false) {
		params["with_count"] = g.WithMemberCount
	}

	return params.URLQueryString()
}

func (g *GetMessagesParams) URLQueryString() string {
	params := make(paramHolder)

	if !(g.Around == 0) {
		params["around"] = g.Around
	}

	if !(g.Before == 0) {
		params["before"] = g.Before
	}

	if !(g.After == 0) {
		params["after"] = g.After
	}

	if !(g.Limit == 0) {
		params["limit"] = g.Limit
	}

	return params.URLQueryString()
}

func (g *GetReactionURLParams) URLQueryString() string {
	params := make(paramHolder)

	if !(g.Before == 0) {
		params["before"] = g.Before
	}

	if !(g.After == 0) {
		params["after"] = g.After
	}

	if !(g.Limit == 0) {
		params["limit"] = g.Limit
	}

	return params.URLQueryString()
}

func (g *GetGuildMembersParams) URLQueryString() string {
	params := make(paramHolder)

	if !(g.After == 0) {
		params["after"] = g.After
	}

	if !(g.Limit == 0) {
		params["limit"] = g.Limit
	}

	return params.URLQueryString()
}

func (b *BanMemberParams) URLQueryString() string {
	params := make(paramHolder)

	if !(b.DeleteMessageDays == 0) {
		params["delete_message_days"] = b.DeleteMessageDays
	}

	if !(b.Reason == "") {
		params["reason"] = b.Reason
	}

	return params.URLQueryString()
}

func (p *pruneMembersParams) URLQueryString() string {
	params := make(paramHolder)

	params["days"] = p.Days

	params["compute_prune_count"] = p.ComputePruneCount

	return params.URLQueryString()
}

func (g *GetCurrentUserGuildsParams) URLQueryString() string {
	params := make(paramHolder)

	if !(g.Before == 0) {
		params["before"] = g.Before
	}

	if !(g.After == 0) {
		params["after"] = g.After
	}

	if !(g.Limit == 0) {
		params["limit"] = g.Limit
	}

	return params.URLQueryString()
}

func (e *execWebhookParams) URLQueryString() string {
	params := make(paramHolder)

	params["wait"] = e.Wait

	return params.URLQueryString()
}
