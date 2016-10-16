.PHONY: go all build clean vendor package-darwin package-linux package-windows release

all: build

build:
	go build

darwin: *.go
	GOOS=darwin GOARCH=amd64 go build

linux: *.go
	GOOS=linux GOARCH=amd64 go build

windows: *.go
	GOOS=windows GOARCH=386 go build

clean:
	rm -f get_iplayer_rss

vendor:
	godep save

package-darwin: darwin
	tar zcf get_iplayer_rss-darwin-amd64.tar.gz get_iplayer_rss

package-linux: linux
	tar zcf get_iplayer_rss-linux-amd64.tar.gz get_iplayer_rss

package-windows: windows
	zip get_iplayer_rss-win32-i386.zip -xi get_iplayer_rss.exe

release: package-linux package-darwin package-windows