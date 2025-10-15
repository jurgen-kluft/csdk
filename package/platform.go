package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
)

// The user has to set add the following variables to the returned Vars object:
// 	"build.path":         "build",     // Output directory for build files
// 	"build.project_name": "MyProject", // Name of the project (used for naming output files)
// 	"build.arch":         "x64",       // Architecture: x86, x64, arm64
// 	"build.defines":      "",          // Additional preprocessor defines
// 	"include_paths":      "",          // Additional include paths for the compiler

func getVars(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, hardwareId string) (vars *corepkg.Vars) {
	vars = corepkg.NewVars(corepkg.VarsFormatCurlyBraces)
	if buildTarget.Windows() {
		getVarsWindows(buildTarget, buildConfig, vars)
	} else if buildTarget.Mac() {
		getVarsMac(buildTarget, buildConfig, vars)
	} else if buildTarget.Linux() {
		getVarsLinux(buildTarget, buildConfig, vars)
	} else if buildTarget.Arduino() {
		getVarsArduino(buildTarget, buildConfig, hardwareId, vars)
	}
	return
}
