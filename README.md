# ev1write
A golang program to write to a magic ntag 21x, specifically while its emulating a ev1 48b

## Build
```/bin/bash
go mod tidy
go build
```

### Usage
```/bin/bash
ev1write ~/Nextcloud/proxmark/condo5.json
```

### Scripts
I wrote some scripts to aid me in remembering the commands to use for dumping, cloning, or just restoring from a json file, see these in the scripts dir
