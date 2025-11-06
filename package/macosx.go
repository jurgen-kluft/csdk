package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
)

func getVarsMac(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, vars *corepkg.Vars) {
	// Using the buildTarget and buildConfig, we iterate over the platformVarsWindows map and
	// set the appropriate variables in the vars object. Some variables may depend on the buildConfig.
	for key, varList := range platformVarsMacOSX {
		for _, v := range varList {
			// If the variable has a Config field, we check if it matches the current buildConfig
			if len(v.Config) == 0 || denv.BuildConfigFromString(v.Config).Contains(buildConfig) {
				if v.Append {
					vars.Append(key, v.Value...)
				} else {
					vars.Set(key, v.Value...)
				}
			}
		}
	}
}

var platformVarsMacOSX = map[string][]denv.Var{
	// # Extensions
	//"build.obj.prefix": {{Value: []string{""}}},       // Object file prefix
	"build.obj.prefix": Value(""),       // Object file prefix
	"build.obj.suffix": Value(".o"),     // Object file suffix
	"build.dep.prefix": Value(""),       // Dependency file prefix
	"build.dep.suffix": Value(".d"),     // Dependency file suffix
	"build.dll.prefix": Value("lib"),    // Dynamic library file prefix
	"build.dll.suffix": Value(".dylib"), // Dynamic library file suffix
	"build.lib.prefix": Value("lib"),    // Static library file prefix
	"build.lib.suffix": Value(".a"),     // Static library file suffix
	"build.exe.prefix": Value(""),       // Executable file prefix (none for macOS)
	"build.exe.suffix": Value(""),       // Executable file suffix (none for macOS)

	// # Frameworks
	"compiler.frameworks.default": Values("-framework", "Foundation", "-framework", "CoreFoundation", "-framework", "IOKit", "-framework", "CoreServices"),
	"compiler.frameworks.metal":   Values("-framework", "Metal", "-framework", "MetalPerformanceShaders", "-framework", "MetalKit"),
	"compiler.frameworks.cocoa":   Values("-framework", "Cocoa"),
	"compiler.frameworks.appkit":  Values("-framework", "AppKit"),
	"compiler.frameworks.uikit":   Values("-framework", "UIKit"),

	// # Frameworks, current
	"compiler.frameworks": {{Value: []string{"{compiler.frameworks.default}"}}},

	// # Debug Info
	"compiler.debug_flags": {{Value: []string{"-g3"}}},

	// # Floating Point
	"compiler.floating_point_flags": {{Value: []string{"-ffp-model=precise"}}},

	// # Optimization flags
	"compiler.optimization_flags": {
		{Value: []string{"-O0"}, Config: "debug-*-*"},
		{Value: []string{"-O2"}, Config: "release-dev-*"},
		{Value: []string{"-O3"}, Config: "release-final-*"},
	},

	// # Compile Warning Levels
	"compiler.warning_flags": {{Value: []string{"-Wall", "-Wextra", "-Wno-unused-function", "-Wno-unused-parameter"}}},

	// # Compilers
	"compiler.c.cmd":    {{Value: []string{"clang"}}},
	"compiler.cpp.cmd":  {{Value: []string{"clang++"}}},
	"compiler.asm.cmd":  {{Value: []string{"clang"}}},
	"compiler.lib.cmd":  {{Value: []string{"ar"}}},
	"compiler.link.cmd": {{Value: []string{"clang++"}}},
	"compiler.size.cmd": {{Value: []string{"size"}}},

	// # Compiler/Archiver/Linker flags
	"compiler.cpreprocessor.flags": {{Value: []string{""}}},

	"compiler.c.flags": {
		{Value: []string{"-c", "-MMD", "{compiler.warning_flags}", "{compiler.common_werror_flags}", "{compiler.floating_point_flags}", "{build.c_standard}"}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
		{Value: []string{"-fexceptions"}, Config: "*-*-test", Append: true},
	},

	"compiler.cpp.flags": {
		{Value: []string{"-c", "-MMD", "{compiler.warning_flags}", "{compiler.common_werror_flags}", "{compiler.floating_point_flags}", "{build.cpp_standard}"}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
		{Value: []string{"-fexceptions"}, Config: "*-*-test", Append: true},
	},

	"compiler.asm.flags": {
		{Value: []string{"-c", "-x", "-MMD", "assembler-with-cpp", "{compiler.warning_flags}"}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
	},

	"compiler.lib.flags": {
		{Value: []string{"-rc"}},
		{Value: []string{}, Config: "debug-*-*", Append: true},
	},

	"compiler.link.flags": {
		{Value: []string{"{library.paths}", "{library.files}", "{compiler.frameworks}"}},
		{Value: []string{"-flto"}, Config: "*-final-*", Append: true},
	},

	// ## Compile c files
	"recipe.c.pattern": {{Value: []string{"{compiler.c.cmd}", "{compiler.c.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{build.includes}"}}},

	// ## Compile c++ files
	"recipe.cpp.pattern": {{Value: []string{"{compiler.cpp.cmd}", "{compiler.cpp.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{build.includes}"}}},

	// ## Compile S files
	"recipe.asm.pattern": {{Value: []string{"{compiler.asm.cmd}", "{compiler.asm.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{build.includes}"}}},

	// ## Create archives
	"recipe.ar.pattern": {{Value: []string{"{compiler.lib.cmd}", "{compiler.lib.flags}"}}},

	// ## Combine gc-sections, archives, and objects
	"recipe.link.pattern": {{Value: []string{"{compiler.link.cmd}", "{compiler.link.flags}"}}},

	// ## Compute size (text, data, bss)
	"recipe.size.pattern": {{Value: []string{"{compiler.size.cmd}", "--format=berkeley", `"{build.path}/{build.project_name}{build.link.extension}"`}}},

	// ## Boards

	// # Build options
	"build.code_debug":  {{Value: []string{"0"}}},
	"build.extra_flags": {{Value: []string{"-DCORE_DEBUG_LEVEL={build.code_debug}"}}},

	"build.defines": {
		{Value: []string{"-DTARGET_MAC"}},
		{Value: []string{"-DTARGET_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-DTARGET_RELEASE"}, Config: "release-*-*", Append: true},
		{Value: []string{"-DTARGET_FINAL"}, Config: "*-final-*", Append: true},
		{Value: []string{"-DTARGET_TEST"}, Config: "*-*-test", Append: true},
	},

	"build.includes": {},

	"build.warnings": {{Value: []string{""}}},

	"build.optimize": {
		{Value: []string{"-O0"}, Config: "debug-*-*"},
		{Value: []string{"-O2"}, Config: "release-*-*"},
		{Value: []string{"-O3"}, Config: "release-final-*"},
	},

	"build.exception_handling": {
		{Value: []string{"-fexceptions"}, Config: "*-*-test"},
	},

	"build.inline_expansion": {
		{Value: []string{"-finline-functions"}, Config: "release-*-*"},
	},

	"build.intrinsic_functions": {
		{Value: []string{"-ffunction-sections"}, Config: "release-*-*"},
	},

	"build.c_standard":   {{Value: []string{"-std=c17"}}},
	"build.cpp_standard": {{Value: []string{"-std=c++17"}}},
}
