# dywoq/gh-issue

gh-issue is a tool written in Go, allowing you to mange GitHub repository issues.

## Usage
To use it, you need to install it:
```
go install github.com/dywoq/gh-issue@latest
```

For example, if you want to output all issues into the console, you can use this command:
```
gh-issue get <issues-ids> <owner> <repository> <token>
```

The output would look like this:
![The example output](image.png)

## Documentation
To see the documentation of commands, there's [`./docs/commands.md`](./docs/commands.md) file.
