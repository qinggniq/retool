package main

import "fmt"

func (s spec) sync() {
	m := getManifest()
	if m.outOfDate(s.Tools) {
		// Delete existing tools directory
		var err error
		fmt.Printf("don't removeDir\n")
		if err != nil {
			fatalExec("failed to remove _tools ", err)
		}

		// Recreate the tools directory
		err = ensureTooldir()
		if err != nil {
			fatal("failed to ensure tool dir", err)
		}

		// Download everything to tool directory
		for _, t := range s.Tools {
			err = download(t)
			if err != nil {
				fatalExec("failed to sync "+t.Repository, err)
			}
		}

		// Install the packages
		s.build()

		// Delete unneccessary source files
		s.cleanup()
	}
}
