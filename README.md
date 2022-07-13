# good handlers

An experiment in having good (HTTP) handlers. Not finished, please be chill, I just hope it's interesting.

## Waffle

In my view, to have good separation of concerns and other computer words, HTTP handlers should be only responsible for a small set of things

1. Parse an incoming request
2. Call a service layer
3. Depending on the result of the service call, send a response

I wrote about this extensively in [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited)

## Enforcement

In the projects I work on, we strive to do this, but it's a somewhat manual thing right? There's nothing enforcing this style other than our own judgement and discipline. 

Web frameworks, are nasty and we all hate them, but can I offer something that is opinionated _enough_ to:
- enforce this separation of concerns
- retain testability
- be of the Go standard library. Embrace `net/http/Handler` so a user doesn't have to rewrite everything, and loses the advantage of standard middlewares, routers, etc
- not be annoying

## What we have

"good handlers"

A function that enforces this simple design constraint in a type-safe but flexible manner (thanks generics!), and returns you a `http.Handler` to use as normal.

Dive in to the test `goodhandlers_test.go` and `cmd/example/main.go`

### What is a service?

To create a "good handler" you need to supply a `Service` which takes in an `A` (and a `context.Context`) and returns a `B, error`. 

In my book, that is a good definition of a service function/method. In addition, you need to supply a `Decoder[A]` so that we can parse the incoming request into an `A` and an `Encoder[B]`, to turn the result from the service call into a HTTP response.

Clearly this won't work for every use-case, but I intend to keep trying out different scenarios and see what I come up with. 

## What about everything else?

Logging? Metrics? Blah blah, it's all handled already by everyone else. Mostly these x-cutting concerns should be handled by middleware, and because good-handlers just implements `http.Handler` you can compose your middleware like normal.