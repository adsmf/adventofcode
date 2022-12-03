# Benchmarks
The following are the benchmarks for the Go implementations of solutions for each day. The results are as measured by a `BenchmarkMain` benchmark in each solution.

## CPU time

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 88.7µs | 102µs | 46.4µs | 19ms | - | 396µs | 14.4µs | 75.1µs
Day 2 | 390µs | 176µs | 47.8µs | 76.4ms | - | 3.06ms | 1.22µs | 4.78µs
Day 3 | 1.53ms | 507µs | 2ns | 45.7ms | - | 83.5µs | 48.2µs | **🔴 976µs**
Day 4 | 2.24s | 23ms | 1.3ms | 745µs | - | 927µs | 692µs | -
Day 5 | 2.1ms | 15.8s | 70.5ms | 112ms | - | 28µs | 704µs | -
Day 6 | 5.73s | 183µs | 11.4ms | 37.5ms | - | 1.34ms | 656ns | -
Day 7 | 22.9ms | 5.81ms | 2.4ms | 3.2ms | - | 1.25ms | 57.1µs | -
Day 8 | 182µs | 131µs | 2.42ms | 681µs | - | 1.81ms | 535µs | -
Day 9 | 19.6ms | 5.1ms | 82.8µs | 650ms | - | 819µs | 244µs | -
Day 10 | 176ms | 4.04ms | 1.54ms | 348ms | - | 117µs | 74.4µs | -
Day 11 | 38ms | **🔴 42.5s** | 462µs | 1.78s | - | 20.7ms | 360µs | -
Day 12 | 1.82ms | 537ms | 18.4ms | - | - | 61.8µs | 5.09ms | -
Day 13 | 270ms | 370µs | **🔴 1.15s** | 6.02ms | - | 24.7µs | 145µs | -
Day 14 | 3.54ms | **🔴 16.8s** | 524ms | - | - | 6.64ms | 405µs | -
Day 15 | 9.97s | 27.6µs | 440ms | - | - | **🔴 620ms** | **🔴 296ms** | -
Day 16 | 1.7ms | 115ms | 9.32ms | - | - | 8.75ms | 45.2µs | -
Day 17 | 63.6ms | 33.9ms | 566ms | 38.5ms | - | 25.7ms | 301µs | -
Day 18 | 680ms | 228ms | 17.7ms | 412ms | - | 2.89ms | 18.1ms | -
Day 19 | 4.39s | 205ms | 2.95ms | 246µs | - | 15.8ms | 59ms | -
Day 20 | 5.1s | 452µs | 25.6ms | 2.02ms | - | 224ms | 18.9ms | -
Day 21 | 29.8µs | 335ms | 562ms | **🔴 44.5s** | - | 705µs | 7.7ms | -
Day 22 | 1.36s | 2.65s | 654ms | 1.71s | - | 288ms | 16.7ms | -
Day 23 | 81.4µs | 5.53ms | 4.65ms | 8.02ms | - | **🔴 682ms** | **🔴 142ms** | -
Day 24 | **🔴 31.9s** | 211ms | 485ms | 140ms | - | 101ms | 1.42ms | -
Day 25 | 69.9ms | 432ms | 561ms | 31.7ms | - | 179µs | 139ms | -
*Total* | *1m2s* | *1m19.9s* | *5.1s* | *50s* | *0s* | *2.01s* | *707ms* | *1.06ms*

## Heap memory

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 17.0 KB | 97.4 KB | 4.9 KB | 6.8 MB | - | 15.0 KB | - | **🔴 15.6 KB**
Day 2 | 372 KB | 12.5 KB | 17.6 KB | 18.9 MB | - | 545 KB | - | 256 B
Day 3 | 513 KB | 659 KB | - | 57.3 MB | - | 55.0 KB | 24.6 KB | **🔴 25.3 KB**
Day 4 | 454 MB | 10.9 MB | 504 KB | 360 KB | - | 488 KB | 391 KB | -
Day 5 | 1.5 MB | 3.8 GB | 66.7 KB | 129 MB | - | 9.8 KB | 45.9 KB | -
Day 6 | 218 MB | 31.1 KB | 1.7 MB | 14.3 KB | - | 535 KB | - | -
Day 7 | 9.9 MB | 3.3 MB | 922 KB | 170 KB | - | 710 KB | 61.5 KB | -
Day 8 | 64.4 KB | 163 KB | 238 KB | 1.2 MB | - | 798 KB | 116 KB | -
Day 9 | 10.6 MB | 163 KB | 22.1 KB | 383 MB | - | 218 KB | 273 KB | -
Day 10 | 185 MB | 5.2 MB | 6.0 KB | **🔴 1.6 GB** | - | 27.0 KB | 14.3 KB | -
Day 11 | 7.9 MB | **🔴 23.0 GB** | 442 KB | 1.6 MB | - | 1.6 MB | 1.1 KB | -
Day 12 | 643 KB | 12.0 KB | 2.8 MB | - | - | 20.0 KB | 6.5 MB | -
Day 13 | 107 MB | 183 KB | 8.0 KB | 2.9 MB | - | 6.0 KB | 26.6 KB | -
Day 14 | 1.1 MB | 2.7 GB | 27.3 MB | - | - | 2.7 MB | 77.9 KB | -
Day 15 | **🔴 3.3 GB** | 6.5 KB | - | - | - | 240 MB | **🔴 103 MB** | -
Day 16 | 825 KB | 260 MB | 3.9 MB | - | - | 1.3 MB | 29.0 KB | -
Day 17 | 2.5 KB | 24.3 MB | 32.3 KB | 9.5 MB | - | 9.0 MB | 192 B | -
Day 18 | 280 MB | 44.8 MB | 14.1 MB | 168 MB | - | 1.4 MB | 1.9 MB | -
Day 19 | 411 MB | 48.2 MB | 2.0 MB | 13.7 KB | - | 7.2 MB | 3.5 MB | -
Day 20 | **🔴 2.6 GB** | 343 KB | 21.1 MB | 1.0 MB | - | **🔴 1.1 GB** | **🔴 67.1 MB** | -
Day 21 | - | 191 MB | 283 MB | 463 KB | - | 372 KB | 2.0 MB | -
Day 22 | 483 MB | 5.3 GB | 6.2 MB | 329 MB | - | 110 MB | 868 KB | -
Day 23 | 13.5 KB | 3.9 MB | 2.0 MB | 9.1 MB | - | 8.0 MB | 16.1 MB | -
Day 24 | 19.4 MB | 35.9 MB | **🔴 1.0 GB** | 188 MB | - | 58.4 MB | 105 KB | -
Day 25 | - | 290 MB | 430 KB | 1.4 MB | - | 166 KB | 21.4 MB | -
*Total* | *8.0 GB* | *35.8 GB* | *1.4 GB* | *2.9 GB* | *0.0 B* | *1.6 GB* | *224 MB* | *41.1 KB*