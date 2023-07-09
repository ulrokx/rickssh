package algorithms

type EncryptionAlgorithm string

const (
	THREEDES_CBC      EncryptionAlgorithm = "3des-cbc"
	BLOWFISH_CBC      EncryptionAlgorithm = "blowfish-cbc"
	TWOFISH_CBC       EncryptionAlgorithm = "twofish-cbc"
	AES128_CBC        EncryptionAlgorithm = "aes128-cbc"
	AES192_CBC        EncryptionAlgorithm = "aes192-cbc"
	AES256_CBC        EncryptionAlgorithm = "aes256-cbc"
	AES128_CTR        EncryptionAlgorithm = "aes128-ctr"
	AES192_CTR        EncryptionAlgorithm = "aes192-ctr"
	AES256_CTR        EncryptionAlgorithm = "aes256-ctr"
	AES128_GCM        EncryptionAlgorithm = "aes128-gcm@openssh.com"
	AES256_GCM        EncryptionAlgorithm = "aes256-gcm@openssh.com"
	CHACHA20_POLY1305 EncryptionAlgorithm = "chacha20-poly1305@openssh.com"
)

var EncryptionAlgorithmFromString = map[string]EncryptionAlgorithm{
	"3des-cbc":                      THREEDES_CBC,
	"blowfish-cbc":                  BLOWFISH_CBC,
	"twofish-cbc":                   TWOFISH_CBC,
	"aes128-cbc":                    AES128_CBC,
	"aes192-cbc":                    AES192_CBC,
	"aes256-cbc":                    AES256_CBC,
	"aes128-ctr":                    AES128_CTR,
	"aes192-ctr":                    AES192_CTR,
	"aes256-ctr":                    AES256_CTR,
	"aes128-gcm@openssh.com":        AES128_GCM,
	"aes256-gcm@openssh.com":        AES256_GCM,
	"chacha20-poly1305@openssh.com": CHACHA20_POLY1305,
}
