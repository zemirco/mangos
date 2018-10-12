## mangos <img src=mangos.jpg align=right>

[![Linux Status](https://img.shields.io/circleci/project/github/nanomsg/mangos.svg?label=linux)](https://circleci.com/gh/nanomsg/mangos)
[![Windows Status](https://img.shields.io/appveyor/ci/nanomsg/mangos.svg?label=windows)](https://ci.appveyor.com/project/nanomsg/mangos)
[![Apache License](https://img.shields.io/badge/license-APACHE2-blue.svg)](https://github.com/nanomsg/mangos/blob/master/LICENSE)
[![Gitter](https://img.shields.io/badge/gitter-join-brightgreen.svg)](https://gitter.im/go-mangos/mangos)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/nanomsg.org/go-mangos)
[![Go Report Card](https://goreportcard.com/badge/nanomsg.org/go-mangos)](https://goreportcard.com/report/nanomsg.org/go-mangos)

> WARNING:  THIS IS A WORK-IN-PROGRESS AND NOT READY FOR GENERAL USE YET.
> This represents version 2 of mangos, and contains breaking changes and
> rather different API.  To use the the original version of mangos please
> see https://github.com/nanomsg/mangos

Package mangos is an implementation in pure Go of the SP
("Scalability Protocols") messaging system.
These are colloquially  known as a "nanomsg".

The modern C implemenation of the SP protocols is available as
[NNG&trade;](https://github.com/nanomsg/nng).

The original implementation of the SP protocols is available as
[nanomsg&trade;](http://www.nanomsg.org).

Generally (modulo a few caveats) all of these implementations can
interoperate.

The design is intended to make it easy to add new transports with almost trivial
effort, as well as new topologies ("protocols" in SP terminology.)

At present, all of the Req/Rep, Pub/Sub, Pair, Bus, Push/Pull, and
Surveyor/Respondent patterns are supported.

Additionally, there is an experimental new pattern called STAR available.  This
pattern is like Bus, except that the messages are delivered not just to
immediate peers, but to all members of the topology.  Developers must be careful
not to create cycles in their network when using this pattern, otherwise
infinite loops can occur.

Supported transports include TCP, inproc, IPC, Websocket, Websocket/TLS and TLS.
Use addresses of the form "tls+tcp://<host>:<port>" to access TLS.
Note that ipc:// is not supported on Windows (by either this or the reference
implementation.)  Forcing the local TCP port in Dial is not supported yet (this
is rarely useful).

Basic interoperability with nanomsg and NNG has been verified (you can do
so yourself with nanocat and macat) for all protocols and transports
that NNG and nanomsg support.
Additionally there are a number of projects that use the two products together.

There is a third party experimental QUIC transport available at
[quic-mangos](https://github.com/lthibault/quic-mangos).  (An RFE to make this
transport official exists.)

If you find this useful, I would appreciate knowing about it.  I can be reached
via my email address, garrett -at- damore -dot- org

## Installing

### Using *go get*

    $ go get -u nanomsg.org/go-mangos

After this command *mangos* is ready to use. Its source will be in:

    $GOPATH/src/pkg/nanomsg.org/go-mangos

You can use `go get -u -a` to update all installed packages.

## Documentation

For docs, see http://godoc.org/nanomsg.org/go-mangos or run:

    $ godoc nanomsg.org/go-mangos

## Testing

This package supports internal self tests, which can be run in
the idiomatic Go way.  (Note that most of the tests are in a test
subdirectory.)

    $ go test nanomsg.org/go-mangos/...

There are also internal benchmarks available:

    $ go test -bench=. nanomsg.org/go-mangos/test

## Commercial Support

[Staysail Systems, Inc.](mailto:info@staysail.tech) offers
[commercial support](http://staysail.tech/support/mangos) for mangos.

## Examples

Some examples are posted in the directories under examples/
in this project.

These examples are rewrites (in Go) of Tim Dysinger's
[Getting Started with Nanomsg](http://nanomsg.org/gettingstarted/index.html).

godoc in the example directories will yield information about how to run
each example program.

Enjoy!

Copyright 2018 The Mangos Authors

mangos&trade;, Nanomsg&trade; and NNG&trade; are [trademarks](http://nanomsg.org/trademarks.html) of Garrett D'Amore.
