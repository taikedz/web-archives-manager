# Web Archives Manager (WAM)

Some organisations use a folder structure down a simple diectory tree, and serve this over HTTP/S as a way of publishing and organising files.

It's rudimentary artifacting, probably leftover from startup days. How about gaining a bit of control?

This tool allows publishing to a local folder ("prefix"), say `/mnt/artifacts/releases/`, and using a release channel path to help resolving releases.

It also specifies a directory structure to help organise the existing tree.

## Consumer benefit

If using WAM as a client, download becomes simply:

```sh
# Download from a channel
TARFILE="$(wam get "http://files.lan/releases/ProjectAlpha" latest)"
# or for a specific version, `wam get "http://files.lan/releases/ProjectAlpha" v/3.0.1`
tar xzf "$TARFILE"
```

If using generic tools like curl or wget, the consumer of the file repository can make firm assumptions on the strucutre of the release site:

```sh
PROJECT=ProjectAlpha
PROJECT_URL="http://files.lan/releases/$PROJECT"
VERSION="$(curl "$PROJECT_URL/chan/latest")"
FILE="$PROJECT-$VERSION.tar.gz"

curl "$PROJECT_URL/v/$VERSION/$FILE" -O "$FILE"
tar xzf "$FILE"

```

## Settings

```sh
# Select the deployment base target
wam prefix NAME

# Unselect a prefix - return command to state where no current prefix is set
# (prevent accidental publishing into wrong spaces)
wam prefix --unset

# Set a prefix name and a prefix path
wam prefix PREFIX_NAME PREFIX_PATH

# Set a project name under the current prefix. If no prefix is set, the command fails.
wam project PROJECT
```

For most subsequent commands, a prefix must be currently active, else the commands fail. WAM does not attempt to auto-choose anything.

## Publish

An existing project under the current prefix must exist, else the action fails; this avoids publishing to a mistype project name.

```sh
# Publish files as a Gzip tarball, optionally include a sidecar readme file, optionally renaming it to NAME
wam publish PROJECT VERSION [-r README[:NAME]] -- FILES ...

# A channel is a name that points to a specific version. Typically "latest" or "stable" are names to expect. Channels can update over time.

# Specify the channels that should point to the given project/version
wam channel PROJECT VERSION -- CHANNELS ...

# Delete a channel from a project
wam chan-del PROJECT CHANNEL
```

## Cleanup

```sh
# Mark for retention
wam retain PROJECT VERSION

# Remove mark for retention
wam unretain PROJECT VERSION

# Remove all versions older than N days if not marked for retention, nor referred to by a channel
# Prompts user for prefix confirmation and each deletion, unless `-f` is specified
# `-y` performs the cleanup, else the items that would be removed are merely printed.
wam cleanup -y -d N [-f] PROJECT
```

Cleanup control:

If a version folder contains a file `.no-cleanup`, then the cleanup process skips the folder entirely.

## Query

```sh
# Show registered prefixes
wam prefix

# List all projects tracked in the system under the current prefix
wam project

# List the versions and channels of a specific project.
# Versions numerically sorted, ascending.
# Channels alphabetically sorted and printed with their corresponding versions.
#  Use  `-r` to reverse-sort
wam list PROJECT [-r] { channels | versions }

# List files for given VERSION of PROJECT
wam ls PROJECT VERSION

# List all versions not marked for retention, or referred to by a channel
wam cleanup PROJECT
```

## Retrieve

WAM can be used as a download client for a location published under this suite

```sh
# Performs download, and prints the downloaded file's name to stdout
# Requires curl or wget present on system
wam get PROJECT_URL {CHANNEL | v/VERSION} [-O OUTPUT_FILE]
```

## Tree structure

```
PREFIX/
  |
  +- PROJECT/
      |
      +- chan/
      |   |
      |   +- channel files ...
      |
      +- v/
         |
         +- version dirs ...
           |
           +- PROJECT-VERSION.tar.gz
           |
           +- README.txt
```

Each channel file is a simple text file containing a version name, corresponding to a version dir

Each version dir contains a file with the PROJECT name followed by the version.

So for a project "ProAlpha" with a stable channel pointing at version 1.0.0, and a latest pointing at 1.2 , the tree would have

```
$PREFIX/ProAlpha/chan/latest # a file containing "1.2"
$PREFIX/ProAlpha/chan/stable # a file containing "1.0.0"
$PREFIX/ProAlpha/v/1.0.0/ProAlpha-1.0.0.tar.gz
$PREFIX/ProAlpha/v/1.2/ProAlpha-1.2.tar.gz
```

## Example session

First setup

```sh
# Register a prefix to publish into
wam prefix artifacts /mnt/artif-server/www
wam project CoolProg
```

Publish a new release of a given project

```sh
# Ensure we are using the `artifacts` prefix
wam prefix artifacts

# Publish the contents of the bin/ directory as version 1.0.1 of the CoolProg project
# Include a readme.txt from the docs folder, and publish it as index.txt
wam publish CoolProg 1.0.1 -r docs/readme.txt:index.txt -- bin/*

# Update the `latest` and `1.0` labels to point to the new version
wam channels CoolProg 1.0.1 -- latest 1.0

# If there are multiple prefixes possible, prevent the next user from
#   accidentally pushing to wrong place
wam prefix --unset
```

