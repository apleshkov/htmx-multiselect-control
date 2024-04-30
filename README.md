# Multiselect Form Control with htmx

A proof-of-concept for multiselect form control using (https://htmx.org)[htmx] and (https://getbootstrap.com)[bootstrap]. No single line of JS. The back-end is written in (https://go.dev)[Go].

The data is taken from (https://randomuser.me)[Random User Generator] and stored locally for the sake of demonstration (consistent user names).

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
