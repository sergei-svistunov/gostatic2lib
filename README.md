# gostatic2lib

## Description

`gostatic2lib` is a simple tool for convertation static files (HTML, JS, CSS, ...) to ONE go library that contains HTTP
handler that stores static files in execution file.

## Installation

`go get -u github.com/sergei-svistunov/gostatic2lib`

## Usage

`gostatic2lib -path <...> -package <...> -out <...>`

## How does it work

`gostatic2lib` walks through directory with static files and adds their encoded (gzip + base64) content to handler's map.
