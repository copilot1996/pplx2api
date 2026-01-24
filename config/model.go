package config

var ModelReverseMap = map[string]string{}
var ModelMap = map[string]string{
	"gpt-5.2":                 "gpt52",
	"gpt-5.1":                 "gpt51",
	"claude-4.5-sonnet":       "claude45sonnet",
	"claude-4.5-sonnet-think": "claude45sonnetthinking",
	"claude-4.5-opus-think":   "claude45opusthinking",
	"gemini-3-pro":            "gemini30pro",
	"gemini-3-flash":          "gemini3flash",
	"kimi-k2":                 "kimik2thinking",
	"grok-4.1":                "grok41",
	"sonar":                   "turbo",
}
var MaxModelMap = map[string]string{
	"claude-4.5-opus-think": "claude45opusthinking",
}

// Get returns the value for the given key from the ModelMap.
// If the key doesn't exist, it returns the provided default value.
func ModelMapGet(key string, defaultValue string) string {
	if value, exists := ModelMap[key]; exists {
		return value
	}
	return defaultValue
}

// GetReverse returns the value for the given key from the ModelReverseMap.
// If the key doesn't exist, it returns the provided default value.
func ModelReverseMapGet(key string, defaultValue string) string {
	if value, exists := ModelReverseMap[key]; exists {
		return value
	}
	return defaultValue
}

var ResponseModels []map[string]string

func init() {
	// 构建反向映射
	for k, v := range ModelMap {
		ModelReverseMap[v] = k
	}
	buildResponseModels()
}

// buildResponseModels 构建响应模型列表
func buildResponseModels() {
	ResponseModels = make([]map[string]string, 0, len(ModelMap)*2)

	for modelID := range ModelMap {
		// 添加普通模型
		ResponseModels = append(ResponseModels, map[string]string{
			"id": modelID,
		})

		// 添加搜索模型
		ResponseModels = append(ResponseModels, map[string]string{
			"id": modelID + "-search",
		})
	}
}
