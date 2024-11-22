# Decision Record: go testing and testify for unit tests

|                      |              |
|----------------------|--------------|
| *Date*               | 2024-11-22   |
| *Status*             | Decided      |
| *Decision Author(s)* | Tim Steffens |

## Decision

* In the context of the knowledger project backend development,
* facing the need to write unit tests,
* we decided for the [go testing package](https://pkg.go.dev/testing) with [testify](https://github.com/stretchr/testify) assertions
* and against Ginkgo / Gomega and  GoConvey,
* to achieve an easy testing approach with and easy syntax,
* accepting that there is no special support for bdd style testing.

## Links

* [go testing package](https://pkg.go.dev/testing)
* [testify](https://github.com/stretchr/testify)