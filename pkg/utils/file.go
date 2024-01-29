package utils

import (
	"encoding/json"
	"os"
)

const cacheFileName = "resume.json"
const optFileName = "config.json"

func Write(datas interface{}, filePath string) error {
	files, err := json.MarshalIndent(datas, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(filePath, files, 0644); err != nil {
		return err
	}
	return nil
}

func Read(datas any, filePath string) error {
	b, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(b, datas); err != nil {
		return err
	}
	return nil
}

func WriteResume(datas interface{}) error {
	return Write(datas, cacheFileName)
}

func ReadResume(datas any) error {
	return Read(datas, cacheFileName)
}

func WriteOpt(datas interface{}) error {
	return Write(datas, optFileName)
}

func ReadOpt(datas any) error {
	return Read(datas, optFileName)
}
