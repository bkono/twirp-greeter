# twirp-greeter

Most polite lambda service using twirp, guaranteed.

## Umm, what?

Hello world, people. This is hello world, using twirp and testing deployment to lambda via [apex/up](https://github.com/apex/up). The good news? It Just Works™. Zero issues sending JSON or proto requests through to it. Mission accomplished.

## I wanna try!

Here's the quick and easy local-only approach:

```
go run main.go server -port 10101

... split terminal ...

go run main.go client -addr 'http://localhost:10101' -name 'you'

> 14:37:37 [client] resp received: greeting:"Hello, you!"

curl -H 'Content-Type:application/json' -X POST 'localhost:10101/twirp/sh.kono.examples.greeter.Greeter/Greet' -d '{"name": "john"}'

> {"greeting":"Hello, john!"}
```

Like I said, nice and easy. If you're setup already with [up](https://github.com/apex/up), use `up` to get it deployed, and replace the url (localhost:10101) with ``up url`` in the above commands, ie: 

```
go run main.go client -addr `up url` -name 'you'
```

For reference, the curl is obviously running with application/json, while the `go run main.go client` uses the generated proto client. Both worked interchangeably against the up deployed version.
