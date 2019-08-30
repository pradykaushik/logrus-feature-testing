package main // import github.com/pradykaushik/logrus-feature-testing

import (
	"github.com/pradykaushik/logrus-feature-testing/sensitiveFields"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	user := &sensitiveFields.User{
		Username:     "johndoe",
		Password: "johndoepassword1234!@#$",
		Email: "johndoe@noreply.com",
		FirstName: "John",
		LastName: "Doe",
	}

	logger := &logrus.Logger{
		Formatter: &logrus.JSONFormatter{
			PrettyPrint: true,
		},
		Level: logrus.InfoLevel,
		Out: os.Stdout,
	}

	logger.WithFields(logrus.Fields{
		"user": user,
	}).Log(logrus.InfoLevel, "user information")
}
