#### Tool for mounting sshfs resources

##### Reason to create

I needed a way to use mounted sshfs resources for [ranger](https://github.com/ranger/ranger) file manager, like in mc.

##### Requirements

- golang
- sshfs tool

##### Instalation

```
git clone https://gitlab.com/handy-tools/sshmount.git
go build -o sshmount
```

##### Usage

For correct work you need to have almost one server record in `~/.ssh/config` 

Then execute 

`./sshmount <ssh server>` 

and by default it creates `~/.mount` directory and mounts sshfs to it.