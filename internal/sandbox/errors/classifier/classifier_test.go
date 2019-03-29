package classifier_test

import (
	"errors"
	"testing"

	. "github.com/kamilsk/platform/internal/sandbox/errors/classifier"
)

func TestDefaultClassifier_Classify(t *testing.T) {
	defaultClassifier := DefaultClassifier{}

	if defaultClassifier.Classify(nil) != Succeed {
		t.Error("succeed is expected")
	}

	if defaultClassifier.Classify(errors.New("error")) != Retry {
		t.Error("retry is expected")
	}
}