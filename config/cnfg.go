package config

var EnvVariable = new(EnvConf)

type (
	EnvConf struct {
		GithubToken string `envconfig:"GITHUB_TOKEN" required:"true"`
	}
)
