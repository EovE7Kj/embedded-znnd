## linux-amd64 embedded znnd 
single push-button spinup of embedded `znnd` binary for linux-amd64 

### NOTE: For use in unikernel/container node development

- Run with default znn root directory (`/root/znn`):
```sh 
$ env CGO_ENABLED=0 go build -o znnd main.go
```

- Run with specified `<path>`: 
```sh 
$ awk -F'"' '/dataDir =/ { print $2 }' \
 | xargs -I{} sed -i "/dataDir/s/{}/<path>/g" main.go \
 && CGO_ENABLED=0 go build -o znnd main.go 
```
