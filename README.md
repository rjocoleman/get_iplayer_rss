# get_iplayer_rss

This is a cross platform CLI (Linux, Windows and macOS) to create RSS feeds from [get_iplayer](https://github.com/get-iplayer/get_iplayer).


## `gen` command

Generates RSS - parses your `get_iplayer` `download_history` file and creates an iTunes RSS feed per show.

```
 $ get_iplayer_rss gen --help
Generate an RSS feed from a get_iplayer download_history file

Usage:
  get_iplayer_rss gen [flags]

Flags:
  -d, --directory string     Path to get_iplayer directory (default "/etc/get_iplayer")
  -o, --output-path string   RSS file output path e.g. /var/www
  -u, --url string           URL to webroot e.g. https://example.com/path/to/dir
```