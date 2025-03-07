# Web Archives Manager (WAM)

Some organisations use a folder structure down a simple diectory tree, and serve this over HTTP/S as a way of publishing and organising files.

It's rudimentary artifacting, probably leftover from startup days. How about gaining a bit of control?

This tool allows publishing to a local folder ("prefix"), say `/var/www/releases/`, and using a release channel path to help resolving releases.

It also specifies a directory structure to help organise the existing tree.

## Consumer benefit

If using WAM as a client, download becomes simply:

```sh
# Download from a channel
TARFILE="$(wam get "http://files.lan/releases/ProjectAlpha" -c latest)"
# or for a specific version, `wam get "http://files.lan/releases/ProjectAlpha" -v 3.0.1`
tar xzf "$TARFILE"
```

If using generic tools like curl or wget, the consumer of the file repository can make firm assumptions on the strucutre of the release site:

```sh
PROJECT=ProjectAlpha
PROJECT_URL="http://files.lan/releases/$PROJECT"
VERSION="$(curl "$PROJECT_URL/chan/latest")"
FILE="$PROJECT-$VERSION.tar.gz"

curl "$PROJECT_URL/$VERSION/$FILE" -O "$FILE"
tar xzf "$FILE"

```

## Settings

```sh
# Select the deployment base target
wam prefix NAME

# Unselect a prefix - return command to state where no current prefix is set
# (prevent accidental publishing into wrong spaces)
wam prefix --none

# Set a prefix name and a prefix path
wam prefix PREFIX_NAME PREFIX_PATH

# Set a label under the current prefix. If no prefix is set, the command fails.
wam label LABEL
```

For most subsequent commands, a prefix must be currently active, else the commands fail. WAM does not attempt to auto-choose anything.

## Publish

An existing label under the current prefix must exist, else the action fails.

```sh
# Publish files as a Gzip tarball, optionally include a sidecar readme file, optionally renaming it to NAME
wam publish LABEL VERSION [-r README[:NAME]] -- FILES ...

# Specify the channels that should point to the given label/version
wam channel LABEL VERSION -- CHANNELS ...

# Delete a channel from a label
wam chan-del LABEL CHANNEL
```

## Cleanup

```sh
# Mark for retention
wam retain LABEL VERSION

# Remove mark for retention
wam unretain LABEL VERSION

# Remove all versions not marked for retention, or referred to by a channel
# Prompts user for prefix confirmation and each deletion, unless `-f` is specified
wam cleanup -y [-f] LABEL
```

Cleanup control:

If a folder contains a file `.no-cleanup`, then the cleanup process skips the folder entirely. This can be specified at any level.

## Query

```sh
# Show registered prefixes
wam prefix

# List all labels tracked in the system under the current prefix
wam label

# List the versions and channels of a specific label. Versions numerically sorted, descending.
wam list LABEL [-n N] { channels | versions }

# List files for given VERSION of LABEL
wam ls LABEL VERSION

# List all versions not marked for retention, or referred to by a channel
wam cleanup LABEL
```

## Retrieve

WAM can be used as a download client for a location published under this suite

```sh
# Performs download, and prints the downloaded file's name to stdout
# Requires curl or wget present on system
wam get PROJECT_URL {-c CHANNEL | -v VERSION} [-O OUTPUT_FILE]
```

## Tree structure

```
PREFIX/
  |
  +- LABEL/
      |
      +- chan/
          |
          +- channel files ...
          |
          +- v/
             |
             +- version dirs ...
```

Each channel file is a simple text file containing a version name, corresponding to a version dir

Each version dir contains a file with the LABEL name followed by the version.

So for a project "ProAlpha" with a stable channel pointing at version 1.0.0, and a latest pointing at 1.2 , the tree would have

```
$PREFIX/ProAlpha/chan/latest # a file containing "1.2"
$PREFIX/ProAlpha/chan/stable # a file containing "1.0.0"
$PREFIX/ProAlpha/v/1.0.0/ProAlpha-1.0.0.tar.gz
$PREFIX/ProAlpha/v/1.2/ProAlpha-1.2.tar.gz
```
