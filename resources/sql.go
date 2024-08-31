package resources

type SqlResource interface {
	ConnectionString() string
	SqlDialect() string
	MigrationFolder() string
}

type MigrationsFolder string

func (m MigrationsFolder) MigrationFolder() string {
	return string(m)
}
