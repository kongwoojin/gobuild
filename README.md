## gobuild
gobuild is cli tool for build Go program

### Install
#### Windows
* Put gobuild binary for Windows in C:\Windows
* Rename gobuild*.exe to gobuild.exe

#### Linux
* Put gobuild binary for Linux in /usr/local/bin
* Rename gobuild* to gobuild

### Options
* -a: Set architecture, default is local machine's architecture
* -p: Set platform, default is local machine's platform
* -o: Set output, default is build/{name}_{platform}-{architecture}
* -help: Show help

### Usage
* Build package for same architecture and platform with local machine
```bash
gobuild
```

* Build package for linux/arm64
```bash
gobuild -a arm64 -p linux
```

* Build package and saved it at out/
```bash
gobuild -o out/
```

* Build package and saved it as hello
```bash
gobuild -o build/hello
```