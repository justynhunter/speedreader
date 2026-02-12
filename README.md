# sread

A command line tool to speed read documents.

## Installing

If you have the go tooling installed simply do

```bash
go install github.com/justynhunter/sread
```

otherwise...

1. Download from the current release to the right 👉
2. Unzip the file and move the binary to somewhere in your $PATH (e.g. `mv sread /usr/local/bin`)
3. Use following the usage instructions below.

## Usage

```bash
cat your_doc | sread -w 200
```

or 

```bash
sread my_file.txt --no-highlight
```

| flag | short | default | purpose |
| --- | --- | --- | --- |
| --words-per-minute | -w | 300 | the speed the words are shown |
| --highlight-color | -c | #FF0088 | color of the 'center' character |
| --no-highlight | -n | false | don't highlight the 'center' character |
