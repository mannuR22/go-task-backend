package business

import hello "github.com/mannuR22/go-task-backend.git/models"

func HelloMessage() hello.Hello {
	var hello hello.Hello
	hello.Message = "Hello World!"
	return hello
}
