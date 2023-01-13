echo $1 > file.c

gcc -nostartfiles -g file.c -o file

./file
