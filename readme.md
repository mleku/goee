# goee.realy.lol

A DSL and widget sets for the [GIO](https:gioui.org) immediate mode GUI library for Go.

## Versioning

This repo uses a variant of Semver that keeps the initial version number, but after that, uses numbers that 
represent the date, eg first release v0.24.12.14 which is an unstable initial version with no API stability guarantee -
ie `v0.x` with the date being 14 December 2024, in lexicographical, semver sortable order.

The rationale for this is that other than API stability guarantees, the minor and patch numbers don't really have any 
clear semantics without being associated with a feature and because Go Modules pick the semver with the biggest numeric
lexicographic value as the newest (not by date or position in the log) this scheme keeps the most recent tag as the
`latest` which is how `go mod tidy` will try to update it to.

The tag that will be applied is also maintained as accessible in the root package with the name `goee.Version` and most 
of the time the tag on a commit will match the value in the embedded version file. Possibly a convenience script could 
do this automatically but for now it will be manual.

There can be a 5th field in the version, also, this will be populated if there is more than one tagged commit added on 
the same calendar day.

This scheme is mostly ripped off the one used by Flatbuffers, except with retaining the API versioning guarantee 
semantics.