package numbers

type MessageNumber uint8

const (
	SSH_MSG_DISCONNECT                MessageNumber = 1
	SSH_MSG_IGNORE                    MessageNumber = 2
	SSH_MSG_UNIMPLEMENTED             MessageNumber = 3
	SSH_MSG_DEBUG                     MessageNumber = 4
	SSH_MSG_SERVICE_REQUEST           MessageNumber = 5
	SSH_MSG_SERVICE_ACCEPT            MessageNumber = 6
	SSH_MSG_KEXINIT                   MessageNumber = 20
	SSH_MSG_NEWKEYS                   MessageNumber = 21
	SSH_MSG_KEXDH_INIT                MessageNumber = 30
	SSH_MSG_KEXDH_REPLY               MessageNumber = 31
	SSH_MSG_KEX_DH_GEX_REQUEST_OLD    MessageNumber = 30
	SSH_MSG_KEX_DH_GEX_GROUP          MessageNumber = 31
	SSH_MSG_KEX_DH_GEX_INIT           MessageNumber = 32
	SSH_MSG_KEX_DH_GEX_REPLY          MessageNumber = 33
	SSH_MSG_KEX_DH_GEX_REQUEST        MessageNumber = 34
	SSH_MSG_USERAUTH_REQUEST          MessageNumber = 50
	SSH_MSG_USERAUTH_FAILURE          MessageNumber = 51
	SSH_MSG_USERAUTH_SUCCESS          MessageNumber = 52
	SSH_MSG_USERAUTH_BANNER           MessageNumber = 53
	SSH_MSG_USERAUTH_INFO_REQUEST     MessageNumber = 60
	SSH_MSG_USERAUTH_INFO_RESPONSE    MessageNumber = 61
	SSH_MSG_GLOBAL_REQUEST            MessageNumber = 80
	SSH_MSG_REQUEST_SUCCESS           MessageNumber = 81
	SSH_MSG_REQUEST_FAILURE           MessageNumber = 82
	SSH_MSG_CHANNEL_OPEN              MessageNumber = 90
	SSH_MSG_CHANNEL_OPEN_CONFIRMATION MessageNumber = 91
	SSH_MSG_CHANNEL_OPEN_FAILURE      MessageNumber = 92
	SSH_MSG_CHANNEL_WINDOW_ADJUST     MessageNumber = 93
	SSH_MSG_CHANNEL_DATA              MessageNumber = 94
	SSH_MSG_CHANNEL_EXTENDED_DATA     MessageNumber = 95
	SSH_MSG_CHANNEL_EOF               MessageNumber = 96
	SSH_MSG_CHANNEL_CLOSE             MessageNumber = 97
	SSH_MSG_CHANNEL_REQUEST           MessageNumber = 98
	SSH_MSG_CHANNEL_SUCCESS           MessageNumber = 99
	SSH_MSG_CHANNEL_FAILURE           MessageNumber = 100
	SSH_MSG_USERAUTH_PASSWD_CHANGEREQ MessageNumber = 60
	SSH_MSG_USERAUTH_PK_OK            MessageNumber = 60
)
