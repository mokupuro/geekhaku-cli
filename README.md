# geekhaku-cli
技育博用にGoで作ったCLIツール

## Installation

- go

```
$ go get github.com/mokupuro/geekhaku-cli
```

- brew

```
$ brew tap mokupuro/geekhaku-cli
or
$ brew install mokupuro/geekhaku-cli/geekhaku-cli
```

## Update

```
$ brew upgrade mokupuro/geekhaku-cli/geekhaku-cli
```

## Usage

```
$ geekhaku-cli
Version: -
                                                                                              
 ________  _______   _______   ___  __            ___  ___  ________  ___  __    ___  ___     
|\   ____\|\  ___ \ |\  ___ \ |\  \|\  \         |\  \|\  \|\   __  \|\  \|\  \ |\  \|\  \    
\ \  \___|\ \   __/|\ \   __/|\ \  \/  /|_       \ \  \\\  \ \  \|\  \ \  \/  /|\ \  \\\  \   
 \ \  \  __\ \  \_|/_\ \  \_|/_\ \   ___  \       \ \   __  \ \   __  \ \   ___  \ \  \\\  \  
  \ \  \|\  \ \  \_|\ \ \  \_|\ \ \  \\ \  \       \ \  \ \  \ \  \ \  \ \  \\ \  \ \  \\\  \ 
   \ \_______\ \_______\ \_______\ \__\\ \__\       \ \__\ \__\ \__\ \__\ \__\\ \__\ \_______\
    \|_______|\|_______|\|_______|\|__| \|__|        \|__|\|__|\|__|\|__|\|__| \|__|\|_______|

Usage:
  geekhaku-cli [command]

Available Commands:
  cat         Print the ascii art of cat
  completion  Generate the autocompletion script for the specified shell
  doc         Print the ascii art of doc
  geekhaku    Print the ascii art of geekhaku
  help        Help about any command

Flags:
  -h, --help     help for geekhaku-cli
  -t, --toggle   Help message for toggle

Additional help topics:
  geekhaku-cli version    Print the version number of geekhaku-cli

Use "geekhaku-cli [command] --help" for more information about a command.
```

```
$ geekhaku-cli geekhaku
Version: -
Use the arrow keys to navigate: ↓ ↑ → ← 
? What format do you want to display?: 
  ▸ Default
    Kanji
    Shadow
    Moon
    3D
    
✔ Moon
You choose no.3 "Moon"

🌑🌑🌕🌖🌑🌑🌑🌒🌕🌗🌑🌑🌑🌑🌑🌑🌑🌑🌓🌕🌗🌑🌑🌑🌑🌑🌑🌒🌕🌗🌑🌑🌑🌒🌕🌘🌔🌘🌑
🌑🌑🌕🌖🌑🌑🌑🌒🌕🌗🌑🌑🌑🌒🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌘🌑🌓🌕🌗🌑🌑🌑🌓🌕🌕🌕🌕🌑
🌓🌕🌕🌕🌖🌔🌕🌕🌕🌕🌕🌕🌘🌒🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌘🌑🌓🌕🌗🌔🌕🌕🌕🌕🌕🌕🌕🌗
🌓🌕🌕🌕🌕🌑🌒🌕🌕🌖🌑🌒🌑🌑🌑🌑🌕🌕🌗🌑🌑🌕🌕🌖🌑🌑🌒🌕🌕🌕🌗🌕🌕🌕🌕🌕🌗🌕🌑
🌑🌒🌕🌖🌑🌑🌑🌒🌕🌖🌑🌑🌑🌑🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌕🌑🌓🌕🌕🌕🌗🌕🌕🌕🌕🌕🌕🌕🌘
🌑🌒🌕🌖🌑🌔🌕🌕🌕🌕🌕🌕🌘🌑🌑🌓🌕🌕🌕🌖🌑🌕🌕🌕🌗🌑🌑🌓🌕🌗🌑🌕🌕🌕🌕🌕🌕🌕🌑
🌓🌕🌕🌕🌕🌔🌕🌕🌕🌕🌕🌕🌑🌑🌑🌔🌕🌕🌕🌕🌕🌕🌕🌖🌑🌑🌑🌓🌕🌗🌑🌕🌕🌕🌕🌕🌕🌕🌑
🌓🌕🌕🌕🌖🌒🌕🌕🌑🌓🌕🌗🌑🌑🌑🌔🌕🌕🌕🌕🌕🌕🌕🌖🌑🌑🌑🌓🌕🌗🌑🌕🌕🌕🌕🌕🌕🌕🌑
🌑🌒🌕🌖🌑🌑🌓🌕🌕🌕🌕🌑🌑🌑🌑🌓🌕🌕🌕🌕🌕🌕🌕🌖🌑🌑🌑🌒🌕🌗🌔🌕🌕🌕🌕🌕🌕🌕🌖
🌑🌒🌕🌖🌑🌑🌑🌕🌕🌕🌗🌑🌑🌑🌑🌓🌕🌕🌕🌑🌑🌕🌕🌗🌑🌑🌑🌓🌕🌗🌔🌕🌕🌕🌕🌕🌕🌕🌖
🌑🌓🌕🌖🌑🌔🌕🌕🌕🌕🌕🌕🌑🌑🌑🌓🌕🌕🌕🌕🌕🌕🌕🌗🌑🌑🌑🌓🌕🌗🌑🌓🌕🌗🌑🌕🌕🌑🌑
🌒🌕🌕🌗🌒🌕🌕🌖🌑🌔🌕🌕🌗🌑🌑🌔🌕🌘🌑🌑🌕🌕🌕🌘🌑🌑🌑🌓🌕🌗🌑🌑🌓🌕🌕🌕🌖🌑🌑
🌑🌗🌑🌑🌑🌑🌑🌑🌑🌑🌑🌑🌑🌑🌑🌓🌕🌑🌑🌑🌕🌕🌖🌑🌑🌑🌑🌒🌕🌗🌑🌑🌑🌑🌕🌕🌘🌑🌑
```

## License

geekhaku-cli is released under the MIT license. See [LICENSE](https://github.com/mokupuro/geekhaku-cli/](https://github.com/mokupuro/geekhaku-cli/blob/main/LICENSE))
