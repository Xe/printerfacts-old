# Contributing #

Thanks for helping make Twirp better! This is great!

First, if you have run into a bug, please file an issue. We try to get back to
issue reporters within a day or two. We might be able to help you right away.

If you'd rather not publicly discuss the issue, please email spencer@twitch.tv
and/or security@twitch.tv.

Issues are also a good place to present experience reports or requests for new
features.

If you'd like to make changes to Twirp, read on:

## Setup Requirements ##

You will need git, Go 1.9+, and Python 2.7 installed and on your system's path.
Install them however you feel.

We work on a branch called `develop`. We periodically release this branch as a
new version, then accumulate more changes on the develop branch until the next
release. Use `develop` as the base for your branches.

## Developer Loop ##

Generally you want to make changes and run `make`, which will install all
dependencies we know about, build the core, and run all of the tests that we
have against all of the languages we support.

Most tests of the Go server are in `internal/twirptest/service_test.go`. Tests
of cross-language clients are in the [clientcompat](./clientcompat) directory.

## Contributing Code ##

Twirp uses github pull requests. Fork a branch from `develop`, hack away at your
changes, run the test suite with `make`, and submit a PR.

## Releasing Versions ##

Releasing versions is the responsibility of the core maintainers. Most people
don't need to know this stuff.

Twirp uses [Semantic versioning](http://semver.org/): `v<major>.<minor>.<patch>`.

 * Increment major if you're making a backwards-incompatible change.
 * Increment minor if you're adding a feature that's backwards-compatible.
 * Increment patch if you're making a bugfix.

To make a release, remember to update the version number in
[internal/gen/version.go](./internal/gen/version.go).

Twirp uses Github releases. To make a new release:
 1. Merge all changes that should be included in the release into the master
    branch.
 2. Add a new commit to master with a message like "Version vX.X.X release".
 3. Tag the commit you just made: `git tag <version number>` and `git push
    origin --tags`
 3. Go to Github https://github.com/twitchtv/twirp/releases and
    "Draft a new release".
 4. Make sure to document changes, specially when upgrade instructions are
    needed.
