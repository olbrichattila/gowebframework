package db

type MemoryConfig struct {
}

func newMemoryDBConfig() DBConfiger {
	return &MemoryConfig{}
}

func (c *MemoryConfig) getConnectionString() string {
	return ":memory:"
}

func (c *MemoryConfig) getConnectionName() string {
	return DriverNameSqLite
}
