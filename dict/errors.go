package dict

import "github.com/pkg/errors"

// errors enumerations
var (
	// ErrDisabledTrie signals trie for a given language has been disabled
	ErrDisabledTrie = errors.New("trie is disabled")
	// ErrOccupiedLanguage has been registered somewhere else
	ErrOccupiedLanguage = errors.New("language already occupied")
	// ErrUnknownLanguage signals a language as a parameter isn't registered
	ErrUnknownLanguage = errors.New("the provided language isn't registered")
)
