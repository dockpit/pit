// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tool

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// downloadPackage runs the create or download command
// to make the first copy of or update a copy of the given package.
func DownloadPackage(p *Package, to string, update bool) error {
	var (
		vcs            *vcsCmd
		repo, rootPath string
		err            error
	)
	if p.SrcRoot != "" {
		// Directory exists.  Look for checkout along path to src.
		vcs, rootPath, err = vcsForDir(p)
		if err != nil {
			return err
		}
		repo = "<local>" // should be unused; make distinctive
	} else {
		// Analyze the import path to determine the version control system,
		// repository, and the import path for the root of the repository.
		rr, err := repoRootForImportPath(p.ImportPath)
		if err != nil {
			return err
		}
		vcs, repo, rootPath = rr.vcs, rr.repo, rr.root
	}

	if p.SrcRoot == "" {
		// Package not found.  Put in first directory of $GOPATH.
		list := filepath.SplitList(to)
		if len(list) == 0 {
			return fmt.Errorf("cannot download, $GOPATH not set. For more details see: go help gopath")
		}

		p.SrcRoot = filepath.Join(list[0], "deps")
	}
	root := filepath.Join(p.SrcRoot, rootPath)

	if buildV {
		fmt.Fprintf(os.Stderr, "%s (download)\n", rootPath)
	}

	// Check that this is an appropriate place for the repo to be checked out.
	// The target directory must either not exist or have a repo checked out already.
	meta := filepath.Join(root, "."+vcs.cmd)
	st, err := os.Stat(meta)
	if err == nil && !st.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", meta)
	}
	if err != nil {
		// Metadata directory does not exist.  Prepare to checkout new copy.
		// Some version control tools require the target directory not to exist.
		// We require that too, just to avoid stepping on existing work.
		if _, err := os.Stat(root); err == nil {
			return fmt.Errorf("%s exists but %s does not - stale checkout?", root, meta)
		}
		// Some version control tools require the parent of the target to exist.
		parent, _ := filepath.Split(root)
		if err = os.MkdirAll(parent, 0777); err != nil {
			return err
		}
		if err = vcs.create(root, repo); err != nil {
			return err
		}
	} else if update {

		// Metadata directory does exist; if we want to download incremental updates.
		if err = vcs.download(root); err != nil {
			return err
		}
	}

	if buildN {
		// Do not show tag sync in -n; it's noise more than anything,
		// and since we're not running commands, no tag will be found.
		// But avoid printing nothing.
		fmt.Fprintf(os.Stderr, "# cd %s; %s sync/update\n", root, vcs.cmd)
		return nil
	}

	// Select and sync to appropriate version of the repository.
	tags, err := vcs.tags(root)
	if err != nil {
		return err
	}
	vers := runtime.Version()
	if i := strings.Index(vers, " "); i >= 0 {
		vers = vers[:i]
	}
	if err := vcs.tagSync(root, selectTag(vers, tags)); err != nil {
		return err
	}

	return nil
}

// selectTag returns the closest matching tag for a given version.
// Closest means the latest one that is not after the current release.
// Version "goX" (or "goX.Y" or "goX.Y.Z") matches tags of the same form.
// Version "release.rN" matches tags of the form "go.rN" (N being a floating-point number).
// Version "weekly.YYYY-MM-DD" matches tags like "go.weekly.YYYY-MM-DD".
//
// NOTE(rsc): Eventually we will need to decide on some logic here.
// For now, there is only "go1".  This matches the docs in go help get.
func selectTag(goVersion string, tags []string) (match string) {
	for _, t := range tags {
		if t == "go1" {
			return "go1"
		}
	}
	return ""
}
