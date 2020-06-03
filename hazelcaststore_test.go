package hazelcaststore

import (
	"net/http"
	"testing"

	"github.com/hazelcast/hazelcast-go-client"
)

const exampleURL string = "http://www.example.com"

var hazelcastClient hazelcast.Client = nil

type TestNewCase struct {
	name  string
	input string
}

func Disconnect() {
	if hazelcastClient != nil {
		if hazelcastClient.LifecycleService().IsRunning() {
			hazelcastClient.Shutdown()
		}
	}
}

func New(t *testing.T, name string, hazelcastStore HazelcastStore) {
	request, err := http.NewRequest(http.MethodGet, exampleURL, nil)
	if err != nil {
		t.Errorf("New('%s'): %s for %s => %s", name, http.MethodGet, exampleURL, err)
		return
	}

	session, err := hazelcastStore.New(request, name)
	if err != nil {
		t.Errorf("New('%s'): hazelcastStore.New => %s", name, err)
		return
	}
	if session == nil {
		t.Errorf("New('%s'): hazelcastStore.New => nil session", name)
		return
	}
	if session.IsNew != true {
		t.Errorf("New('%s'): hazelcastStore.New => session '%v' not new", name, session)
		return
	}

	return
}

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("Test: Skip as testing.Short()")
	}

	// Connect to external processes, skip tests if failed
	hazelcastClient, err := hazelcast.NewClient()
	if err != nil {
		t.Errorf("Test: No connection => %s", err)
	}

	hazelcastStore := NewHazelcastStoreFromClient(hazelcastClient)

	testsNew := []TestNewCase{
		{name: "empty input", input: ""},
		{name: "valid input", input: "xyz"},
	}

	for _, testNew := range testsNew {
		t.Run(testNew.name, func(t *testing.T) {
			New(t, testNew.input, *hazelcastStore)
		})
	}

	t.Cleanup(Disconnect)
}
