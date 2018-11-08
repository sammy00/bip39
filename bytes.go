package bip39

func paddedPrepend(dst, src []byte, size int) []byte {
	for i := size - len(src); i > 0; i-- {
		dst = append(dst, 0)
	}

	return append(dst, src...)
}
