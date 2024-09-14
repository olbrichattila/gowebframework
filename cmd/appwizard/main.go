package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	envFileName = ".env"
)

type envData struct {
	key   string
	value string
}

type answer struct {
	value        string
	nextQuestion *question
}
type answers map[string]answer
type question struct {
	key           string
	prompt        string
	defaultAnswer string
	mandatory     bool
	answers       answers
	nextQuestion  *question
}

func main() {
	databaseWizard()
}

func databaseWizard() {
	envContent := getEnvContent()
	currentQuestion := appUrlQuestion
	responses := make([]envData, 0)
	for {
		currentValue := lookupValue(envContent, currentQuestion.key)
		if currentValue == "" {
			currentValue = currentQuestion.defaultAnswer
		}
		answer := selection(currentQuestion, currentValue)
		fmt.Println("")
		if answer == nil {
			break
		}

		if currentQuestion.key != "" && answer.value != "" {
			responses = append(responses, envData{key: currentQuestion.key, value: answer.value})
		}

		if answer.nextQuestion != nil {
			currentQuestion = *answer.nextQuestion
			continue
		}

		break
	}

	storages := getStorages(responses)
	for _, storageName := range storages {
		if storageQuestion, ok := storageQuestionMap[storageName]; ok {
			if storageQuestion == nil {
				continue
			}
			currentQuestion := storageQuestion
			for {
				currentValue := lookupValue(envContent, currentQuestion.key)
				if currentValue == "" {
					currentValue = currentQuestion.defaultAnswer
				}

				answer := selection(*currentQuestion, currentValue)
				fmt.Println("")

				if currentQuestion.key != "" && answer.value != "" {
					responses = append(responses, envData{key: currentQuestion.key, value: answer.value})
				}
				if currentQuestion.nextQuestion == nil {
					break
				}

				currentQuestion = currentQuestion.nextQuestion
			}
		}
	}

	envStr := mergeEnv(envContent, responses)
	saveEnvContent(envStr)
}

func input(defaultValue string) string {
	if defaultValue != "" {
		fmt.Printf("(current: %s, press enter to keep it): ", defaultValue)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	val := strings.TrimSpace(input)
	if val == "" {
		return defaultValue
	}

	return val
}

func selection(q question, currentValue string) *answer {
	if len(q.answers) > 0 {
		fmt.Println(q.prompt)
		fmt.Print("Please choose: ")
	} else {
		fmt.Print(q.prompt + ": ")
	}

	for {
		response := ""
		if len(q.answers) == 0 {
			response = input(currentValue)
		} else {
			resolvedAnswer := resolveAnswer(q.answers, currentValue)
			response = input(resolvedAnswer)
		}

		if response == "" && q.mandatory {
			if len(q.answers) == 0 {
				fmt.Println("Please provide a value")
			} else {
				fmt.Println("Please select an option")
			}
			continue
		}

		if len(q.answers) > 0 {
			if selected, ok := q.answers[response]; ok {
				if selected.nextQuestion == nil && q.nextQuestion != nil {
					return &answer{value: selected.value, nextQuestion: q.nextQuestion}
				}
				return &selected
			}

			fmt.Println("invalid selection")
			continue
		}

		return &answer{value: response, nextQuestion: q.nextQuestion}
	}
}

func mergeEnv(currentEnv string, data []envData) string {
	currentLines := strings.Split(currentEnv, "\n")

	for _, envLine := range data {
		envRow := envLine.key + "=" + envLine.value
		if keyId, ok := lookup(currentLines, envLine.key); ok {
			currentLines[keyId] = envRow
			continue
		}

		currentLines = append(currentLines, envRow)
	}

	return strings.Join(currentLines, "\n")
}

func lookup(lines []string, keyValue string) (int, bool) {
	for i, line := range lines {
		lineParts := strings.Split(line, "=")
		key := strings.TrimSpace(lineParts[0])
		if key == keyValue {
			return i, true
		}
	}

	return 0, false
}

func lookupValue(envContent, keyValue string) string {
	lines := strings.Split(envContent, "\n")
	for _, line := range lines {
		lineParts := strings.Split(line, "=")
		key := strings.TrimSpace(lineParts[0])
		if key == keyValue {
			if len(lineParts) > 1 {
				return strings.TrimSpace(strings.Join(lineParts[1:], "="))
			}
		}
	}

	return ""
}

func getEnvContent() string {
	content, err := os.ReadFile(envFileName)
	if err != nil {
		return ""
	}

	return string(content)
}

func saveEnvContent(fileContent string) error {
	return os.WriteFile(envFileName, []byte(fileContent), 0644)
}

func resolveAnswer(a answers, value string) string {
	for key, answer := range a {
		if answer.value == value {
			return key
		}
	}

	return ""
}

func getStorages(data []envData) []string {
	re := regexp.MustCompile(`.*_STORAGE`)

	storages := make([]string, 0)
	for _, env := range data {
		if re.MatchString(env.key) && !sliceContains(storages, env.value) {
			storages = append(storages, env.value)
		}
	}

	return storages
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
