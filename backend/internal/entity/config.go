package entity

type Config struct {
	Github `yaml:"GITHUB"`
	App    `yaml:"APP"`
}

type App struct {
	Port string `yaml:"port"`
}
type Github struct {
	BaseUrl      string `yaml:"base_url"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
