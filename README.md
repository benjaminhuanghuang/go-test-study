# Go Test Study



## Command
```
  go test 

  go test -v   # display details

  go test -bench=.
```


## Naming convention
File:   xxx_test.go

Function: func TestXXX(t *testing.T)

Benchark Function:  BenchmarkXXX

## Control the testing order
Using sub test
```
  func TestAll(t *testing.T){
    t.Run("test1", testPrint)
    t.Run("test2", testPrint2)
  }
```

## Skip test
```
  t.SkipNow()
```

## Main Test
```
  func TestMain(m *testing.M){
    //Init test suit

    m.Run()
  }
```