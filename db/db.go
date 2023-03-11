package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jaredwarren/macroPi/macro"
	bolt "go.etcd.io/bbolt"
)

var (
	MacroBucket   = []byte("MacroBucket")
	ProfileBucket = []byte("ProfileBucket")
	ErrNotFound   = errors.New("not found")
)

type DBer interface {
	GetMacro(id string) (*macro.Macro, error)
	ListtMacros() ([]*macro.Macro, error)
	UpdateMacro(id string, m *macro.Macro) error
	DeleteMacro(id string) error

	Close() error
	// OldListSongs() ([]*model.Song, error) // Still need for migrate

	// // V2
	// GetSong(rfid string) (*model.Song, error)
	// ListSongs() ([]*model.Song, error)
	// UpdateSong(song *model.Song) error
	// DeleteSong(id string) error
	// SongExists(id string) (bool, error)

	// // RFID stuff
	// GetRFIDSong(rfid string) (*model.RFIDSong, error)
	// GetSongRFID(songID string) (*model.RFIDSong, error)
	// AddRFIDSong(rfid, songID string) error
	// RemoveRFIDSong(rfid, songID string) error
	// DeleteRFID(id string) error
	// ListRFIDSongs() ([]*model.RFIDSong, error)
	// RFIDExists(rfid string) (bool, error)
	// DeleteSongFromRFID(songID string) error
}

func NewMacroDB(path string) (DBer, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(MacroBucket)
		if err != nil {
			return fmt.Errorf("create bucke(%s)t: %s", MacroBucket, err)
		}

		_, err = tx.CreateBucketIfNotExists(ProfileBucket)
		if err != nil {
			return fmt.Errorf("create bucke(%s)t: %s", ProfileBucket, err)
		}

		// _, err = tx.CreateBucketIfNotExists([]byte(RFIDBucket))
		// if err != nil {
		// 	return fmt.Errorf("create bucke(%s)t: %s", RFIDBucket, err)
		// }
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &MacroDB{
		db: db,
	}, nil
}

type MacroDB struct {
	db *bolt.DB
}

func (s *MacroDB) Close() error {
	return s.db.Close()
}

func (s *MacroDB) GetMacro(id string) (*macro.Macro, error) {
	var m *macro.Macro
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(MacroBucket)
		v := b.Get([]byte(id))
		if v == nil {
			return ErrNotFound
		}
		err := json.Unmarshal(v, &m)
		if err != nil {
			return err
		}
		return nil
	})
	if m == nil {
		return nil, ErrNotFound
	}
	return m, err
}

func (s *MacroDB) ListtMacros() ([]*macro.Macro, error) {
	macros := []*macro.Macro{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(MacroBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var m *macro.Macro
			err := json.Unmarshal(v, &m)
			if err != nil {
				return err
			}
			macros = append(macros, m)
		}
		return nil
	})
	return macros, err
}

func (s *MacroDB) UpdateMacro(id string, m *macro.Macro) error {
	if id == "" {
		return fmt.Errorf("macro ID required")
	}
	m.ID = id
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MacroBucket)

		buf, err := json.Marshal(m)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), buf)
	})
}

func (s *MacroDB) DeleteMacro(id string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MacroBucket)
		return b.Delete([]byte(id))
	})
	return err
}
