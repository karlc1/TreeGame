package config

type Config struct {
	ScreenHeight float64
	ScreenWidth  float64
	TargetFPS    float64
}

func Default() *Config {
	return &Config{
		ScreenHeight: 800,
		ScreenWidth:  1200,
		TargetFPS:    60,
	}
}
