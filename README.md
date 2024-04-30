# Multiselect Form Control with htmx

A proof-of-concept for multiselect form control using [htmx](https://htmx.org) and [bootstrap](https://getbootstrap.com). No single line of JS. The back-end is written in [Go](https://go.dev).

The data is taken from [Random User Generator](https://randomuser.me) and stored locally for the sake of demonstration (consistent user names).

![Demo](https://github.com/apleshkov/htmx-multiselect-control/blob/main/demo.gif)

## Features

* Multiselect
* Live search (excluding already selected user names)
* Removing of selected user names
* Loading indicator
* No JS

## Run

1. `go run main.go`
2. Open <http://localhost:1234>

## License

MIT
