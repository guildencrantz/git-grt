# grt

grt aims to unify git and [Gerrit Code Review](http://code.google.com/p/gerrit/)
by unifying gerrit commands into the git interface.

## Installation

No packaged version is currently available, so you'll need to have [go installed](http://golang.org/doc/install).

Once go is installed just run:

`go get github.com/guildencrantz/git-grt`

Make sure that `$GOPATH/bin` is in your `$PATH` and you should be set.

## Usage

Yeah, we really should write some real documentation. Even a --help. For now:

* `git grt list`
* `git grt details _id_`
* `git grt track _id_`
* `git grt fetch`

Id can be found in `git grt list`, or in the gerrit http URI.

## History

grt was prototyped during the Spring 2014 hack-@-thon at Return Path with the
intention of helping simplify the workflow for developers sick of switching between
the git CLI and gerrit web page.

If quarterly hackathons and challenging problems sound interesting to you check
out [the current open positions at Return Path](http://jobvite.com/m?3xCrlfwF).

# Copyright and License

Copyright (c) 2014 Matt Henkel. Licensed under the Apache License (see LICENSE
for details).
