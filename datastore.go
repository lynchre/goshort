package main

import "errors"

type URLEntry struct {
	Shortcode string
	URL       string
	Hits      int
}

type DataStore interface {
	AddURL(entry URLEntry) error
	GetURL(shortcode string) (URLEntry, error)
	ListURLs() ([]URLEntry, error)
	HitURL(shortcode string) (URLEntry, error)
}

type MapStore struct {
	urls map[string]URLEntry
}

func NewMapStore() *MapStore {
	store := &MapStore{
		urls: make(map[string]URLEntry),
	}
	return store
}

func (store *MapStore) AddURL(entry URLEntry) error {
	short := entry.Shortcode

	// check for duplicate entry
	if _, ok := store.urls[short]; !ok {
		return errors.New("Duplicate shortcode entry")
	}

	store.urls[short] = entry
	return nil
}

func (store *MapStore) GetURL(shortcode string) (URLEntry, error) {
	entry, ok := store.urls[shortcode]
	if !ok {
		return entry, errors.New("Shortcode does not exist")
	}
	return entry, nil
}

func (store *MapStore) ListURLs() ([]URLEntry, error) {
	// TODO 1: return a list of all URLEntrys in the MapStore
	return []URLEntry{}, errors.New("not implemented")
}

func (store *MapStore) HitURL(shortcode string) (URLEntry, error) {
	// TODO 2: Get the URLEntry from the shortcode provided, increment its Hits
	// property, and then return the URLEntry
	return URLEntry{}, errors.New("not implemented")
}
