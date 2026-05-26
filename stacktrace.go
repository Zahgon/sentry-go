package sentry

import (
	"go/build"
	"reflect"
	"runtime"
	"strings"
)

const unknown string = "unknown"

// The module download is split into two parts: downloading the go.mod and downloading the actual code.
// If you have dependencies only needed for tests, then they will show up in your go.mod,
// and go get will download their go.mods, but it will not download their code.
// The test-only dependencies get downloaded only when you need it, such as the first time you run go test.
//
// https://github.com/golang/go/issues/26913#issuecomment-411976222

// Stacktrace holds information about the frames of the stack.
type Stacktrace struct {
	Frames        []Frame `json:"frames,omitempty"`
	FramesOmitted []uint  `json:"frames_omitted,omitempty"`
}

// NewStacktrace creates a stacktrace using runtime.Callers.
func NewStacktrace() *Stacktrace { _ = "STUB: not implemented"; return nil }

// TODO: Make it configurable so that anyone can provide their own implementation?
// Use of reflection allows us to not have a hard dependency on any given
// package, so we don't have to import it.

// ExtractStacktrace creates a new Stacktrace based on the given error.
func ExtractStacktrace(err error) *Stacktrace { _ = "STUB: not implemented"; return nil }

func extractReflectedStacktraceMethod(err error) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// https://github.com/go-errors/errors

// https://github.com/pkg/errors

// https://github.com/pingcap/errors

func extractPcs(method reflect.Value) []uintptr { _ = "STUB: not implemented"; return nil }

// extractXErrorsPC extracts program counters from error values compatible with
// the error types from golang.org/x/xerrors.
//
// It returns nil if err is not compatible with errors from that package or if
// no program counters are stored in err.
func extractXErrorsPC(err error) []uintptr {
	_ = "STUB: not implemented"
	// This implementation uses the reflect package to avoid a hard dependency
	// on third-party packages.
	return nil
}

// We don't know if err matches the expected type. For simplicity, instead
// of trying to account for all possible ways things can go wrong, some
// assumptions are made and if they are violated the code will panic. We
// recover from any panic and ignore it, returning nil.
//nolint: errcheck

// type Frame struct{ frames [3]uintptr }

// drop first pc pointing to xerrors.New

// Frame represents a function call and it's metadata. Frames are associated
// with a Stacktrace.
type Frame struct {
	Function string `json:"function,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	// Module is, despite the name, the Sentry protocol equivalent of a Go
	// package's import path.
	Module      string                 `json:"module,omitempty"`
	Filename    string                 `json:"filename,omitempty"`
	AbsPath     string                 `json:"abs_path,omitempty"`
	Lineno      int                    `json:"lineno,omitempty"`
	Colno       int                    `json:"colno,omitempty"`
	PreContext  []string               `json:"pre_context,omitempty"`
	ContextLine string                 `json:"context_line,omitempty"`
	PostContext []string               `json:"post_context,omitempty"`
	InApp       bool                   `json:"in_app"`
	Vars        map[string]interface{} `json:"vars,omitempty"`
	// Package and the below are not used for Go stack trace frames.  In
	// other platforms it refers to a container where the Module can be
	// found.  For example, a Java JAR, a .NET Assembly, or a native
	// dynamic library.  They exists for completeness, allowing the
	// construction and reporting of custom event payloads.
	Package         string `json:"package,omitempty"`
	InstructionAddr string `json:"instruction_addr,omitempty"`
	AddrMode        string `json:"addr_mode,omitempty"`
	SymbolAddr      string `json:"symbol_addr,omitempty"`
	ImageAddr       string `json:"image_addr,omitempty"`
	Platform        string `json:"platform,omitempty"`
	StackStart      bool   `json:"stack_start,omitempty"`
}

// NewFrame assembles a stacktrace frame out of runtime.Frame.
func NewFrame(f runtime.Frame) Frame { _ = "STUB: not implemented"; return *new(Frame) }

// Like filepath.IsAbs() but doesn't care what platform you run this on.
// I.e. it also recognizies `/path/to/file` when run on Windows.
func isAbsPath(path string) bool { _ = "STUB: not implemented"; return false }

// If the volume name starts with a double slash, this is an absolute path.

// Windows absolute path, see https://learn.microsoft.com/en-us/dotnet/standard/io/file-path-formats

func newFrame(module string, function string, file string, line int) Frame {
	_ = "STUB: not implemented"
	return *new(Frame)
}

// Leave abspath as the empty string to be omitted when serializing event as JSON.

// TODO: in the general case, it is not trivial to come up with a
// "project relative" path with the data we have in run time.
// We shall not use filepath.Base because it creates ambiguous paths and
// affects the "Suspect Commits" feature.
// For now, leave relpath empty to be omitted when serializing the event
// as JSON. Improve this later.

// f.File is a relative path. This may happen when the binary is built
// with the -trimpath flag.

// Omit abspath when serializing the event as JSON.

// splitQualifiedFunctionName splits a package path-qualified function name into
// package name and function name. Such qualified names are found in
// runtime.Frame.Function values.
func splitQualifiedFunctionName(name string) (pkg string, fun string) {
	_ = "STUB: not implemented"
	return "", ""
}

func extractFrames(pcs []uintptr) []runtime.Frame { _ = "STUB: not implemented"; return nil }

// createFrames creates Frame objects while filtering out frames that are not
// meant to be reported to Sentry, those are frames internal to the SDK or Go.
func createFrames(frames []runtime.Frame) []Frame { _ = "STUB: not implemented"; return nil }

// Fix issues grouping errors with the new fully qualified function names
// introduced from Go 1.21

// TODO ID: why do we want to do this?
// I'm not aware of other SDKs skipping all Sentry frames, regardless of their position in the stactrace.
// For example, in the .NET SDK, only the first frames are skipped until the call to the SDK.
// As is, this will also hide any intermediate frames in the stack and make debugging issues harder.
func shouldSkipFrame(module string) bool {
	_ = "STUB: not implemented"
	// Skip Go internal frames.
	return false
}

// Skip Sentry internal frames, except for frames in _test packages (for testing).

// On Windows, GOROOT has backslashes, but we want forward slashes.
var goRoot = strings.ReplaceAll(build.Default.GOROOT, "\\", "/")

func setInAppFrame(frame *Frame) { _ = "STUB: not implemented"; return }

func callerFunctionName() string { _ = "STUB: not implemented"; return "" }

// packageName returns the package part of the symbol name, or the empty string
// if there is none.
// It replicates https://golang.org/pkg/debug/gosym/#Sym.PackageName, avoiding a
// dependency on debug/gosym.
func packageName(name string) string { _ = "STUB: not implemented"; return "" }

// baseName returns the symbol name without the package or receiver name.
// It replicates https://golang.org/pkg/debug/gosym/#Sym.BaseName, avoiding a
// dependency on debug/gosym.
func baseName(name string) string { _ = "STUB: not implemented"; return "" }

func isCompilerGeneratedSymbol(name string) bool {
	_ = "STUB: not implemented"
	// In versions of Go 1.20 and above a prefix of "type:" and "go:" is a
	// compiler-generated symbol that doesn't belong to any package.
	// See variable reservedimports in cmd/compile/internal/gc/subr.go
	return false
}

// Walk backwards through the results and for the current function name
// remove it's parent function's prefix, leaving only it's actual name. This
// fixes issues grouping errors with the new fully qualified function names
// introduced from Go 1.21.
func cleanupFunctionNamePrefix(f []Frame) []Frame { _ = "STUB: not implemented"; return nil }
