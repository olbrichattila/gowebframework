package session

import (
	"encoding/json"
	"framework/internal/app/storage"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	sessionName        = "go-session"
	sessionStoragePath = "./sessions/"
)

func New(storage storage.Storager) Sessioner {
	return &Session{
		storage: storage,
	}
}

type Sessioner interface {
	Init(w http.ResponseWriter, r *http.Request)
	Set(string, string)
	Delete(string)
	Get(string) string
	Redirect(string)
	Close()
	RemoveSession()
}

type Session struct {
	storage storage.Storager
	id      string
	r       *http.Request
	w       http.ResponseWriter
}

func (s *Session) Init(w http.ResponseWriter, r *http.Request) {
	s.r = r
	s.w = w
	cookie, err := r.Cookie(sessionName)
	if err != nil {
		s.id = uuid.New().String()
		cookie := &http.Cookie{
			Name:    sessionName,
			Value:   s.id,
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		}

		http.SetCookie(w, cookie)
		return
	}

	s.id = cookie.Value
}

func (s *Session) Set(key string, value string) {
	fileName := s.getSessionFileName()
	values := make(map[string]string, 0)

	hasKey, err := s.storage.HasKey(fileName)
	if err != nil {
		// TODO may log
		return
	}

	if hasKey {
		content, err := s.storage.Get(fileName)
		if err != nil {
			// TODO may log
			return
		}

		err = json.Unmarshal([]byte(content), &values)
		if err != nil {
			// TODO may log
			return
		}
	}

	values[key] = value
	newContent, err := json.Marshal(values)
	if err != nil {
		// TODO may log
		return
	}

	err = s.storage.Put(fileName, string(newContent))
	if err != nil {
		// TODO may log
	}
}

func (s *Session) Delete(key string) {
	fileName := s.getSessionFileName()
	values := make(map[string]string, 0)

	hasKey, err := s.storage.HasKey(fileName)
	if err != nil {
		// TODO may log
		return
	}

	if hasKey {
		content, err := s.storage.Get(fileName)
		if err != nil {
			// TODO may log
			return
		}

		err = json.Unmarshal([]byte(content), &values)
		if err != nil {
			// TODO may log
			return
		}
	}

	delete(values, key)
	newContent, err := json.Marshal(values)
	if err != nil {
		log.Fatalf("Failed to marshall to map: %s", err)
		return
	}

	err = s.storage.Put(fileName, string(newContent))
	if err != nil {
		// TODO may log
	}
}

func (s *Session) Get(key string) string {
	if s.id == "" {
		return ""
	}

	values := make(map[string]string, 0)
	fileName := s.getSessionFileName()
	hasKey, err := s.storage.HasKey(fileName)
	if err != nil || !hasKey {
		return ""
	}

	content, err := s.storage.Get(fileName)
	if err != nil {
		return ""
	}

	err = json.Unmarshal([]byte(content), &values)
	if err != nil {
		return ""
	}

	if val, ok := values[key]; ok {
		return val
	}

	return ""
}

func (s *Session) Close() {
	fileName := s.getSessionFileName()

	s.storage.Delete(fileName)
}

func (s *Session) Redirect(path string) {
	http.Redirect(s.w, s.r, path, http.StatusSeeOther)
}

func (s *Session) getSessionFileName() string {
	return sessionStoragePath + s.id + ".json"
}

func (s *Session) RemoveSession() {
	cookie := &http.Cookie{
		Name:    sessionName,
		Value:   "",
		Expires: time.Now().Add(-24 * time.Hour),
		Path:    "/",
	}

	http.SetCookie(s.w, cookie)
}
