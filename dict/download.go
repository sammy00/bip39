// +build download

package dict

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/chinese_simplified.txt -o golden/chinese_simplified.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/chinese_traditional.txt -o golden/chinese_traditional.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt -o golden/english.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/french.txt -o golden/french.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/italian.txt -o golden/italian.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/japanese.txt -o golden/japanese.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/korean.txt -o golden/korean.txt

//go:generate curl https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/spanish.txt -o golden/spanish.txt
