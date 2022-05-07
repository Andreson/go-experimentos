package boostrap

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func loadEnvInt(name string, defaultValue int) int {

	if value := os.Getenv(name); value == "" {
		return defaultValue
	} else {
		valueCast, _ := strconv.Atoi(value)
		return valueCast
	}
}

func loadEnvBool(name string, defaultValue string) bool {

	readValue, err := strconv.ParseBool(loadEnvString(name, defaultValue))

	if err != nil {
		log.Errorf("Erro ao carregar env  configurado %s | err : ", name, err.Error())
		readValue, _ = strconv.ParseBool(defaultValue)
	}
	return readValue
}

func loadEnvString(name string, defaultValue string) string {

	if value := os.Getenv(name); value == "" {
		return defaultValue
	} else {
		return value
	}
}

func loadPathParamIndex(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
