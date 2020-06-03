package hazelcaststore

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/config"
)

/**
 * Use https://hazelcast.org/ as an implement of this:
 *
 * type Store interface {
 *	Get(r *http.Request, name string) (*Session, error)
 *	New(r *http.Request, name string) (*Session, error)
 *	Save(r *http.Request, w http.ResponseWriter, s *Session) error
 * }
 */

// HazelcastStore is a remote session store, spread across multiple remote server
// processes, connecting using a multiplexing client.
type HazelcastStore struct {
	client hazelcast.Client
}

// -- Implementation of the Store interface from Gorilla sessions

// New TODO
func (hazelcastStore *HazelcastStore) New(r *http.Request, name string) (*sessions.Session, error) {
	log.Printf("START New(%s)\n", name)

	session := sessions.NewSession(hazelcastStore, name)
	//session.IsNew = true

	log.Printf("END   New(%s)\n", name)
	return session, nil
}

// Get TODO
func (hazelcastStore *HazelcastStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	//TODO
	return nil, nil
}

// Save TODO
func (hazelcastStore *HazelcastStore) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	//TODO
	return nil
}

// -- HazelcastStore : Connection options

// NewHazelcastStore builds a store using default client configuration
func NewHazelcastStore() (hazelcastStore *HazelcastStore, err error) {
	clientConfig := hazelcast.NewConfig()
	return NewHazelcastStoreFromConfig(clientConfig)
}

// NewHazelcastStoreFromConfig builds a store using the supplied client configuration
func NewHazelcastStoreFromConfig(clientConfig *config.Config) (hazelcastStore *HazelcastStore, err error) {
	hazelcastClient, err := hazelcast.NewClientWithConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return NewHazelcastStoreFromClient(hazelcastClient), nil
}

// NewHazelcastStoreFromClient builds a store using the supplied client connection
func NewHazelcastStoreFromClient(hazelcastClient hazelcast.Client) *HazelcastStore {
	hazelcastStore := &HazelcastStore{
		client: hazelcastClient,
	}
	return hazelcastStore
}

// --
