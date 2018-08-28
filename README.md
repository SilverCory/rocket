# Rocket

[![GoDoc](https://godoc.org/github.com/astrocorp42/rocket?status.svg)](https://godoc.org/github.com/astrocorp42/rocket)
[![GitHub release](https://img.shields.io/github/release/astrocorp42/rocket.svg)](https://github.com/astrocorp42/rocket/releases/latest)
[![Build Status](https://travis-ci.org/astrocorp42/rocket.svg?branch=master)](https://travis-ci.org/astrocorp42/rocket)

Deploy software as fast and easily as possible

1. [Installation](#installation)
2. [Usage](#usage)
3. [Available providers](#available-providers)
4. [Roadmap](#roadmap)

-------------------


## Installation

### Using go (nightly)
```bash
$ go get -u github.com/astrocorp42/rocket
```

### Binary releases
[https://github.com/astrocorp42/rocket/releases/latest](https://github.com/astrocorp42/rocket/releases/latest)




## Usage

Go to your project's root directory then
```bash
$ rocket init # create a configuration .rocket.toml file with default configuration
# edit the file with the desired configuration, then
$ rocket # to deploy
```





## Available providers

| Provider              | Status |
| --------------------- | -------|
| [AWS S3](https://aws.amazon.com/s3) `s3` | :construction: |
| Custom script `script` | :heavy_check_mark: |
| [Google Firebase](https://firebase.google.com) `firebase` | :clock1: |
| [Google Cloud Storage](https://cloud.google.com/storage) `gcs` | :clock1: |
| [GitHub releases](https://help.github.com/categories/releases) `github_releases` | :heavy_check_mark: |
| [Heroku](https://www.heroku.com) `heroku` | :heavy_check_mark: |
| [Netlify](https://www.netlify.com) `netlify` | :clock1: |
| [NPM](https://www.npmjs.com) `npm` | :clock1: |



## Roadmap

See [https://github.com/astrocorp42/rocket/projects/2](https://github.com/astrocorp42/rocket/projects/2)
