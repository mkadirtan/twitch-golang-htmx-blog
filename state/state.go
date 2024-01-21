package state

import "github.com/mailgun/raymond/v2"

type Site struct {
	Title             string
	Url               string
	MembersEnabled    bool `handlebars:"members_enabled"`
	Locale            string
	Logo              string
	MembersInviteOnly bool `handlebars:"members_invite_only"`
}

type Custom struct {
	ColorScheme          string `handlebars:"color_scheme"`
	NavigationLayout     string `handlebars:"navigation_layout"`
	TitleFont            string `handlebars:"title_font"`
	BodyFont             string `handlebars:"body_font"`
	ShowPublicationCover bool   `handlebars:"show_publication_cover"`
	HeaderStyle          string `handlebars:"header_style"`
}

type TplData struct {
	Site       Site              `handlebars:"@site"`
	Custom     Custom            `handlebars:"@custom"`
	MetaTitle  string            `handlebars:"meta_title"`
	GhostHead  string            `handlebars:"ghost_head"`
	BodyClass  string            `handlebars:"body_class"`
	Navigation string            `handlebars:"navigation"`
	Date       string            `handlebars:"date"`
	GhostFoot  string            `handlebars:"ghost_foot"`
	Body       *raymond.Template `handlebars:"body"`
}
