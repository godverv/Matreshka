package servers

const UnknownServerType = "unknown"

type Unknown struct {
	Name
	Values map[string]string
}

func (u *Unknown) GetType() string {
	return UnknownServerType
}

func (u *Unknown) GetPort() uint16 {
	return 0
}

func (u *Unknown) GetPortStr() string {
	return "0"
}

func (u *Unknown) ToEnv() map[string]string {
	u.Values[EnvServerName] = u.GetName()
	return u.Values
}

func (u *Unknown) FromEnv(in map[string]string) (err error) {
	u.Values = in
	return nil
}
