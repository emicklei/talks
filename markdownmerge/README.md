# markdown merge

Reads the entries of `slides.txt` (via -i flag) and writes a composed Markdown file `main.md` (via -o flag)

## install

    go install

## setup

create a folder `slides` and write a markdown file per slide.
create a file `slides.txt` with, per line, a local reference to each markdown file.
prefix the line with `!` to skip it.

in that folder, run:

    markdownmerge -i slides.txt -o ../main.md