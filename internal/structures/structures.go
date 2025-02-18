package structures

import "time"

type UserStruct struct {
	DisabledAt         any       `json:"disabled_at"`
	Email              string    `json:"email"`
	ID                 string    `json:"id"`
	InsertedAt         time.Time `json:"inserted_at"`
	LastSignedInAt     any       `json:"last_signed_in_at"`
	LastSignedInMethod any       `json:"last_signed_in_method"`
	Role               string    `json:"role"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type GetUsersStruct struct {
	Users []UserStruct `json:"data"`
}

type CreateUserBody struct {
	User struct {
		Email                string `json:"email"`
		Role                 string `json:"role"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	} `json:"user"`
}

type SoloUserStruct struct {
	User UserStruct `json:"data"`
}

type DelUserStruct struct {
	Error string `json:"error"`
}
type DeviceStruct struct {
	AllowedIps                    []string    `json:"allowed_ips"`
	Description                   interface{} `json:"description"`
	DNS                           []string    `json:"dns"`
	Endpoint                      string      `json:"endpoint"`
	ID                            string      `json:"id"`
	InsertedAt                    time.Time   `json:"inserted_at"`
	Ipv4                          string      `json:"ipv4"`
	Ipv6                          string      `json:"ipv6"`
	LatestHandshake               interface{} `json:"latest_handshake"`
	Mtu                           int         `json:"mtu"`
	Name                          string      `json:"name"`
	PersistentKeepalive           int         `json:"persistent_keepalive"`
	PresharedKey                  string      `json:"preshared_key"`
	PublicKey                     string      `json:"public_key"`
	RemoteIP                      interface{} `json:"remote_ip"`
	RxBytes                       interface{} `json:"rx_bytes"`
	ServerPublicKey               string      `json:"server_public_key"`
	TxBytes                       interface{} `json:"tx_bytes"`
	UpdatedAt                     time.Time   `json:"updated_at"`
	UseDefaultAllowedIps          bool        `json:"use_default_allowed_ips"`
	UseDefaultDNS                 bool        `json:"use_default_dns"`
	UseDefaultEndpoint            bool        `json:"use_default_endpoint"`
	UseDefaultMtu                 bool        `json:"use_default_mtu"`
	UseDefaultPersistentKeepalive bool        `json:"use_default_persistent_keepalive"`
	UserID                        string      `json:"user_id"`
}

type ListDevicesCmd struct {
	Devices []DeviceStruct `json:"data"`
}

type SoloDeviceStruct struct {
	Device DeviceStruct `json:"data"`
}
