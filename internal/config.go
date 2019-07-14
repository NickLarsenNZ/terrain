package internal

type Config struct {
	Diff           bool
	DiffDirectoryA string
	DiffDirectoryB string
	ScanDirectory  string
	ShowResources  bool
	OutputFormat   string
}

func (c Config) Validate() error {
	return nil
}
