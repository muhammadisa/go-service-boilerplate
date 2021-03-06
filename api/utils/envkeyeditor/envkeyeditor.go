package envkeyeditor

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/mapsmerger"
)

// EnvKeyEditor edit key inside .env file
func EnvKeyEditor(targetKey string, newKeyValue string) (string, string, error) {
	if targetKey == "" || newKeyValue == "" {
		return "", "", fmt.Errorf("TargetKey or NewKeyValue can't be blank string")
	}

	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return "", "", err
	}
	lastKeyValue := myEnv[targetKey]

	debugChanged, err := godotenv.Unmarshal(
		fmt.Sprintf(
			"%s=%s",
			targetKey,
			newKeyValue,
		),
	)
	if err != nil {
		return "", "", err
	}

	maps := []map[string]string{myEnv, debugChanged}
	mergedMaps := mapsmerger.MapsMerger(maps...)
	err = godotenv.Write(mergedMaps, ".env")
	if err != nil {
		return "", "", err
	}

	return lastKeyValue, newKeyValue, nil
}
