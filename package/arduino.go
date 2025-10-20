package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
	cespressif "github.com/jurgen-kluft/ccode/espressif"
)

func getVarsArduino(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, hardwareId string, vars *corepkg.Vars) {
	if tc, err := cespressif.ParseToolchain(buildTarget.Arch().String()); err == nil {
		cespressif.GetVars(tc, hardwareId, vars)

		// Override some specific settings for Arduino based on build configuration
		defines := make([]string, 0, 8)
		if buildConfig.IsDebug() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.debug}")
			defines = append(defines, "-DTARGET_DEBUG")
		} else if buildConfig.IsRelease() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.release}")
			defines = append(defines, "-DTARGET_RELEASE")
		}

		if buildConfig.IsFinal() {
			defines = append(defines, "-DTARGET_FINAL")
		}

		if buildConfig.IsTest() {
			defines = append(defines, "-DTARGET_TEST")
		}

		defines = append(defines, "-DTARGET_ARDUINO")
		if buildTarget.Esp32() {
			defines = append(defines, "-DTARGET_ESP32")
		} else if buildTarget.Esp8266() {
			defines = append(defines, "-DTARGET_ESP8266")
		}

		vars.Prepend("compiler.cpreprocessor.flags", "{build.defines}")
		vars.Append("build.defines", defines...)
	}
}
