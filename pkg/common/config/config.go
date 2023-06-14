package config

//This line specifies that the code in this file belongs to the config package.

import "github.com/spf13/viper"

/*This line imports the "viper" package, which is a popular Go library used for configuration management.*/

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

/*This code defines a struct type called Config. It represents the configuration settings that will be loaded and stored using Viper. The Port field is associated with the environment variable PORT, and the DBUrl field is associated with the environment variable DB_URL. The mapstructure tag is used to specify the mapping between the struct fields and the environment variables.
Port, DBUrl are struct fields and PORT, DB_URL are env variables*/

func LoadConfig() (c Config, err error) {

	//This function, LoadConfig, loads the configuration settings using Viper.

	viper.AddConfigPath("./pkg/common/config/envs")

	//This line adds the directory path ./pkg/common/config/envs to the search paths for the configuration file.

	viper.SetConfigName("dev")

	//This line sets the name of the configuration file to be loaded as "dev".

	viper.SetConfigType("env")

	//This line sets the configuration file type to "env".

	viper.AutomaticEnv()

	//This line enables automatic binding of environment variables to the corresponding fields in the Config struct.

	err = viper.ReadInConfig()

	//This line reads the configuration file specified earlier (dev.env) and populates the Viper configuration object with the values.

	if err != nil {
		return
	}

	//This block of code checks if there was an error while reading the configuration file.

	err = viper.Unmarshal(&c)

	//This line unmarshals the configuration values from the Viper configuration object into the Config struct c.

	return
}
