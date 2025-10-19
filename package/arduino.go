package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
	cespressif "github.com/jurgen-kluft/csdk/package/espressif"
)

func getVarsArduino(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, hardwareId string, vars *corepkg.Vars) {
	if tc, err := cespressif.ParseToolchain(buildTarget.Arch().String()); err == nil {
		cespressif.GetVars(tc, hardwareId, vars)

		// Override some specific settings for Arduino based on build configuration
		if buildConfig.IsDebug() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.debug}")
		} else if buildConfig.IsRelease() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.release}")
		}
	}
}
