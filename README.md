# speedreader

### Usage

```bash
cat your_doc | sread -d 400
```

or 

```bash
sread my_file.txt --no-highlight
```

| flag | short | default | purpose |
| --- | --- | --- | --- |
| --delay | -d | 300 | delay between words in ms |
| --no-highlight | -n | false | don't highlight the 'center' character |

### TODO
- [x] center text on screen
- ~~[ ] style text~~
- [x] add color highlight for focus
- [x] support for file name, so you don't need to cat
- [x] show help when nothing is piped in
- [ ] add xdg config for default --delay
