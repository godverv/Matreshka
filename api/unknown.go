package api

type Unknown struct {
	Name
	Values map[string]string
}

func (u *Unknown) GetPort() uint16 {
	return 0
}

func (u *Unknown) GetPortStr() string {
	return "0"
}

func (u *Unknown) ToEnv() map[string]string {
	return u.Values
}

func (u *Unknown) FromEnv(in map[string]string) (err error) {
	u.Values = in
	return nil
}
