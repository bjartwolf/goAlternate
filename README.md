```shell
go build alternate.go
./alternate --cpuprofile=alternate.prof
    Normal switch
    2000000000	         1.39 ns/op
    Fast switch
    2000000000	         0.89 ns/op
go tool pprof alternate alternate.prof
    Total: 478 samples
         291  60.9%  60.9%      291  60.9% main.AlternateBranch
         187  39.1% 100.0%      187  39.1% main.AlternateXor
           0   0.0% 100.0%      478 100.0% gosched0
           0   0.0% 100.0%      478 100.0% testing.(*B).launch
           0   0.0% 100.0%      478 100.0% testing.(*B).runN

(pprof) weblist main
```
