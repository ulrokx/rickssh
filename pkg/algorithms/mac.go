package algorithms

type MACAlgorithm string

const (
	HMAC_SHA1         MACAlgorithm = "hmac-sha1"
	HMAC_SHA1_96      MACAlgorithm = "hmac-sha1-96"
	HMAC_SHA2_256     MACAlgorithm = "hmac-sha2-256"
	HMAC_SHA2_512     MACAlgorithm = "hmac-sha2-512"
	HMAC_MD5          MACAlgorithm = "hmac-md5"
	HMAC_MD5_96       MACAlgorithm = "hmac-md5-96"
	UMAC_64           MACAlgorithm = "umac-64@openssh.com"
	UMAC_128          MACAlgorithm = "umac-128@openssh.com"
	HMAC_SHA1_ETM     MACAlgorithm = "hmac-sha1-etm@openssh.com"
	HMAC_SHA1_96_ETM  MACAlgorithm = "hmac-sha1-96-etm@openssh.com"
	HMAC_SHA2_256_ETM MACAlgorithm = "hmac-sha2-256-etm@openssh.com"
	HMAC_SHA2_512_ETM MACAlgorithm = "hmac-sha2-512-etm@openssh.com"
	HMAC_MD5_ETM      MACAlgorithm = "hmac-md5-etm@openssh.com"
	HMAC_MD5_96_ETM   MACAlgorithm = "hmac-md5-96-etm@openssh.com"
	UMAC_64_ETM       MACAlgorithm = "umac-64-etm@openssh.com"
	UMAC_128_ETM      MACAlgorithm = "umac-128-etm@openssh.com"
)

var MacAlgorithmFromString = map[string]MACAlgorithm{
	"hmac-sha1":                     HMAC_SHA1,
	"hmac-sha1-96":                  HMAC_SHA1_96,
	"hmac-sha2-256":                 HMAC_SHA2_256,
	"hmac-sha2-512":                 HMAC_SHA2_512,
	"hmac-md5":                      HMAC_MD5,
	"hmac-md5-96":                   HMAC_MD5_96,
	"umac-64@openssh.com":           UMAC_64,
	"umac-128@openssh.com":          UMAC_128,
	"hmac-sha1-etm@openssh.com":     HMAC_SHA1_ETM,
	"hmac-sha1-96-etm@openssh.com":  HMAC_SHA1_96_ETM,
	"hmac-sha2-256-etm@openssh.com": HMAC_SHA2_256_ETM,
	"hmac-sha2-512-etm@openssh.com": HMAC_SHA2_512_ETM,
	"hmac-md5-etm@openssh.com":      HMAC_MD5_ETM,
	"hmac-md5-96-etm@openssh.com":   HMAC_MD5_96_ETM,
	"umac-64-etm@openssh.com":       UMAC_64_ETM,
	"umac-128-etm@openssh.com":      UMAC_128_ETM,
}
