package main

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int16  `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBDatabase string `mapstructure:"DB_NAME"`
}

func LoadConfig(configPath string, configFileName string) (config Config) {
	//viper.SetDefault("DB_HOST", "database")
	//viper.SetDefault("DB_PORT", int16(5432))
	//viper.SetDefault("DB_USER", "provisioning")
	//viper.SetDefault("DB_NAME", "provisioning")
	//viper.SetDefault("DB_PASSWORD", "")
	//
	//viper.SetConfigName(configFileName)
	//viper.AddConfigPath(configPath)
	//viper.SetConfigType("env")
	//viper.AutomaticEnv()
	//
	//err := viper.ReadInConfig()
	//if err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		log.Printf("could not find config file: %s/%s\n", configPath, configFileName)
	//	} else {
	//		log.Fatalf("error during config loading: %v", err)
	//	}
	//}
	//
	//err = viper.Unmarshal(&config)
	//if err != nil {
	//	panic(err)
	//}
	return
}
