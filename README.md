# Go markdown daemon (gmdd)

Gmdd is a daemon for serving and highlighting markdown files with an
opinionated set of features. It compiles into a single executable that
encompasses all assets and configuration.

Features
--------

#### Math expressions
Math expressions are supported in TeX format and are processed with the
[KaTeX][katex] library. (See the [KaTeX wiki][katex-supported] for the list of
supported functions.)

Math expressions must be in fenced or inline code blocks. Fenced code blocks
must use the `math` language identifier. Inline code blocks must start and
end in `$` delimiters.

#### Diagrams
Diagrams are supported in the mermaid format and are processed with the
[mermaid][mermaid] library. (See the [mermaid documentation][mermaid] for the
list of supported diagram types and syntax.)

Diagrams must be in fenced code blocks and use the `mermaid` language
identifier.

#### Code snippets
Codes snippets are highlighted with the [highlight.js][hljs] library.

Code snippers must be in fenced code blocks and use their respective language
identifier.


[hljs]: https://highlightjs.org/
[katex]: https://khan.github.io/KaTeX/
[katex-supported]: https://github.com/Khan/KaTeX/wiki/Function-Support-in-KaTeX
[mermaid]: https://mermaidjs.github.io/


Usage
-----

#### Build
```sh
make
```

#### Run
```sh
./gmdd [${ROOT_DIR}]
```