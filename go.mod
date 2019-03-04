module github.com/sammyne/bip39

require (
	github.com/derekparker/trie v0.0.0-20180212171413-e608c2733dc7
	github.com/pkg/errors v0.8.0
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/text v0.3.0
)

replace golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd => github.com/golang/crypto v0.0.0-20181106171534-e4dc69e5b2fd

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
