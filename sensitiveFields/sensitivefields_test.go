package sensitiveFields

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"testing"
)

var loggerForTesting *logrus.Logger

func TestMain(m *testing.M) {
	loggerForTesting = &logrus.Logger{
		Formatter: &logrus.JSONFormatter{
			PrettyPrint: true,
		},
		Level: logrus.InfoLevel,
	}
}

type MyIOWriter struct {
	writer io.Writer
	testData func(data []byte) bool
}

func (w *MyIOWriter) Write(p []byte) (n int, err error) {
	if !w.testData(p) {
		return 0, errors.New("invalid struct being written")
	}

	return w.writer.Write(p)
}

// Testing whether tagging struct fields with '-' ignores them when marshaling.
func TestSensitiveFieldsUser(t *testing.T) {
	// Creating a user.
	user := &User{
		Username:     "johndoe",
		Password: "johndoepassword1234!@#$",
		Email: "johndoe@noreply.com",
		FirstName: "John",
		LastName: "Doe",
	}

	ioWriter := &MyIOWriter{
		writer: os.Stdout,
		testData: func(data []byte) bool {
			marshaledJSON, err := json.Marshal(user)
			if err != nil {
				return false
			} else {
				return reflect.DeepEqual(marshaledJSON, data)
			}
		},
	}

	loggerForTesting.Out = ioWriter
	loggerForTesting.WithFields(logrus.Fields{
		"user": user,
	}).Log(logrus.InfoLevel, "user information")
}
