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
			vars.Append("build.defines", "-DTARGET_DEBUG")
		} else if buildConfig.IsRelease() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.release}")
			vars.Append("build.defines", "-DTARGET_RELEASE")
		}

		if buildConfig.IsFinal() {
			vars.Append("build.defines", "-DTARGET_FINAL")
		}

		if buildConfig.IsTest() {
			vars.Append("build.defines", "-DTARGET_TEST")
		}

		vars.Append("build.defines", "-DTARGET_ARDUINO")
		if buildTarget.Esp32() {
			vars.Append("build.defines", "-DTARGET_ESP32")
		} else if buildTarget.Esp8266() {
			vars.Append("build.defines", "-DTARGET_ESP8266")
		}

	}
}
