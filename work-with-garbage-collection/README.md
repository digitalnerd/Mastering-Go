# GC

---

```bash
$ time go run sliceGC.go

real    0m2.899s
user    0m2.597s
sys     0m2.608s
```

```bash
$ time go run mapStar.go

real    0m25.561s
user    0m35.801s
sys     0m0.796s
```

```bash
$ time go run mapNoStar.go

real    0m13.487s
user    0m13.217s
sys     0m0.468s
```

```bash
$ time go run mapSplit.go

real    0m13.200s
user    0m12.961s
sys     0m0.451s

```
