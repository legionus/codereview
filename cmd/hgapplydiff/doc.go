// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

Hgapplydiff applies a patch to the local Mercurial repository.
The patch should have been been generated by
a version control system like CVS, Git, Mercurial, or Subversion.
If successful, hgapplydiff writes a list of affected files to standard output.

Hgapplydiff is meant to be used by the Mercurial codereview extension.
It was formerly named hgpatch, but executables containing the name
patch cause problems on Windows.

Usage:
	hgapplydiff [patchfile]

*/
package documentation
