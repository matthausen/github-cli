package model

type(
	Profile struct {
		Login string `json:"login"`
		Id int32 `json:"id"`
		NodeId string `json:"node_id"`
		Avatar string `json:"avatar_url"`
		Name string `json:"name"`
		Company string `json:"company"`
		Location string `json:"location"`
		Email string `json:"email"`
		Bio string `json:"bio"`
		TwitterHandle string `json:"twitter_username"`
		PublicRepos int `json:"public_repos"`
	}

	Repository struct {
		HTMLUrl string `json:"html_url"`
		Description string `json:"description"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		PushedAt string `json:"pushed_at"`
		Size int `json:"size"`
		Language string `json:"language"`
	}

	CreateRepoResponse struct {
		Id int32 `json:"id"`
		NodeId string `json:"node_id"`
		Name string `json:"name"`
		FullName string `json:"full_name"`
		Private bool `json:"private"`
		HTMLUrl string `json:"html_url"`
		Description string `json:"description"`
		Owner Profile `json:"owner"`
		GitUrl string `json:"git_url"`
		SSHUrl string `json:"ssh_url"`
		CloneUrl string `json:"clone_url"`
		Language string `json:"language"`
	}
)



