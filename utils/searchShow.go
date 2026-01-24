package utils

import (
	"fmt"
	"strings"
	"pplx2api/config"
)

// cleanURL 对 URL 进行简单的转义处理，防止破坏 Markdown 结构
func cleanURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	// 替换空格为 %20
	rawURL = strings.ReplaceAll(rawURL, " ", "%20")
	// 替换 Markdown 链接语法的关键字符
	rawURL = strings.ReplaceAll(rawURL, "(", "%28")
	rawURL = strings.ReplaceAll(rawURL, ")", "%29")
	return rawURL
}

// cleanTitle 清洗标题，防止破坏 Markdown 结构
func cleanTitle(title string) string {
	// 移除换行符
	title = strings.ReplaceAll(title, "\n", " ")
	title = strings.ReplaceAll(title, "\r", " ")
	// 转义方括号
	title = strings.ReplaceAll(title, "[", "\\[")
	title = strings.ReplaceAll(title, "]", "\\]")
	return strings.TrimSpace(title)
}

func searchShowDetails(index int, title, url, snippet string) string {
	// 优化：直接把标题作为链接文本，去掉 snippet（摘要），因为文末引用通常只需要标题
	// 如果需要摘要，可以保留 snippet
	return fmt.Sprintf("[%d] [%s](%s)", index, cleanTitle(title), cleanURL(url))
}

func searchShowCompatible(index int, title, url, snippet string) string {
	// 优化：更紧凑的 Markdown 列表格式
	// 格式：1. [网页标题](URL) - 摘要前50个字...
	
	// 截断过长的 snippet
	if len([]rune(snippet)) > 50 {
		runeSnippet := []rune(snippet)
		snippet = string(runeSnippet[:50]) + "..."
	}
	
	// 移除 snippet 中的换行符，保持一行
	snippet = strings.ReplaceAll(snippet, "\n", " ")
	
	return fmt.Sprintf("%d. [%s](%s) - %s", index, cleanTitle(title), cleanURL(url), snippet)
}

func SearchShow(index int, title, url, snippet string) string {
	index++
	// 初步去除首尾空格
	url = strings.TrimSpace(url)
	if url == "" {
		return ""
	}

	if config.ConfigInstance.SearchResultCompatible {
		return searchShowCompatible(index, title, url, snippet)
	}
	return searchShowDetails(index, title, url, snippet)
}
