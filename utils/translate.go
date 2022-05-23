package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var (
	Dictionaries = map[string]map[string]string{}
)

func InitTranslationDictionary() error {
	rootPath := "utils/data/"
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()
		dictionaryName := fileName[:len(fileName)-5]

		f, err := os.Open(rootPath + fileName)
		if err != nil {
			return err
		}

		singleDictionary := make(map[string]string)
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&singleDictionary)
		if err != nil {
			return err
		}

		Dictionaries[dictionaryName] = singleDictionary
	}
	return nil
}

func TranslateText(dictionaryName string, reqText string) (string, error) {
	if singleDictionary, ok := Dictionaries[dictionaryName]; ok {
		if resText, ok := singleDictionary[reqText]; ok {
			return resText, nil
		}
		return reqText, nil
	}

	return reqText, errors.New("The dictionary named: " + dictionaryName + ", in the \"utils/data/\" path, was not found")
}
