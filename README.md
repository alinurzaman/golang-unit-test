## HOW TO RUN TEST

- `go test` (run all unit test)
- `go test -v` (run all unit test with details)
- `go test -v -run=TestFunctionName` (run specific unit test)
- `go test ./...` (run from parent folder)

## HOW TO HANDLE FAILED TEST

- `t.Fail()` (fail the unit test but continue to run the rest code)
- `t.FailNow()` (fail the unit test and stop run the test)
- `t.Error()` (show the error log then call t.Fail())
- `t.Fatal()` (show the error log then call t.FailNow())

## HOW TO USE ASSERTION

- Add this library: github.com/stretchr/testify
- Run script `go get github.com/stretchr/testify`

## HOW TO SKIP TEST

- `t.Skip()` (to skip the unit test)

## BEFORE AND AFTER TEST

- Use `func TestMain()`
- Use `testing.M` instead of `testing.T`
- `m.Run()` (to run all unit test)

## HOW TO IMPLEMENT SUB TEST

- `t.Run("Sub Test Name", func(t *testing.T){})`
- `go test -run TestFunctionName/SubTestName` (run only specific sub test in a function)
- `go test -run /SubTestName` (run all sub test with prefix)

## WHY USE TABLE TEST

- To simplify when we use many data.
- Use struct to implement table test.

## HOW TO USE BENCHMARK TEST

- Use `testing.B`
- Don't return value.
- `go test -v -bench=.` (run all benchmark and unit test)
- `go test -v -run=NotMatchUnitTest -bench=.` (run benchmark only)
- `go test -v -run=NotMatchUnitTest -bench=BenchmarkName` (run specific benchmark)
- `go test -v -run=NotMatchUnitTest -bench=. ./...` (run benchmark from parent folder)
