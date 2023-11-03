package resources

type Unknown struct {
	AppResource

	Content map[string]string
}

func (u *Unknown) GetType() string {
	return "Unknown"
}

func (u *Unknown) ToEnv() map[string]string {
	return u.Content
}

func (u *Unknown) FromEnv(in map[string]string) (err error) {
	u.Content = in
	return nil
}
