package requester

import (
	"strings"
	"testing"
)

func TestRequester_Get(t *testing.T) {
	r := New()
	_, err := r.Get("https://www.example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRequester_Post(t *testing.T) {
	r := New()
	body := strings.NewReader(`{"name": "John", "age": 30}`)
	_, err := r.Post("https://jsonplaceholder.typicode.com/posts", body)
	if err != nil {
		t.Fatalf("got an error while sending a post request. error: %s", err)
	}
}

func TestRequester_PostJSON(t *testing.T) {
	r := New()
	body := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	res, err := r.PostJSON("https://jsonplaceholder.typicode.com/posts", body)
	if err != nil {
		t.Fatalf("got an error while sending a post request. error: %s", err)
	}

	if !strings.Contains(string(res), `id": 101`) {
		t.Fatalf("Expected response to be id 101, got %s", string(res))
	}

}
