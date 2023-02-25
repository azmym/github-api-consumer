package domain

type (
	Repo struct {
		Id           int             `json:"id"`
		NodeId       string          `json:"node_id"`
		Name         string          `json:"name"`
		FullName     string          `json:"full_name"`
		Owner        RepoOwner       `json:"owner"`
		Description  string          `json:"description"`
		AllowForking bool            `json:"allow_forking"`
		Visibility   string          `json:"visibility"`
		Permissions  RepoPermissions `json:"permissions"`
		Languages    Languages       `json:"languages",omitempty`
	}

	RepoOwner struct {
		Login   string `json:"login"`
		Id      int    `json:"id"`
		HtmlUrl string `json:"html_url"`
	}

	RepoPermissions struct {
		Admin    bool `json:"admin"`
		Maintain bool `json:"maintain"`
		Push     bool `json:"push"`
		Triage   bool `json:"triage"`
		Pull     bool `json:"pull"`
	}

	User struct {
		UserName string `json:"login"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Location string `json:"location"`
	}

	Languages map[string]int
)
