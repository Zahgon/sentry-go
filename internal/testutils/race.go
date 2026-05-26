//go:build !race

package testutils

func IsRaceTest() bool { _ = "STUB: not implemented"; return false }
