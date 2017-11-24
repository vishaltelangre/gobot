Gobot
======

**tl;dr**

- I am your pet, baby! I can do anything for you, just tell me...
- Sneak Peek Preview:

![Gobot Screenshot](https://raw.github.com/vishaltelangre/gobot/master/gobot_preview.png)

## Installation

- Before getting started, make sure that you have [installed Go](http://golang.org/doc/install) and have set workspace (`$GOPATH`, etc.), or [RTFM](http://golang.org/doc/code.html) yourself how to do it!
- Also, [Readline](http://en.wikipedia.org/wiki/GNU_Readline) is a dependency of this project, for which, I needed to install `libedit-dev` package on my Ubuntu system.
- Currently Gobot has support for Unix-based systems only.

Download (or update) and build Gobot and its dependencies:

```sh
$ go get -u github.com/vishaltelangre/gobot
```

Test whether `gobot` command works fine by checking version of it:

```
$ gobot --version
```

## Usage

There are two modes of gobot, viz. an **interactive mode** and **command line mode (default)**.

### Interactive Mode
  - This starts a REPL-like interactive session, where you can feed commands to gobot and it will interactively respond to you.
  - To get in this mode, just type either

    ```
    gobot -i
    ```

  - Also you can name your bot with other name (default is _Gobo_) by starting gobot interactive session using command:

    ```
    gobot -i --name=Darling
    ```

  - To list all available/supported commands, type in `help me` or `?` anytime when in interactive mode.

### Command Line Mode (Default)
  - This mode works like rest of other commands work on your system.
  - Usage examples could be:

    ```
    gobot what is my ip
    ```

    or

    ```
    gobot what does mean by a Bot
    ```

## Contributing

  Hey, you're encouraged to contribute to this project.

  * Fork this project. Add extension for stuff you like.
  * Report bugs, comment on and close open issues.
  * Make a pull request!

## Copyright and License

Copyright (c) 2014, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).