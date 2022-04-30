# Error-Tags

`et`, short for `error-tags`, is a small library for tagging errors, and inspecting the tag of other errors. `et` operates well with the `pkg/errors` packages for providing convenient, ergonomic, and robust error handling.

`error`s can easily be tagged using `et.Tag`.
```go
const (
  FooFailure et.ErrorTag = iota
)

func DoFooFallibleThing() error {
  err := foo.OperationThatMayFail()
  return et.Tag(err, FooFailure) // if err is nil, Tag will return nil as well
}
```

Any `error` can easily be checked for a given tag.

```go
func DoABunchOfFallibleThings() error {
  if err := DoFooFallibleThing(); err != nil {
    return err
  }
  return DoSomeOtherFallibleThing()
}

// ...
e := DoABunchOfFallibleThings()
if et.Tagged(e, FooFailure) {
  fmt.Println("DoFooFallibleThing failed")
}
```

Tags are even preserved through wrapped errors.
```go
e := errors.Wrap(
  et.Tag(
    errors.New("an untagged error"),
    SomeTag,
  ),
  "..."
)

fmt.Println(et.Tagged(e, SomeTag))
// true
```

You can also extract the tag directly using `errors.As`.
```go
const (
  BadRequest et.ErrorTag = iota
  NotFound
  Unauthorized
)

func ErrorToHTTPStatus(e error) int {
  var clientErr et.TaggedError
  
  if errors.As(e, &clientErr) {
    switch clientErr.Tag {
    case BadRequest:
      return 400
    case NotFound:
      return 404
    case Unauthorized:
      return 401
    default:
      return 500
    }
  }

  return 500
}
```