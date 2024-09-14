package main

var appUrlQuestion = question{
	key:           "APP_URL",
	prompt:        "Please provide app URL",
	defaultAnswer: "http://localhost:8080",
	nextQuestion:  &appPortQuestion,
}

var appPortQuestion = question{
	key:           "HTTP_LISTENING_PORT",
	prompt:        "Please provide app URL",
	defaultAnswer: "http://localhost:8080",
	nextQuestion:  &sessionStorageQuestions,
}
