package config

type Config struct {
	EmailTemplatePath    string `mapstructure:"EMAIL_TEMPLATE_PATH" json:"email_template_path"`
	CustomerDataFilePath string `mapstructure:"CUSTOMER_DATA_FILE_PATH" json:"customer_data_path"`
	OutputEmailPath      string `mapstructure:"OUTPUT_EMAIL_PATH" json:"out_email_path"`
	ErrorPath            string `mapstructure:"ERROR_PATH" json:"error_path"`
	SmtpHost             string `mapstructure:"SMTP_HOST" json:"smtp_host"`
	SmtpPort             string `mapstructure:"SMTP_PORT" json:"smtp_port"`
	Password             string `mapstructure:"PASSWORD" json:"password"`
}
