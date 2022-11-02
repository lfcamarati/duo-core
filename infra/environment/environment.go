package environment

import (
	"errors"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(errors.New("Erro fatal no arquivo de configuração: " + err.Error()))
	}
}
