package resources

type Unknown struct {
	Name `yaml:"resource_name"`

	Content map[string]string
}

func (u *Unknown) GetType() string {
	return "Unknown"
}
