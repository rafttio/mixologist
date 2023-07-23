# mixologist
Your personal bartender

## Building
Run `go build ./...` in you terminal from this directory, a `mixologist` binary will be created.

## Enabling auto completion
Execute the `completion` subcommand for your desired shell and write the output to a temporary file.

```
./mixologist completion zsh > /tmp/completion
```

Source the `completion` script file

```
source /tmp/completion
```

## Executing
Execute the binary as `./mixologist`
