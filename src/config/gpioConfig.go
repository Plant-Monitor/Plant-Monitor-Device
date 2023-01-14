package config

type GpioConfig struct{}

func GetGpioConfigInstance() *GpioConfig {
	return &GpioConfig{}
}
