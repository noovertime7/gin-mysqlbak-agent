package clean

type Cleaner interface {
	CleanFile() error
}

type CleanerFactory struct {
	cleaner Cleaner
}

func NewCleanerFactory(cleaner Cleaner) *CleanerFactory {
	return &CleanerFactory{
		cleaner: cleaner,
	}
}

func (c *CleanerFactory) Run() error {
	return c.cleaner.CleanFile()
}
