# Decision Record: Y-Statements for Decision Records

|                      |              |
|----------------------|--------------|
| *Date*               | 2024-11-21   |
| *Status*             | Decided      |
| *Decision Author(s)* | Tim Steffens |

## Decision

* In the context of the [knowledger project](https://github.com/tmstff/knowledger),
* facing the need to document decisions to be able to understand them later on,
* we decided for [Y-Statements](https://medium.com/olzzio/y-statements-10eb07b5a177) stored in the
  [knowledger GitHub project](https://github.com/tmstff/knowledger)
  as [MarkDown](https://en.wikipedia.org/wiki/Markdown)
* and against [other decision record templates](https://github.com/joelparkerhenderson/architecture-decision-record) and
  other document storage options like [Atlassian Confluence](https://www.atlassian.com/software/confluence) or
  [Google Docs](https://docs.google.com/),
* to achieve well understandable decision documentation and that is stored close to the code and easy to edit for
  project contributors,
* accepting that we don't get WYSIWYG and will have a little bit more trouble creating diagrams.

## Template for Decision Records

```
# Decision Record: [title of the decision]

|                      |                                           |
|----------------------|-------------------------------------------|
| *Date*               | [YYYY-MM-DD]                              |
| *Status*             | [WIP | Decided | Deprecated | Superseded] |
| *Decision Author(s)* | [author name(s)]                          |

## Decision

* In the context of [context],  
* facing the need to [need],  
* we decided for [selected option]  
* and against [neglected options],
* to achieve [desired qualities],   
* accepting that [downsides].

## Details

## Links
```

## Links
* [Y-Statements](https://medium.com/olzzio/y-statements-10eb07b5a177)
* [knowledger GitHub project](https://github.com/tmstff/knowledger)
* [MarkDown](https://en.wikipedia.org/wiki/Markdown)
