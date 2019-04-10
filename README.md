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

```bash
./sshmount <ssh server>
```

and by default it creates `~/.mount` directory and mounts sshfs volume to it.

**!** You can put the binary to `~/.bin` and set `$PATH` env variable with

```bash
echo "export PATH=$PATH:$HOME/.bin" >> ~/.bashrc
```
or to `~/.bash_profile`

if you have mac.

---

##### Todo

- [ ] Add make file for building
- [ ] Add autocompletion file check
- [ ] Add clean support of OSX