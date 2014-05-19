# Gourd #

Gourd is a simple automation tool that will watch for
specific file changes in the local directory.

When a matching file is changed the commands specified in
your Pumpkin will be executed.

Gourd uses fsnotify and as such should be able to work on
any of the major platforms such as Windows, Linux and Mac.

## Example Usage ##

Here is an example of a Pumpkin
```json
{
  "pattern": ".*.go",
  "commands": [
    "go test"
  ]
}
```

## Why are the files called Pumpkin? ##

Because pumpkins are a type of gourd.
