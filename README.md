# osvdb

Go library and CLI tool for consuming OSV data, building a vulnerability database, and querying the database

## Why?

The data coming from the **Open Source Vulnerabilities (OSV)** community is some of the best vulnerability data in existence. It's usually made available as git repositories with thousands of files, or as an online database accessible through APIs.

This project aims to make it easier to consume OSV data locally and/or offline by providing a library and CLI tool for building and querying custom SQLite databases of OSV data.

## Installation

### Library

```shell
go get -u github.com/luhring/osvdb@latest
```

### CLI

```shell
go install github.com/luhring/osvdb/cmd/osvdb@main
```

## Examples

### Build a local database of GitHub Security Advisories

Clone GitHub's advisory-database repository to get the latest OSV data.

```shell
git clone --depth 1 https://github.com/github/advisory-database.git ghsa
```

Build the SQLite database!

```shell
osvdb build -R ghsa/advisories/github-reviewed -o ghsa.db
```

Query the database!

```shell
sqlite3 ghsa.db "SELECT * FROM vulnerabilities WHERE vulnerability_id = 'GHSA-2fr7-cc7p-p45q';"
```

## Ideas for the future...

- [ ] Smart updates to the database (only update what's changed)
- [ ] Namespaces for vulnerabilities to allow for multiple data sources in the same database
- [ ] Native support for consuming well-known OSV data sources
- [ ] Optimize the performance of database builds
- [ ] Standardize the database location on the local filesystem
- [ ] More... open an issue with your idea!
