# Decision Record: Fiber as http framework

|                      |              |
|----------------------|--------------|
| *Date*               | 2024-11-21   |
| *Status*             | Decided      |
| *Decision Author(s)* | Tim Steffens |

## Decision

* In the context of the [knowledger project](https://github.com/tmstff/knowledger),
* facing the need to provide a RESTful API for the backend,
* we decided for [Fiber](https://github.com/gofiber/fiber) 
* and against other popular frameworks like [gin](https://github.com/gin-gonic/gin) and [echo](https://github.com/labstack/echo),
* to achieve easy usage, easily understandable and concise code, good community support, longevity and 
the possibility to try out a new framework claiming to be highly performant to gain new insights,
* accepting that other options are more widely used and might offer even more longevity and better community support.

## Details

The main USP of [Fiber](https://github.com/gofiber/fiber) - ultra high performance - is not really important for this
project (yet). But as the desired qualities (see above) are given, Fiber's an interesting candidate for gaining some
experience to maybe later use it in a performance critical context when needed.

## Links

* [Fiber](https://github.com/gofiber/fiber)
* [gin](https://github.com/gin-gonic/gin)
* [echo](https://github.com/labstack/echo)