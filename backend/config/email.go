package config

import "github.com/spf13/viper"

type EmailConfig struct {
	Host         string `mapstructure:"SMTP_HOST"`
	Port         int    `mapstructure:"SMTP_PORT"`
	SenderName   string `mapstructure:"SMTP_SENDER_NAME"`
	AuthEmail    string `mapstructure:"SMTP_AUTH_EMAIL"`
	AuthPassword string `mapstructure:"SMTP_AUTH_PASSWORD"`
}

func NewEmailConfig() (*EmailConfig, error) {
	viper.AutomaticEnv()

	viper.BindEnv("SMTP_HOST")
	viper.BindEnv("SMTP_PORT")
	viper.BindEnv("SMTP_SENDER_NAME")
	viper.BindEnv("SMTP_AUTH_EMAIL")
	viper.BindEnv("SMTP_AUTH_PASSWORD")

	config := EmailConfig{
		Host:         viper.GetString("SMTP_HOST"),
		Port:         viper.GetInt("SMTP_PORT"),
		SenderName:   viper.GetString("SMTP_SENDER_NAME"),
		AuthEmail:    viper.GetString("SMTP_AUTH_EMAIL"),
		AuthPassword: viper.GetString("SMTP_AUTH_PASSWORD"),
	}

	return &config, nil
}
