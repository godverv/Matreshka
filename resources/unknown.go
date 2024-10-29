package resources

type Unknown struct {
	Name `yaml:"resource_name" env:"-"`

	Content map[string]string
}

func (u *Unknown) GetType() string {
	return "Unknown"
}
