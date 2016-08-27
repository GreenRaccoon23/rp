# rp
Go program to replace strings within files.  

Install:  
`go install github.com/GreenRaccoon23/rp`  
Clone:  
`go get github.com/GreenRaccoon23/rp`  

    [hiro@nakamura ~]$ rp -h
    rp <options> <file/directory>
      -o="": (old)
          string in file to replace
      -n="": (new)
          string to replace old string with
      -x="": (exclude)
          Patterns to exclude from matches, separated by commas
      -e=false: (expression)
          Treat '-o' and '-n' as regular expressions
      -r=false: (recursive)
          Edit matching files recursively [down to the bottom of the directory]
      -d=${PWD}: (directory)
          Directory under which to edit files recursively
      -s=1000: (semaphore-size)
          Max number of files to edit at the same time
          WARNING: Setting this too high will cause the program to crash,
          corrupting the files it was editing
      -c=false: (color)
          Colorize output
      -q=false: (quiet)
          Don't list edited files
      -Q=false: (Quiet)
          Don't show any output at all

Basically, it works like `sed -i "s/${old}/${new}/" "${file}"` but with a recursive option.  
I got tired of typing `find "${dir}" -type f -exec sed -i "s/${old}/${new}/" {}\;` because it's long and ugly and I had to deal with escaping regex for sed, so I made this program as a replacement.  

This program does allow regex as well, but it must be explicitly turned on by the `-e` flag. The regex syntax is slightly different than it is for `sed`; you can see the specifications [here](https://github.com/google/re2/wiki/Syntax) and test regular expressions [here](http://www.regexplanet.com/advanced/golang/index.html).  

The coolest thing about this program is that, when editing a directory recursively, it will edit multiple files at once. So whereas `find...sed...` might finish in a couple minutes, this program will finish in a couple seconds. The default limit is 1000 files at once, and it can be customized with the `-s` flag. **Don't set it above 1000**; if you set it too high, you'll cause the program to crash, which will probably corrupt a lot of your files. If you're editing huge files, set it to 100 (`-s 100`), or for massive files, use 10 (`-s 10`).
