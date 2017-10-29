[![Build status][buildimage]][build] [![Coverage][codecovimage]][codecov] [![GoReportCard][cardimage]][card] [![API documentation][docsimage]][docs]

# This project is parked and n longer maintained

# git ignore parser and matcher for Go
 
The module implements matching file system paths to gitignore patterns that can be automatically read from 
  a git repository tree in the order of definition priorities. It support all pattern formats as specified
  in the original `gitignore` documentation, copied below:

## Pattern format

* A blank line matches no files, so it can serve as a separator for readability.

* A line starting with # serves as a comment. Put a backslash ("\\") 
in front of the first hash for patterns that begin with a hash.

* Trailing spaces are ignored unless they are quoted with backslash ("\\").

* An optional prefix "!" which negates the pattern; any matching file excluded 
by a previous pattern will become included again. It is not possible to 
re-include a file if a parent directory of that file is excluded. 
Git doesn’t list excluded directories for performance reasons, so 
any patterns on contained files have no effect, no matter where they are 
defined. Put a backslash ("\\") in front of the first "!" for patterns 
that begin with a literal "!", for example, "\\!important!.txt".

* If the pattern ends with a slash, it is removed for the purpose of the 
following description, but it would only find a match with a directory. 
In other words, foo/ will match a directory foo and paths underneath it, 
but will not match a regular file or a symbolic link foo (this is consistent 
with the way how pathspec works in general in Git).

* If the pattern does not contain a slash /, Git treats it as a shell glob 
pattern and checks for a match against the pathname relative to the location 
of the .gitignore file (relative to the toplevel of the work tree if not 
from a .gitignore file).

* Otherwise, Git treats the pattern as a shell glob suitable for consumption 
by fnmatch(3) with the FNM_PATHNAME flag: wildcards in the pattern will 
not match a / in the pathname. For example, "Documentation/*.html" matches 
"Documentation/git.html" but not "Documentation/ppc/ppc.html" or 
"tools/perf/Documentation/perf.html".

* A leading slash matches the beginning of the pathname. For example, 
"/*.c" matches "cat-file.c" but not "mozilla-sha1/sha1.c".

Two consecutive asterisks ("**") in patterns matched against full pathname 
may have special meaning:

* A leading "\*\*" followed by a slash means match in all directories. 
For example, "\*\*/foo" matches file or directory "foo" anywhere, the same as 
pattern "foo". "**/foo/bar" matches file or directory "bar" 
anywhere that is directly under directory "foo".

* A trailing "/\*\*" matches everything inside. For example, "abc/**" matches 
all files inside directory "abc", relative to the location of the 
.gitignore file, with infinite depth.

* A slash followed by two consecutive asterisks then a slash matches 
zero or more directories. For example, "a/**/b" matches "a/b", "a/x/b", 
"a/x/y/b" and so on.

* Other consecutive asterisks are considered invalid.

### License and copyright

	Copyright (c) 2017. Oleg Sklyar and teris.io. MIT license applies. All rights reserved.


[build]: https://travis-ci.org/teris-io/gitignore
[buildimage]: https://travis-ci.org/teris-io/gitignore.svg?branch=master

[codecov]: https://codecov.io/github/teris-io/gitignore?branch=master
[codecovimage]: https://codecov.io/github/teris-io/gitignore/coverage.svg?branch=master

[card]: http://goreportcard.com/report/teris-io/gitignore
[cardimage]: https://goreportcard.com/badge/github.com/teris-io/gitignore

[docs]: https://godoc.org/github.com/teris-io/gitignore
[docsimage]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat
