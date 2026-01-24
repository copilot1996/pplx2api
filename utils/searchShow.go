package utils

import (
	"fmt"
	"strings"
	"pplx2api/config"
)

func searchShowDetails(index int, title, url, snippet string) string {
	// 优化：直接把标题作为链接文本，去掉 snippet（摘要），因为文末引用通常只需要标题
	// 如果需要摘要，可以保留 snippet
	return fmt.Sprintf("[%d] [%s](%s)", index, title, url)
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
	
	return fmt.Sprintf("%d. [%s](%s) - %s", index, title, url, snippet)
}

func SearchShow(index int, title, url, snippet string) string {
	index++
	url = strings.TrimSpace(url)
	if url == "" {
		return ""
	}

	if config.ConfigInstance.SearchResultCompatible {
		return searchShowCompatible(index, title, url, snippet)
	}
	return searchShowDetails(index, title, url, snippet)
}
