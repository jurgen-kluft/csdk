package csdk

import (
	denv "github.com/jurgen-kluft/ccode/denv"
)

const (
	repo_path = "github.com\\jurgen-kluft\\"
	repo_name = "csdk"
)

// var UserVariables = map[string]string{
// 	"build.path":         "build",     // Output directory for build files
// 	"build.project_name": "MyProject", // Name of the project (used for naming output files)
// 	"build.arch":         "x64",       // Architecture: x86, x64, arm64
// 	"build.config":       "debug-dev", // Build configuration: debug-dev, release-dev, debug-dev-test, ..
// 	"build.defines":      "",          // Additional preprocessor defines
// 	"include_paths":      "",          // Additional include paths for the compiler
// }

func GetPackage() *denv.Package {
	name := repo_name

	// main package
	mainpkg := denv.NewPackage(repo_path, repo_name)
	mainpkg.SetGetVarsFunc(getVars)

	// main library (header files only)
	mainlib := denv.SetupCppHeaderProject(mainpkg, name)

	mainpkg.AddMainLib(mainlib)
	return mainpkg
}
