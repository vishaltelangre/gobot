Gobot
======

I am your pet, baby! I can do anything for you, just tell me...

## Installation

Before getting started, make sure that you must have setup your go workspace, or [rtfm](http://golang.org/doc/code.html) yourself!

```sh
go get -u github.com/vishaltelangre/gobot
which gobot
```

## Usage!

- There are two modes of gobot, viz. an **interactive mode (default)** and **command line mode**.

### Interactive Mode (Default)
  - This starts a REPL-like interactive session, where you can feed commands to gobot and it will interactively respond to you.
  - To get in this mode, just type either

    ```
    gobot
    ```

    or

    ```
    gobot -i
    ```

  - Also you can name your bot with other name (default is _gobo_) by starting gobot interactive session using command:

    ```
    gobot --name=Darling
    ```

  - To list all available/supported commands, type in `help me` or `?` anytime when in interactive mode.
  - Below is sneak peek preview of the interactive mode:

    -![Gobot Screenshot](https://raw.github.com/vishaltelangre/gobot/master/gobot_preview.png)

### Command Line Mode
  - This mode works like rest of other commands work on your system.
  - Usage examples could be:

    ```
    gobot -c what is my ip
    ```

    or

    ```
    gobot -c what does mean by a Bot
    ```

#### You're encourage to contribute to this project!

## Copyright and License

Copyright (c) 2014, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).