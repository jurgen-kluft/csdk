package csdk

import (
	denv "github.com/jurgen-kluft/ccode/denv"
)

const (
	repo_path = "github.com\\jurgen-kluft\\"
	repo_name = "csdk"
)

func GetPackage() *denv.Package {
	main_pkg := denv.NewPackage(repo_path, repo_name)
	main_pkg.SetGetVarsFunc(getVars)

	mainlib := denv.SetupCppLibProject(main_pkg, repo_name)

	main_pkg.AddMainLib(mainlib)
	return main_pkg
}
