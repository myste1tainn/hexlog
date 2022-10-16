package log

type Config struct {
	// Level indicate minimum logging level that will be logged out to the writer default is debug
	Level string `mapstructure:"level"`

	// ReportCaller set this to true to make it report file and line number of the log message
	ReportCaller bool `mapstructure:"reportCaller"`

	// Infra indicate which infrastructure environment the logger is running in, specify `InfraLocal`
	// to make the logging output console style logging rather than JSON, all `Cloud` variants will output JSON
	Infra string `mapstructure:"infra"`
}

var DefaultConfig = Config{
	Level:        "debug",
	ReportCaller: true,
	Infra:        "local",
}
