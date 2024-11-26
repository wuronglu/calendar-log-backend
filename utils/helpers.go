package utils

import (
	"encoding/json"
	"io"
	"os"
)

// ReadJSONFile 读取 JSON 文件并解析为 map
func ReadJSONFile(filePath string) (map[string]interface{}, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 内容
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}

	return result, nil
}
