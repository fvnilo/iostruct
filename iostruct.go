package iostruct

func identityMapper(d []byte) ([]byte, error) {
	return d, nil
}

func newDecryptionMapper(passphrase string) transform {
	return func(d []byte) ([]byte, error) {
		return decrypt(d, passphrase)
	}
}

func newEncryptionMapper(passphrase string) transform {
	return func(d []byte) ([]byte, error) {
		return encrypt(d, passphrase)
	}
}

func Write(filename string, data interface{}) error {
	return writeFile(filename, identityMapper, data)
}

func Read(filename string, e interface{}) error {
	return readFile(filename, identityMapper, e)
}

func WriteEncrypted(filename string, passphrase string, data interface{}) error {
	return writeFile(filename, newEncryptionMapper(passphrase), data)
}

func ReadEncrypted(filename string, passphrase string, o interface{}) error {
	return readFile(filename, newDecryptionMapper(passphrase), o)
}
