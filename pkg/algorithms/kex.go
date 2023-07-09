package algorithms

type KEXAlgorithm string

const (
	DIFFIE_HELLMAN_GROUP1_SHA1           KEXAlgorithm = "diffie-hellman-group1-sha1"
	DIFFIE_HELLMAN_GROUP14_SHA1          KEXAlgorithm = "diffie-hellman-group14-sha1"
	DIFFIE_HELLMAN_GROUP14_SHA256        KEXAlgorithm = "diffie-hellman-group14-sha256"
	DIFFIE_HELLMAN_GROUP16_SHA512        KEXAlgorithm = "diffie-hellman-group16-sha512"
	DIFFIE_HELLMAN_GROUP18_SHA512        KEXAlgorithm = "diffie-hellman-group18-sha512"
	DIFFIE_HELLMAN_GROUP_EXCHANGE_SHA1   KEXAlgorithm = "diffie-hellman-group-exchange-sha1"
	DIFFIE_HELLMAN_GROUP_EXCHANGE_SHA256 KEXAlgorithm = "diffie-hellman-group-exchange-sha256"
	ECDH_SHA2_NISTP256                   KEXAlgorithm = "ecdh-sha2-nistp256"
	ECDH_SHA2_NISTP384                   KEXAlgorithm = "ecdh-sha2-nistp384"
	ECDH_SHA2_NISTP521                   KEXAlgorithm = "ecdh-sha2-nistp521"
	CURVE25519_SHA256                    KEXAlgorithm = "curve25519-sha256"
	CURVE25519_SHA256_LIBSSH_ORG         KEXAlgorithm = "curve25519-sha256@libssh.org"
	SNTRUP761X25519_SHA512               KEXAlgorithm = "sntrup761x25519-sha512@openssh.org"
)

var KexAlgorithmFromString = map[string]KEXAlgorithm{
	"diffie-hellman-group1-sha1":           DIFFIE_HELLMAN_GROUP1_SHA1,
	"diffie-hellman-group14-sha1":          DIFFIE_HELLMAN_GROUP14_SHA1,
	"diffie-hellman-group14-sha256":        DIFFIE_HELLMAN_GROUP14_SHA256,
	"diffie-hellman-group16-sha512":        DIFFIE_HELLMAN_GROUP16_SHA512,
	"diffie-hellman-group18-sha512":        DIFFIE_HELLMAN_GROUP18_SHA512,
	"diffie-hellman-group-exchange-sha1":   DIFFIE_HELLMAN_GROUP_EXCHANGE_SHA1,
	"diffie-hellman-group-exchange-sha256": DIFFIE_HELLMAN_GROUP_EXCHANGE_SHA256,
	"ecdh-sha2-nistp256":                   ECDH_SHA2_NISTP256,
	"ecdh-sha2-nistp384":                   ECDH_SHA2_NISTP384,
	"ecdh-sha2-nistp521":                   ECDH_SHA2_NISTP521,
	"curve25519-sha256":                    CURVE25519_SHA256,
	"curve25519-sha@libssh.org":            CURVE25519_SHA256_LIBSSH_ORG,
	"sntrup761x25519-sha512@openssh.org":   SNTRUP761X25519_SHA512,
}
