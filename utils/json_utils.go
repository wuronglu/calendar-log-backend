package utils

import (
	"encoding/json"
	"os"
)

// ReadJSONFile 读取 JSON 文件并解析为 map
func ReadJSONFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := make(map[string]string)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&content); err != nil {
		return nil, err
	}

	return content, nil
}

// WriteJSONFile 将数据写入 JSON 文件
func WriteJSONFile(filePath string, data map[string]string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
