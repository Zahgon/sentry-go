package sentry

const (
	MechanismTypeGeneric string = "generic"
	MechanismTypeChained string = "chained"
	MechanismTypeUnwrap  string = "unwrap"
	MechanismSourceCause string = "cause"
)

type visited struct {
	ptrs map[uintptr]struct{}
	msgs map[string]struct{}
}

func (v *visited) seenError(err error) bool { _ = "STUB: not implemented"; return false }

func convertErrorToExceptions(err error, maxErrorDepth int) []Exception {
	_ = "STUB: not implemented"
	return nil
}

// mechanism type is used for debugging purposes, but since we can't really distinguish the origin of who invoked
// captureException, we set it to nil if the error is not chained.

// Add a trace of the current stack to the top level(outermost) error in a chain if
// it doesn't have a stack trace yet.
// We only add to the most recent error to avoid duplication and because the
// current stack is most likely unrelated to errors deeper in the chain.

func convertErrorDFS(err error, exceptions *[]Exception, parentID *int, source string, visited *visited, maxErrorDepth int, currentDepth int) {
	_ = "STUB: not implemented"
	return
}
