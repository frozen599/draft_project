package config

type Config struct {
	EmailTemplatePath    string `mapstructure:"EMAIL_TEMPLATE_PATH"`
	CustomerDataFilePath string `mapstructure:"CUSTOMER_DATA_FILE_PATH"`
	OutputEmailPath      string `mapstructure:"OUTPUT_EMAIL_PATH"`
	ErrorPath            string `mapstructure:"ERROR_PATH"`
}
