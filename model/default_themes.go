// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

var DefaultThemes = map[string]*Theme{
	"default": {
		Id:                         "default",
		DisplayName:                "Mattermost",
		SidebarBackground:          "#145dbf",
		SidebarText:                "#ffffff",
		SidebarUnreadText:          "#ffffff",
		SidebarTextHoverBackground: "#4578bf",
		SidebarTextActiveBorder:    "#579eff",
		SidebarTextActiveColor:     "#ffffff",
		SidebarHeaderBackground:    "#1153ab",
		SidebarHeaderTextColor:     "#ffffff",
		OnlineIndicator:            "#06d6a0",
		AwayIndicator:              "#ffbc42",
		DndIndicator:               "#f74343",
		MentionBackground:          "#ffffff",
		MentionColor:               "#145dbf",
		CenterChannelBackground:    "#ffffff",
		CenterChannelColor:         "#3d3c40",
		NewMessageSeparator:        "#ff8800",
		LinkColor:                  "#2389d7",
		ButtonBackground:           "#166de0",
		ButtonColor:                "#ffffff",
		ErrorTextColor:             "#fd5960",
		MentionHighlightBackground: "#ffe577",
		MentionHighlightLink:       "#166de0",
		CodeTheme:                  "github",
	},
	"organization": {
		Id:                         "organization",
		DisplayName:                "Organization",
		SidebarBackground:          "#2071a7",
		SidebarText:                "#ffffff",
		SidebarUnreadText:          "#ffffff",
		SidebarTextHoverBackground: "#136197",
		SidebarTextActiveBorder:    "#7ab0d6",
		SidebarTextActiveColor:     "#ffffff",
		SidebarHeaderBackground:    "#2f81b7",
		SidebarHeaderTextColor:     "#ffffff",
		OnlineIndicator:            "#7dbe00",
		AwayIndicator:              "#dcbd4e",
		DndIndicator:               "#ff6a6a",
		MentionBackground:          "#fbfbfb",
		MentionColor:               "#2071f7",
		CenterChannelBackground:    "#f2f4f8",
		CenterChannelColor:         "#333333",
		NewMessageSeparator:        "#ff8800",
		LinkColor:                  "#2f81b7",
		ButtonBackground:           "#1dacfc",
		ButtonColor:                "#ffffff",
		ErrorTextColor:             "#a94442",
		MentionHighlightBackground: "#f3e197",
		MentionHighlightLink:       "#2f81b7",
	},
	"mattermostDark": {
		Id:                         "mattermostDark",
		DisplayName:                "Mattermost Dark",
		SidebarBackground:          "#1b2c3e",
		SidebarText:                "#ffffff",
		SidebarUnreadText:          "#ffffff",
		SidebarTextHoverBackground: "#4a5664",
		SidebarTextActiveBorder:    "#66b9a7",
		SidebarTextActiveColor:     "#ffffff",
		SidebarHeaderBackground:    "#1b2c3e",
		SidebarHeaderTextColor:     "#ffffff",
		OnlineIndicator:            "#65dcc8",
		AwayIndicator:              "#c1b966",
		DndIndicator:               "#e81023",
		MentionBackground:          "#b74a4a",
		MentionColor:               "#ffffff",
		CenterChannelBackground:    "#2f3e4e",
		CenterChannelColor:         "#dddddd",
		NewMessageSeparator:        "#5de5da",
		LinkColor:                  "#a4ffeb",
		ButtonBackground:           "#4cbba4",
		ButtonColor:                "#ffffff",
		ErrorTextColor:             "#ff6461",
		MentionHighlightBackground: "#984063",
		MentionHighlightLink:       "#a4ffeb",
		CodeTheme:                  "solarized-dark",
	},
	"windows10": {
		Id:                         "windows10",
		DisplayName:                "Windows Dark",
		SidebarBackground:          "#171717",
		SidebarText:                "#ffffff",
		SidebarUnreadText:          "#ffffff",
		SidebarTextHoverBackground: "#302e30",
		SidebarTextActiveBorder:    "#196caf",
		SidebarTextActiveColor:     "#ffffff",
		SidebarHeaderBackground:    "#1f1f1f",
		SidebarHeaderTextColor:     "#ffffff",
		OnlineIndicator:            "#399fff",
		AwayIndicator:              "#c1b966",
		DndIndicator:               "#e81023",
		MentionBackground:          "#0177e7",
		MentionColor:               "#ffffff",
		CenterChannelBackground:    "#1f1f1f",
		CenterChannelColor:         "#dddddd",
		NewMessageSeparator:        "#cc992d",
		LinkColor:                  "#0d93ff",
		ButtonBackground:           "#0177e7",
		ButtonColor:                "#ffffff",
		ErrorTextColor:             "#ff6461",
		MentionHighlightBackground: "#784098",
		MentionHighlightLink:       "#a4ffeb",
		CodeTheme:                  "monokai",
	},
}
