# *gd3*: next generation GlusterFS management!

[![gd3!](art/gd3.png)](art/)

[![Go Report Card](https://goreportcard.com/badge/github.com/purpleidea/gd3?style=flat-square)](https://goreportcard.com/report/github.com/purpleidea/gd3)
[![Build Status](https://img.shields.io/travis/purpleidea/gd3/master.svg?style=flat-square)](http://travis-ci.org/purpleidea/gd3)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/purpleidea/gd3)
[![IRC](https://img.shields.io/badge/irc-%23mgmtconfig-brightgreen.svg?style=flat-square)](https://webchat.freenode.net/?channels=#mgmtconfig)

## Community:
Come join us in the `gd3` community! We're sharing it with the `mgmt` crew:

| Medium | Link |
|---|---|---|
| IRC | [#mgmtconfig](https://webchat.freenode.net/?channels=#mgmtconfig) on Freenode |
| Twitter | [@mgmtconfig](https://twitter.com/mgmtconfig) & [#mgmtconfig](https://twitter.com/hashtag/mgmtconfig) |
| Mailing list | [mgmtconfig-list@redhat.com](https://www.redhat.com/mailman/listinfo/mgmtconfig-list) |

## Status:
Gd3 is a fairly new project, which is using mgmt as its core library.
We aim to provide powerful automation for GlusterFS and to demonstrate that mgmt
can be a useful foundation for your existing distributed software project.
We're working towards being minimally useful for production environments.
We aren't feature complete for what we'd consider a 1.x release yet.
With your help you'll be able to influence our design and get us there sooner!

## Documentation:
Please read, enjoy and help improve our documentation!

| Documentation | Additional Notes |
|---|---|
| [general documentation](docs/documentation.md) | for everyone |
| [godoc API reference](https://godoc.org/github.com/purpleidea/gd3) | for gd3 developers |

## Questions:
Please ask in the [community](#community)!
If you have a well phrased question that might benefit others, consider asking
it by sending a patch to the documentation [FAQ](https://github.com/purpleidea/gd3/blob/master/docs/documentation.md#usage-and-frequently-asked-questions) section.
I'll merge your question, and a patch with the answer!

## Roadmap:
Please see: [TODO.md](TODO.md) for a list of upcoming work and TODO items.
Please get involved by working on one of these items or by suggesting something else!
Feel free to grab one of the straightforward [#gd3love](https://github.com/purpleidea/gd3/labels/gd3love)
issues if you're a first time contributor to the project or if you're unsure
about what to hack on!

## Bugs:
Please set the `DEBUG` constant in [main.go](https://github.com/purpleidea/gd3/blob/master/main.go)
to `true`, and post the logs when you report the [issue](https://github.com/purpleidea/gd3/issues).
Bonus points if you provide a reproducible test case.
Feel free to read my article on [debugging golang programs](https://ttboj.wordpress.com/2016/02/15/debugging-golang-programs/).

## Patches:
We'd love to have your patches! Please send them by email, or as a pull request.

## On the web:
| Author | Format | Subject |
|---|---|---|
| James Shubin | blog | [Next generation configuration mgmt](https://ttboj.wordpress.com/2016/01/18/next-generation-configuration-mgmt/) |

##

Happy hacking!
