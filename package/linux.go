package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
)

func getVarsLinux(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, vars *corepkg.Vars) {

}

// TODO; implement Linux vars (currently a copy from Mac)
var platformVarsLinux = map[string][]denv.Var{
	// # Extensions
	"build.obj.prefix": {{Value: []string{""}}},       // Object file prefix
	"build.obj.suffix": {{Value: []string{".o"}}},     // Object file suffix
	"build.dep.prefix": {{Value: []string{""}}},       // Dependency file prefix
	"build.dep.suffix": {{Value: []string{".d"}}},     // Dependency file suffix
	"build.dll.prefix": {{Value: []string{"lib"}}},    // Dynamic library file prefix
	"build.dll.suffix": {{Value: []string{".dylib"}}}, // Dynamic library file suffix
	"build.lib.prefix": {{Value: []string{"lib"}}},    // Static library file prefix
	"build.lib.suffix": {{Value: []string{".a"}}},     // Static library file suffix
	"build.exe.prefix": {{Value: []string{""}}},       // Executable file prefix (none for macOS)
	"build.exe.suffix": {{Value: []string{""}}},       // Executable file suffix (none for macOS)

	// # Frameworks
	"compiler.frameworks.default": {{Value: []string{"-framework", "Foundation", "-framework", "CoreFoundation", "-framework", "IOKit", "-framework", "CoreServices"}}},
	"compiler.frameworks.metal":   {{Value: []string{"-framework", "Metal", "-framework", "MetalPerformanceShaders", "-framework", "MetalKit"}}},
	"compiler.frameworks.cocoa":   {{Value: []string{"-framework", "Cocoa"}}},
	"compiler.frameworks.appkit":  {{Value: []string{"-framework", "AppKit"}}},
	"compiler.frameworks.uikit":   {{Value: []string{"-framework", "UIKit"}}},

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
	"compiler.warning_flags": {{Value: []string{"-Wall", "-Wextra"}}},

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
		{Value: []string{`-c","-MMD","{compiler.warning_flags}","{compiler.common_werror_flags}","{compiler.floating_point_flags}","{compiler.c.flags}","{build.cpp_standard}`}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
		{Value: []string{"-fexceptions"}, Config: "*-*-test", Append: true},
	},

	"compiler.cpp.flags": {
		{Value: []string{`-c","-MMD","{compiler.warning_flags}","{compiler.common_werror_flags}","{compiler.floating_point_flags}","{compiler.cpp.flags}`}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
		{Value: []string{"-fexceptions"}, Config: "*-*-test", Append: true},
	},

	"compiler.asm.flags": {
		{Value: []string{"-c", "-x", "-MMD", "assembler-with-cpp", "{compiler.warning_flags}", "{compiler.asm.flags}"}},
		{Value: []string{"-g", "-O0", "-fno-omit-frame-pointer", "-D_DEBUG"}, Config: "debug-*-*", Append: true},
		{Value: []string{"-O2", "-finline-functions", "-ffunction-sections", "-DNDEBUG"}, Config: "release-*-*", Append: true},
	},

	"compiler.lib.flags": {
		{Value: []string{"-rs", "{compiler.lib.flags}"}},
		{Value: []string{""}, Config: "debug-*-*", Append: true},
	},

	"compiler.link.flags": {
		{Value: []string{"-Wl,--Map={build.path}/{build.project_name}.map", "{compiler.link.flags}", "{compiler.frameworks}"}},
		{Value: []string{"-Wl,-nostrip"}, Config: "*-dev-*", Append: true},
		{Value: []string{"-flto", "-Wl,-Strip-all"}, Config: "*-final-*", Append: true},
	},

	// ## Compile c files
	"recipe.c.pattern": {{Value: []string{`{compiler.c.cmd}`, "{compiler.c.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{include_paths}"}}},

	// ## Compile c++ files
	"recipe.cpp.pattern": {{Value: []string{"{compiler.cpp.cmd}", "{compiler.cpp.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{include_paths}"}}},

	// ## Compile S files
	"recipe.asm.pattern": {{Value: []string{"{compiler.asm.cmd}", "{compiler.asm.flags}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{build.defines}", "{include_paths}"}}},

	// ## Create archives
	"recipe.ar.pattern": {{Value: []string{"{compiler.lib.cmd}", "{compiler.lib.flags}", `"{archive_file_path}"`, "{object_files}"}}},

	// ## Combine gc-sections, archives, and objects
	"recipe.link.pattern": {{Value: []string{"{compiler.link.cmd}", "{compiler.link.flags}", "-Wl,--start-group", "{object_files}", "{build.extra_libs}", "-Wl,--end-group", "-Wl,-EL", "-o", `"{build.path}/{build.project_name}{build.link.extension}"`}}},

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

	"build.cpp_standard": {{Value: []string{"-std=c++17"}}},
	"build.c_standard":   {{Value: []string{"-std=c17"}}},
}
