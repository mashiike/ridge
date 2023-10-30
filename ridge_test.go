package ridge_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fujiwara/ridge"
)

func TestNewLambdaHandler(t *testing.T) {
	handler := ridge.NewLambdaHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}),
	)
	if handler == nil {
		t.Error("handler is nil")
	}
	bs, err := os.ReadFile("test/get.json")
	if err != nil {
		t.Fatalf("failed to open test/get.json: %s", err)
	}
	resp, err := handler(json.RawMessage(bs))
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if _, ok := resp.(ridge.Response); !ok {
		t.Errorf("resp is not ridge.Response: %#v", resp)
	}
}

func TestNewLambdaHandler__InvokeModeStreaming(t *testing.T) {
	t.Setenv("RIDGE_INVOKE_MODE", "streaming")
	handler := ridge.NewLambdaHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}),
	)
	if handler == nil {
		t.Error("handler is nil")
	}
	bs, err := os.ReadFile("test/get.json")
	if err != nil {
		t.Fatalf("failed to open test/get.json: %s", err)
	}
	resp, err := handler(json.RawMessage(bs))
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if _, ok := resp.(*events.LambdaFunctionURLStreamingResponse); !ok {
		t.Errorf("resp is not *events.LambdaFunctionURLStreamingResponse: %#v", resp)
	}
}
