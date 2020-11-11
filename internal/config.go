package internal

//Config struct for holding config for exporter and Gitlab
type Config struct {
	Port      string
	Host      string
	User      string
	Password  string
	Database  string
	Retention string
	SSLMode   string
	Metrics   string
}
