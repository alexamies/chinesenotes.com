module e2etest

go 1.13

require (
	github.com/alexamies/cnweb v0.0.1
	github.com/gorilla/mux v1.7.3
)

replace github.com/alexamies/cnweb => ./tmp/cnweb
