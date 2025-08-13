package about

import "github.com/aandrku/personal-website/pkg/services/markdown"

func newDefaultAboutInfo() AboutInfo {
	return AboutInfo{
		Name:        "Andrii Sozonik",
		AvatarURL:   "/uploads/avatar.png",
		DescShort:   "self-taught developer",
		Description: "## Hi, I am andrii",
	}
}

type AboutInfo struct {
	Name        string `json:"name"`
	AvatarURL   string `json:"avatarURL"`
	DescShort   string `json:"descShort"`
	Description string `json:"description"`
}

func (i AboutInfo) DescriptionHTML() (string, error) {
	d, err := markdown.ToHTML(i.Description)
	if err != nil {
		return d, err
	}

	return d, nil
}
