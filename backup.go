package sessiongo

import (
	"encoding/gob"
	"os"
	"time"
)

const backup = "./backup.gob"

type sessionBackup struct {
	SessionID  string
	Data       interface{}
	ExpireTime time.Time
	CreateAt   time.Time
}

type sessionsBackup struct {
	Sessions map[string]*sessionBackup
	Age      int
	Refresh  bool
}

//ReadBackup will read backup and cover session storage in memory
func (s *SessionsStorageInMemory) ReadBackup() error {
	file, err := os.Open(backup)
	if err != nil {
		return err
	}
	defer file.Close()

	sessions := &sessionsBackup{
		Sessions: make(map[string]*sessionBackup),
	}

	dec := gob.NewDecoder(file)
	err = dec.Decode(sessions)
	if err != nil {
		return err
	}

	s = sessions.transfer()

	return nil
}

func (s *SessionsStorageInMemory) BackUp() error {
	file, err := os.Create(backup)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(s.transfer())
	if err != nil {

		return err
	}

	return nil
}

func (s *SessionsStorageInMemory) transfer() *sessionsBackup {
	sessions := &sessionsBackup{
		Age:      s.age,
		Refresh:  s.refresh,
		Sessions: make(map[string]*sessionBackup),
	}

	for k, v := range s.sessions {
		sessbackup := &sessionBackup{
			SessionID:  v.sessionID,
			Data:       v.data,
			ExpireTime: v.expireTime,
			CreateAt:   v.createAt,
		}
		sessions.Sessions[k] = sessbackup
	}

	return sessions
}

func (s *sessionsBackup) transfer() *SessionsStorageInMemory {
	sessions := &SessionsStorageInMemory{
		age:      s.Age,
		refresh:  s.Refresh,
		sessions: make(map[string]*Session),
	}

	for k, v := range s.Sessions {
		session := &Session{
			sessionID:  v.SessionID,
			data:       v.Data,
			expireTime: v.ExpireTime,
			createAt:   v.CreateAt,
		}

		sessions.sessions[k] = session
	}

	return sessions
}
