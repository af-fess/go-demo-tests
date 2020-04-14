package factories

import (
	"os"
	"time"
)

// Responsible to simulate "os" library.
// You can add more methods here
type SystemLibraryProvider interface {
	// os
	Chdir(dir string) error // os.Chdir

	//time
	Sleep(d time.Duration)
}

type SystemLibraryProvide struct {
}

func New()  *SystemLibraryProvide {
	return &SystemLibraryProvide{}
}

// mock helper for 'os'
func (app *SystemLibraryProvide) Chdir(dir string) error {
	return os.Chdir(dir)
}

// mock helper for 'time'
func (app *SystemLibraryProvide) Sleep(d time.Duration) {
	time.Sleep(d)
}