# gomobile-ios-android-library-dev
Android and iOS cross library development with Golang

## How to set up?

Create base folders on your environment

```bash
cd ~
mkdir -p go/pkg go/bin go/src
```

### OSx

```bash
brew install go
```

### Linux and Windows with MinGW (or Unix Like Windows) from TarBall

```bash
wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.12.4.linux-amd64.tar.gz
```

### Dependencies and Golang configuration

> Edit your shell profile file configuration based on your distribution

```bash
# GO CONFIGURATION
export GOPATH="${HOME}/go"

## For OSx
export GOROOT="$(brew --prefix golang)/libexec"

## For linux and windows
export GOROOT="/usr/local/go"

## For legacy compile support C, Objective-C
export GOROOT_BOOTSTRAP=$GOROOT

## Append golang root and installed binaries throw go get
export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"
```

**Open a new terminal or reload your profile file**

To check go installation run following commands:

```bash
go version
go env
```

> Dependencies: gomobile, go-sqlite3, google-geolocate, ini.v1

Copy and paste those commands on terminal:
a
```bash
go get github.com/martinlindhe/google-geolocate
go get github.com/mattn/go-sqlite3
go get golang.org/x/mobile/cmd/gomobile
go get github.com/go-ini/ini
```

**Initialize gomobile in this project, with `gomobile init` command**

## Cloning

> Clone this repository as **googlemapsearch** name

```bash
git clone git@github.com:captaincode0/gomobile-ios-android-library-dev.git $GOPATH/src/googlemapsearch
```

**Note**: If you want to keep your files outside **GOPATH** directory, personally recommend to make a soft symlink in **GOPATH/src** folder, `ln -s <your-custom-path>/googlemapsearch $GOPATH/src`

## Set your google API Key

After Api Key generation for Google Maps, edit ini configuration file *config.ini* and add your own key

Check manual to generate a [Google Maps Api Key](https://developers.google.com/maps/documentation/embed/get-api-key)
