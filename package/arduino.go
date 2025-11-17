package csdk

import (
	"strings"

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
			vars.Set("build.code_debug", "1")
		} else if buildConfig.IsRelease() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.release}")
			defines = append(defines, "-DTARGET_RELEASE")
			vars.Set("build.code_debug", "0")
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
			vars.Prepend("compiler.cpreprocessor.flags", "{build.defines}")
		}

		defines = append(defines, "-DBOARD_HAS_PSRAM")
		vars.Set("build.psram_type", "opi")

		vars.Set("build.partitions", "huge_app")
		vars.Set("build.flash.mode", "dio")
		vars.Set("build.boot", "dio")
		vars.Set("build.boot_freq", "80m")
		vars.Set("build.flash_freq", "80m")

		// Convert mcu string to be able to be marked as a valid C/C++ define
		mcuDefine := strings.ToUpper(strings.ReplaceAll(vars.GetFirstOrEmpty("build.mcu"), "-", "_"))
		defines = append(defines, "-DTARGET_"+mcuDefine)

		// Convert board name string to be able to be marked as a valid C/C++ define
		boardNameDefine := strings.ToUpper(strings.ReplaceAll(vars.GetFirstOrEmpty("board.name"), "-", "_"))
		defines = append(defines, "-DTARGET_"+boardNameDefine)

		vars.Append("build.defines", defines...)
	}
}
