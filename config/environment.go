package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func AutoLoad() error {

	file, err := os.Open(".env")

	if err != nil {
		err = errors.New("error when open .env")
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		key, value, err := parseLine(scanner.Text())
		if err == nil {
			if os.Getenv(key) == "" {
				os.Setenv(key, value)
			}
		} else {
			return err
		}
	}

	return err
}

func parseLine(line string) (key string, value string, err error) {

	if len(line) == 0 {
		err = errors.New("zero length string")
		return
	}

	splitString := strings.SplitN(line, "=", 2)

	if len(splitString) != 2 {
		err = errors.New("Can't separate key from value")
		return
	}

	key = splitString[0]
	key = strings.Trim(key, " ")

	value = splitString[1]
	value = strings.Trim(value, " ")

	if strings.Count(value, "\"") == 2 || strings.Count(value, "'") == 2 {
		value = strings.Trim(value, "\"'")

		value = strings.Replace(value, "\\\"", "\"", -1)
		value = strings.Replace(value, "\\n", "\n", -1)
	}

	return
}
