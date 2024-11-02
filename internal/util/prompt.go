package util

import "github.com/tmc/langchaingo/prompts"

func CreatePrompt() prompts.ChatPromptTemplate {
	return prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("", nil),
		prompts.NewHumanMessagePromptTemplate("{{ .text }}", []string{
			"text",
		}),
	})
}

//func CreatePrompt() prompts.ChatPromptTemplate {
//	//return prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
//	//	prompts.NewSystemMessagePromptTemplate("", nil),
//	//	prompts.NewHumanMessagePromptTemplate("{{ .text }}", []string{
//	//		"text",
//	//	}),
//	//})
//	return prompts.NewChatPromptTemplate(
//		).
//}
