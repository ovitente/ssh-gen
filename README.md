#### Tool for mounting sshfs resources

##### Reason to create

I needed a way to use mounted sshfs resources for [ranger](https://github.com/ranger/ranger) file manager, like in mc.

##### Requirements

- Installed golang to build source code
- Installed sshfs tool

##### Instalation

`go bouild -o sshmount`

##### Usage

For correct work you need to have almost one server record in `~/.ssh/config` 

Then execute `./sshmount <ssh server>` and by default it will create `~/.mount` directory and mount sshfs to it.