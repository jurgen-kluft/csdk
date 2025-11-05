package csdk

import (
	corepkg "github.com/jurgen-kluft/ccode/core"
	denv "github.com/jurgen-kluft/ccode/denv"
)

func getVarsWindows(buildTarget denv.BuildTarget, buildConfig denv.BuildConfig, vars *corepkg.Vars) {

	// TODO
	//    - figure out the Windows SDK versions dynamically
	//    - figure out the MSVC versions dynamically

	// Using the buildTarget and buildConfig, we iterate over the platformVarsWindows map and
	// set the appropriate variables in the vars object. Some variables may depend on the buildConfig.
	for key, varList := range platformVarsWindows {
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

var platformVarsWindows = map[string][]denv.Var{
	// # Extensions
	"build.obj.prefix": {{Value: []string{""}}},     // Object file prefix
	"build.obj.suffix": {{Value: []string{".obj"}}}, // Object file suffix
	"build.dep.prefix": {{Value: []string{""}}},     // Dependency file prefix
	"build.dep.suffix": {{Value: []string{".dep"}}}, // Dependency file suffix
	"build.dll.prefix": {{Value: []string{""}}},     // Dynamic library file prefix
	"build.dll.suffix": {{Value: []string{".dll"}}}, // Dynamic library file suffix
	"build.lib.prefix": {{Value: []string{""}}},     // Static library file prefix
	"build.lib.suffix": {{Value: []string{".lib"}}}, // Static library file suffix
	"build.exe.prefix": {{Value: []string{""}}},     // Executable file prefix
	"build.exe.suffix": {{Value: []string{".exe"}}}, // Executable file suffix

	// # Micrsoft Visual Studio specific paths
	"msvc-libs.path":  {{Value: []string{`{runtime.platform.path}/tools/msvc-libs`}}},
	"msvc-tools.path": {{Value: []string{``}}},

	// # Windows SDK (Windows 10/11) specific paths
	"windows.sdk.path": {{Value: []string{`C:/Program Files (x86)/Windows Kits/10/10.22000.0`}}},

	// # Object files output directory
	"build.objdir": {{Value: []string{`/Fo{build.path}\`}}},

	// # Source dependencies
	"compiler.source_dependencies": {{Value: []string{`/sourceDependencies`, `"{build.path}"`}}},

	// # Debug Info
	"compiler.debug_flags": {{Config: "debug-*-*", Append: true, Value: []string{`{msvc.cl.generatedebuginfo}`}}},

	// # Floating Point
	"compiler.floating_point_flags": {{Value: []string{`{msvc.cl.usefloatingpointprecise}`}}},

	// # SSE / AVX
	"compiler.sse_flags.x64": {{Value: []string{`/arch:AVX`}}}, // Enable AVX instructions for x64 architectur}e

	// # Optimization flags
	"compiler.optimization_flags": {
		{Config: "debug-*-*", Value: []string{`{msvc.cl.optimize_none}`}},
		{Config: "release-dev-*", Value: []string{`{msvc.cl.optimize_size}`}},
		{Config: "release-final-*", Value: []string{`{msvc.cl.optimize_speed}`}},
	},

	// # Compile Warning Levels
	"compiler.warning_flags": {{Value: []string{"{msvc.cl.warnings_level3}", "{msvc.cl.warnings_are_errors}"}}},

	// # Compile Flags
	"compiler.cpreprocessor.flags": {{Value: []string{``}}},
	"compiler.flags":               {{Value: []string{"{msvc.cl.compileonly}", "{msvc.cl.nologo}", "{msvc.cl.diagnostics_columnmode}", "{msvc.cl.diagnostics_emitfullpathofsourcefiles}", "{msvc.cl.buildmultiplesourcefilesconcurrently}", "{compiler.warning_flags}", "{compiler.floating_point_flags}", "{compiler.source_dependencies}", "{compiler.debug_flags}"}}},
	"compiler.c.flags":             {{Value: []string{"{compiler.flags}"}}},
	"compiler.cpp.flags":           {{Value: []string{"{compiler.flags}"}}},
	"compiler.asm.flags":           {{Value: []string{""}}},
	"compiler.c.link.flags":        {{Value: []string{""}}},
	"compiler.c.link.libs":         {{Value: []string{""}}},
	"compiler.lib.flags":           {{Value: []string{""}}},

	// # Compiler Extra Flags
	"compiler.c.extra_flags":      {{Value: []string{""}}},
	"compiler.cpp.extra_flags":    {{Value: []string{""}}},
	"compiler.asm.extra_flags":    {{Value: []string{""}}},
	"compiler.c.link.extra_flags": {{Value: []string{""}}},
	"compiler.lib.extra_flags":    {{Value: []string{""}}},

	// # Compilers (Microsoft Visual Studio)
	"compiler.c.cmd":    {{Value: []string{"cl.exe"}}},
	"compiler.cpp.cmd":  {{Value: []string{"cl.exe"}}},
	"compiler.asm.cmd":  {{Value: []string{"cl.exe"}}},
	"compiler.lib.cmd":  {{Value: []string{"lib.exe"}}},
	"compiler.link.cmd": {{Value: []string{"link.exe"}}},
	"compiler.size.cmd": {{Value: []string{"size.exe"}}},

	// ## Compile c files
	"recipe.c.pattern": {{Value: []string{`"{compiler.c.cmd}"`, "{compiler.c.extra_flags}", "{compiler.c.flags}", "{build.warnings}", "{build.optimize}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{include_paths}"}}},

	// ## Compile c++ files
	"recipe.cpp.pattern": {{Value: []string{`"{compiler.cpp.cmd}"`, "{compiler.cpp.extra_flags}", "{compiler.cpp.flags}", "{build.optimize}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{include_paths}"}}},

	// ## Compile asm files
	"recipe.asm.pattern": {{Value: []string{`"{compiler.c.cmd}"`, "{compiler.S.extra_flags}", "{compiler.S.flags}", "{build.optimize}", "{build.extra_flags}", "{compiler.cpreprocessor.flags}", "{include_paths}"}}},

	// ## Create archives/libraries
	"recipe.lib.pattern": {{Value: []string{`"{compiler.lib.cmd}"`, "{compiler.ar.flags}", "{compiler.ar.extra_flags}"}}},

	// ## Combine libraries, and object files
	"recipe.link.pattern": {{Value: []string{`"{compiler.link.cmd}"`, "{compiler.c.link.flags}", "{compiler.c.link.extra_flags}", "{libpaths}", "{libfiles}"}}},

	// ## Compute size (text, data, bss)
	"recipe.size.pattern": {{Value: []string{`"{compiler.size.cmd}"`, "--format=berkeley"}}},

	// ## Microsoft Visual Studio compiler (cl) flags
	"msvc.cl.compileonly":                           {{Value: []string{`/c`}}},                  // Compile only; do not link. This is useful for generating object files without creating an executable.
	"msvc.cl.nologo":                                {{Value: []string{`/nologo`}}},             // Suppress the display of the compiler's startup banner and copyright message.
	"msvc.cl.diagnostics_columnmode":                {{Value: []string{`/diagnostics:column`}}}, // Enable column mode for diagnostics, which provides more detailed error messages.
	"msvc.cl.diagnostics_emitfullpathofsourcefiles": {{Value: []string{`/FC`}}},                 // Full path of source files in diagnostics.
	"msvc.cl.warnings_level3":                       {{Value: []string{`/W3`}}},                 // Set output warning level to 3 (high warnings).
	"msvc.cl.warnings_are_errors":                   {{Value: []string{`/WX`}}},                 // Treat warnings as errors.
	"msvc.cl.warnings_disable_all":                  {{Value: []string{`/w`}}},                  // Disable all warnings. This is generally not recommended, but can be useful in certain scenarios where you want to suppress all warnings.
	"msvc.cl.buildmultiplesourcefilesconcurrently":  {{Value: []string{`/MP`}}},                 // Build multiple source files concurrently.
	"msvc.cl.generatedebuginfo":                     {{Value: []string{`/Zi`}}},                 // Generate complete debugging information.
	"msvc.cl.disableframepointer":                   {{Value: []string{`/Oy`}}},                 // Do not omit frame pointer.
	"msvc.cl.optimize_none":                         {{Value: []string{`/Od`}}},                 // Disable optimizations for debugging.
	"msvc.cl.optimize_size":                         {{Value: []string{`/O1`}}},                 // Optimize for size.
	"msvc.cl.optimize_speed":                        {{Value: []string{`/O2`}}},                 // Optimize for speed.
	"msvc.cl.inline_expansion_level_0":              {{Value: []string{`/b0`}}},                 // Enable inline expansion for functions that are small and frequently called.
	"msvc.cl.inline_expansion_level_1":              {{Value: []string{`/b1`}}},                 // Enable inline expansion for functions that are small and frequently called, as well as for functions that are explicitly marked with the inline keyword.
	"msvc.cl.inline_expansion_level_2":              {{Value: []string{`/b2`}}},                 // Enable inline expansion for functions that are small and frequently called, as well as for functions that are explicitly marked with the inline keyword.
	"msvc.cl.inline_expansion_level_3":              {{Value: []string{`/b3`}}},                 // Enable inline expansion for all functions, regardless of their size or frequency of use.
	"msvc.cl.enable_intrinsicfunctions":             {{Value: []string{`/Oi`}}},                 // Enable intrinsic functions.
	"msvc.cl.enable_stringpooling":                  {{Value: []string{`/GF`}}},                 // Enable string pooling, which can reduce the size of the generated code by sharing identical string literals.
	"msvc.cl.enable_wholeprogramoptimization":       {{Value: []string{`/Gw`}}},                 // Enable whole program optimization, which allows the compiler to optimize across translation units.
	"msvc.cl.enable_exceptionhandling":              {{Value: []string{`/EHsc`}}},               // Enable C++ exception handling.
	"msvc.cl.rtti.enable":                           {{Value: []string{`/GR`}}},                 // Enable Run-Time Type Information (RTTI), which allows for dynamic type identification and safe downcasting in C++.
	"msvc.cl.rtti.disable":                          {{Value: []string{`/GR`}}},                 // Disable Run-Time Type Information (RTTI), which can reduce code size and improve performance if RTTI is not needed.
	"msvc.cl.omitframepointer":                      {{Value: []string{`/Oy`}}},                 // Omit frame pointer for functions that do not require one.
	"msvc.cl.usemultithreadeddebugruntime":          {{Value: []string{`/MTd`}}},                // Use the multithreaded debug version of the C runtime library.
	"msvc.cl.usemultithreadedruntime":               {{Value: []string{`/MT`}}},                 // Use the multithreaded version of the C runtime library.
	"msvc.cl.usefloatingpointprecise":               {{Value: []string{`/fp:precise`}}},         // Use floating-point model: precise
	"msvc.cl.usecpp14":                              {{Value: []string{`/std:c++14`}}},          // Use C++14 standard.
	"msvc.cl.usecpp17":                              {{Value: []string{`/std:c++17`}}},          // Use C++17 standard.
	"msvc.cl.usecpp20":                              {{Value: []string{`/std:c++20`}}},          // Use C++20 standard.
	"msvc.cl.usecpplatest":                          {{Value: []string{`/std:c++latest`}}},      // Use the latest C++ standard.
	"msvc.cl.usec11":                                {{Value: []string{`/std:c11`}}},            // Use C11 standard.
	"msvc.cl.usec17":                                {{Value: []string{`/std:c17`}}},            // Use C17 standard.
	"msvc.cl.useclatest":                            {{Value: []string{`/std:clatest`}}},        // Use the latest C standard.

	// ## Micrsoft Visual Studio archiver (lib) flags
	"msvc.lib.nologo":     {{Value: []string{`/nologo`}}},      // Suppress the display of the lib's startup banner and copyright message.
	"msvc.lib.machinex64": {{Value: []string{`/MACHINE:X64`}}}, // Target the x64 architecture.
	"msvc.lib.outfile":    {{Value: []string{`/OUT:`}}},        // Specify the name of the output library file.

	// ## Micrsoft Visual Studio linker (link) flags
	"msvc.link.errorreportprompt":             {{Value: []string{`/ERRORREPORT:PROMPT`}}}, // Prompt the user for error reporting.
	"msvc.link.nologo":                        {{Value: []string{`/NOLOGO`}}},             // Suppress the display of the linker's startup banner and copyright message.
	"msvc.link.generatemapfile":               {{Value: []string{`/MAP:`}}},               // Generate a map file.
	"msvc.link.generatedebuginfo":             {{Value: []string{`/DEBUG`}}},              // Generate debug information.
	"msvc.link.optimizereferences":            {{Value: []string{`/OPT:REF`}}},            // Eliminate unreferenced functions and data.
	"msvc.link.optimizeidenticalfolding":      {{Value: []string{`/OPT:ICF`}}},            // Perform identical COMDAT folding.
	"msvc.link.linktimecodegeneration":        {{Value: []string{`/LTCG`}}},               // Enable link-time code generation.
	"msvc.link.disableincrementallinking":     {{Value: []string{`/INCREMENTAL:NO`}}},     // Disable incremental linking.
	"msvc.link.subsystemconsole":              {{Value: []string{`/SUBSYSTEM:CONSOLE`}}},  // Specify the subsystem for the application.
	"msvc.link.subsystemwindows":              {{Value: []string{`/SUBSYSTEM:WINDOWS`}}},  // Specify the subsystem for the application.
	"msvc.link.dynamicbase":                   {{Value: []string{`/DYNAMICBASE`}}},        // Enable address space layout randomization (ASLR).
	"msvc.link.enabledataexecutionprevention": {{Value: []string{`/NXCOMPAT`}}},           // Enable Data Execution Prevention (DEP).
	"msvc.link.machinex64":                    {{Value: []string{`/MACHINE:X64`}}},        // Target the x64 architecture.
	"msvc.link.libpath":                       {{Value: []string{`/LIBPATH:`}}},           // Specify a directory to search for library files.
	"msvc.link.outfile":                       {{Value: []string{`/OUT:`}}},               // Specify the name of the output executable file.

	// --------------------------------------------------------------------------------------------------
	// # Build Options

	"build.code_debug":  {{Value: []string{`0`}}},                                                   // Code debug level (0: none, 1: basic, 2: full)
	"build.extra_flags": {{Append: true, Value: []string{`/DCORE_DEBUG_LEVEL={build.code_debug}`}}}, //
	"build.extra_libs":  {{Append: true, Value: []string{``}}},                                      //

	"build.defines": {
		{Config: "*-*-*", Append: true, Value: []string{`/DTARGET_PC`}},            // Define target as PC
		{Config: "debug-*-*", Append: true, Value: []string{`/DTARGET_DEBUG`}},     //
		{Config: "release-*-*", Append: true, Value: []string{`/DTARGET_RELEASE`}}, //
		{Config: "*-final-*", Append: true, Value: []string{`/DTARGET_FINAL`}},     //
		{Config: "*-*-test", Append: true, Value: []string{`/DTARGET_TEST`}},       //
	},

	"build.warnings": {{Append: true, Value: []string{``}}},

	"build.optimize": {
		{Config: "debug-*-*", Value: []string{`{msvc.cl.optimize_none}`}},        // No optimization
		{Config: "release-*-*", Value: []string{`{msvc.cl.optimize_size}`}},      // Optimize for size
		{Config: "release-final-*", Value: []string{`{msvc.cl.optimize_speed}`}}, // Optimize for speed
	},

	"build.exception_handling": {{Config: "*-*-test", Value: []string{`{msvc.cl.enable_exceptionhandling}`}}}, // Enable C++ exception handling for test builds

	"build.inline_expansion":    {{Value: []string{`{msvc.cl.inline_expansion_level_2}`}}},  //
	"build.intrinsic_functions": {{Value: []string{`{msvc.cl.enable_intrinsicfunctions}`}}}, //
	"build.cpp_standard":        {{Value: []string{`{msvc.cl.usecpp17}`}}},                  //
	"build.c_standard":          {{Value: []string{`{msvc.cl.usec17}`}}},                    //
}
