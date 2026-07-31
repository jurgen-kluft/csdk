package csdk

import (
	"strings"

	cespressif "github.com/jurgen-kluft/csdk/espressif"
	corepkg "github.com/jurgen-kluft/gcore"
	denv "github.com/jurgen-kluft/gide/denv"
)

func getVarsArduino(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, hardwareId string, vars *corepkg.Vars) {
	if tc, err := cespressif.ParseToolchain(buildTarget.Arch().String()); err == nil {
		cespressif.GetVars(tc, hardwareId, vars)

		// Override some specific settings for Arduino based on build configuration

		// 0: None (ARDUHAL_LOG_LEVEL_NONE)
		// 1: Error (ARDUHAL_LOG_LEVEL_ERROR)
		// 2: Warning (ARDUHAL_LOG_LEVEL_WARN)
		// 3: Info (ARDUHAL_LOG_LEVEL_INFO)
		// 4: Debug (ARDUHAL_LOG_LEVEL_DEBUG)
		// 5: Verbose (ARDUHAL_LOG_LEVEL_VERBOSE)

		defines := make([]string, 0, 8)
		if buildConfig.IsDebug() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.debug}")
			defines = append(defines, "-DTARGET_DEBUG")
			vars.Set("build.code_debug", "4")
		} else if buildConfig.IsRelease() {
			vars.Set("compiler.optimization_flags", "{compiler.optimization_flags.release}")
			defines = append(defines, "-DTARGET_RELEASE")
			vars.Set("build.code_debug", "1")
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

			if strings.Contains(strings.ToLower(hardwareId), "esp32s3") {
				defines = append(defines, "-DBOARD_HAS_PSRAM")
				vars.Set("build.psram_type", "opi")
			}

		} else if buildTarget.Esp8266() {
			defines = append(defines, "-DTARGET_ESP8266")
			vars.Prepend("compiler.cpreprocessor.flags", "{build.defines}")
		}

		vars.Set("upload.speed", "115200")

		// Convert mcu string to be able to be marked as a valid C/C++ define
		mcuDefine := strings.ToUpper(strings.ReplaceAll(vars.GetFirstOrEmpty("build.mcu"), "-", "_"))
		defines = append(defines, "-DTARGET_"+mcuDefine)

		// Convert board name string to be able to be marked as a valid C/C++ define
		boardNameDefine := strings.ToUpper(strings.ReplaceAll(vars.GetFirstOrEmpty("board.name"), "-", "_"))
		defines = append(defines, "-DTARGET_"+boardNameDefine)

		vars.Append("build.defines", defines...)
	}
}
