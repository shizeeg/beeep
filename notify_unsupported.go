//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !illumos && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!illumos,!js

package beeep

// Notify sends desktop notification.
func Notify(title, message string) error {
	return ErrUnsupported
}
