# Install

```bash
go install github.com/GreenRaccoon23/rp;
```

# Download

```bash
git clone https://github.com/GreenRaccoon23/rp.git;
```

# Description

Program to replace strings within files.

It supports regex patterns.

It supports globbing patterns (in order to edit multiple files at once).

It supports concurrency so that multi-file edits are insanely fast (i.e., 1000+ files in half a second).

It is like `sed -i "s/${old}/${new}/" "${file}"` but supports multiple lines, supports multiple files, supports non-greedy regex, requires less "kill me now" syntax, and runs faster.

```bash
[hiro@nakamura ~]$ rp -h
rp <options> <path>...
  -o, --old string        old string/pattern to find
  -n, --new string        new string/pattern to replace old one with
  -e, --regex             Treat '-o' and '-n' as regular expressions
  -r, --recursive         Edit files under each <path>
  -i, --include string    File patterns to include, separated by commas
  -x, --exclude string    File patterns to exclude, separated by commas
  -c, --concurrency int   Max number of files to edit simultaneously (default 1)
  -l, --list              List which files would be edited but do not edit them
  -v, --verbose           Show more output
  -q, --quiet             Hide all output except errors

WARNING: Setting concurrency too high will cause the program to crash,
corrupting the files it was editing.

The syntax of the regular expressions accepted is the same general
syntax used by Perl, Python, and other languages. More precisely, it
is the syntax accepted by RE2 and described at
https://golang.org/s/re2syntax, except for \C.
For an overview of the syntax, run:
	go doc regexp/syntax
```

Written in Go/Golang.

# Why

Look at this:

```bash
find . -type f -name="*.svg" -exec sed -i 's/\(fill="#\).*\("\)/\1ff0000\2/g' {}\; ;
```

Kill me now. I needed to read a manual every time I needed to write something like this. Plus, since `sed` does not support the non-greedy `.*?`, this command with the greedy `.*` would erase most of the file content. This also most likely would not work on MacOS's version of `sed`.

I wanted something less rocket science like this:

```bash
rp -re -o '(fill="#).*?(")' -n '${1}ff0000${2}' -i '*.svg' .;
```

So I made it happen.

# Regex

The regex syntax used by this program is similar to what most programs use, so it is different than the one used by `sed` (thankfully). The spec is here:
- [https://github.com/google/re2/wiki/Syntax](https://github.com/google/re2/wiki/Syntax)

Test regex patterns here:
- [http://www.regexplanet.com/advanced/golang/index.html](http://www.regexplanet.com/advanced/golang/index.html).

# Warning

The `-c` option (concurrency) is powerful. Do not set it too high. It can cause the program to edit 1000 files in half a second, but it can also cause the program to erase 9000 files in half a second. For a computer with 4 GB RAM editing average sized files, 1000 is safe. For huge files (> 500 KB), 100 is safer. For massive files (> 1 MB), stick to 1 or 2.
