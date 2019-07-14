package internal

import "fmt"

type ModuleConfiguration struct {
	Resources   []interface{}
	DataSources []interface{}
	Variables   []interface{}
	Outputs     []interface{}
	Providers   []interface{}
	Modules     []interface{}
}

func Scan(directory string) *ModuleConfiguration {
	return &ModuleConfiguration{}
}

// Haven't quite thought this out, but it would be nice to be able to diff:
// Two different directories configurations
// Two configurations based on git commit hashes/branches/tags
// Between two ModuleConfiguration's (eg: from JSON output)

func DiffModuleConfiguration(m1 *ModuleConfiguration, m2 *ModuleConfiguration) interface{} {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func DiffDirectories(dir1 string, dir2 string) interface{} {
	return DiffModuleConfiguration(Scan(dir1), Scan(dir2))
}

func DiffGitCommits(commit1 string, commit2 string) interface{} {
	// See: https://stackoverflow.com/questions/15250343/checkout-old-commit-to-temporary-directory-with-git

	// Checkout each commit
	tmpDir1 := "/tmp/directory/of/${commit1}"
	tmpDir2 := "/tmp/directory/of/${commit2}"

	diff := DiffDirectories(tmpDir1, tmpDir2)

	// Delete temp dirs

	return diff
}
