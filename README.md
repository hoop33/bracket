# Bracket

This is a program for filling out a bracket for the NCAA tournament. It's written in Go, which I'm just starting to learn, so it's pretty rough.

## Installation

Make sure you have a working Go environment (go 1.1 is *required*). [See the install instructions](http://golang.org/doc/install.html).

To install `bracket.go`, run:

```
$ go get github.com/hoop33/bracket
```

Make sure your `PATH` includes to the `$GOPATH/bin` directory so your commands can be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

## Usage

From the command line, type:

```
$ bracket
```

The output plays the games, region by region (ignoring the play-in games), and takes you to the final four. At that point, it's up to you to decide who wins the final games.

I'm not done tweaking the formulas yet for determining the winners.

Note: This program is for fun only, and carries no guarantees. If you win big, good on you. If you lose big, tough luck!

## License

Bracket is released under the MIT license.