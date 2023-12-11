# Benchmarks
The following are the benchmarks for the Go implementations of solutions for each day. The results are as measured by a `BenchmarkMain` benchmark in each solution.

## CPU time

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022 | 2023
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 88.7µs | 102µs | 46.4µs | 19ms | - | 396µs | 14.4µs | 13.9µs | 579µs
Day 2 | 390µs | 176µs | 47.8µs | 76.4ms | - | 3.06ms | 1.22µs | 4.18µs | 22.2µs
Day 3 | 1.53ms | 507µs | 2ns | 45.7ms | - | 83.5µs | 48.2µs | 65.7µs | **🔴 2.78ms**
Day 4 | 2.24s | 23ms | 1.3ms | 745µs | - | 927µs | 692µs | 23.4µs | 62.8µs
Day 5 | 2.1ms | 15.8s | 70.5ms | 112ms | - | 28µs | 704µs | 15.7µs | 35.2µs
Day 6 | 5.73s | 183µs | 11.4ms | 37.5ms | - | 1.34ms | 656ns | 14.5µs | **🔴 3.44ms**
Day 7 | 22.9ms | 5.81ms | 2.4ms | 3.2ms | - | 1.25ms | 57.1µs | 13.2µs | 312µs
Day 8 | 182µs | 131µs | 2.42ms | 681µs | - | 1.81ms | 535µs | 458µs | 1.68ms
Day 9 | 19.6ms | 5.1ms | 82.8µs | 650ms | - | 819µs | 244µs | 446µs | 114µs
Day 10 | 176ms | 4.04ms | 1.54ms | 348ms | - | 117µs | 74.4µs | 919ns | **🔴 3.66ms**
Day 11 | 38ms | **🔴 42.5s** | 462µs | 1.78s | - | 20.7ms | 360µs | 18.3ms | 274µs
Day 12 | 1.82ms | 451ms | 18.4ms | - | - | 61.8µs | 5.09ms | 1.58ms | -
Day 13 | 270ms | 370µs | **🔴 1.15s** | 6.02ms | - | 24.7µs | 145µs | 2.59ms | -
Day 14 | 3.54ms | **🔴 16.8s** | 524ms | - | - | 6.64ms | 405µs | 8.12ms | -
Day 15 | 9.97s | 27.6µs | 440ms | - | - | **🔴 620ms** | **🔴 296ms** | 2.84µs | -
Day 16 | 1.7ms | 115ms | 8.34ms | - | - | 8.75ms | 45.2µs | **🔴 2.05s** | -
Day 17 | 63.6ms | 33.9ms | 566ms | 39ms | - | 25.7ms | 301µs | 1.8ms | -
Day 18 | 680ms | 228ms | 17.7ms | 412ms | - | 2.89ms | 18.1ms | 1.02ms | -
Day 19 | 4.39s | 205ms | 2.95ms | 54µs | - | 15.8ms | 59ms | **🔴 1.56s** | -
Day 20 | 5.1s | 452µs | 25.6ms | 2.02ms | - | 224ms | 18.9ms | 594ms | -
Day 21 | 29.8µs | 326ms | 562ms | **🔴 39.1s** | - | 705µs | 7.7ms | 569µs | -
Day 22 | 1.36s | 2.65s | 654ms | 1.71s | - | 288ms | 16.7ms | 172ms | -
Day 23 | 81.4µs | 5.53ms | 4.65ms | 8.02ms | - | **🔴 682ms** | **🔴 142ms** | 154ms | -
Day 24 | **🔴 31.9s** | 211ms | 438ms | 101ms | - | 101ms | 1.42ms | 190ms | -
Day 25 | 69.9ms | 432ms | 561ms | 31.7ms | - | 179µs | 139ms | 4.48µs | -
*Total* | *1m2s* | *1m19.8s* | *5.06s* | *44.5s* | *0s* | *2.01s* | *707ms* | *4.75s* | *13ms*

## Heap memory

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022 | 2023
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 17.0 KB | 97.4 KB | 4.9 KB | 6.8 MB | - | 15.0 KB | None | None | None
Day 2 | 372 KB | 12.5 KB | 17.6 KB | 18.9 MB | - | 545 KB | None | None | None
Day 3 | 513 KB | 659 KB | None | 57.3 MB | - | 55.0 KB | 24.6 KB | None | **🔴 2.8 MB**
Day 4 | 454 MB | 10.9 MB | 504 KB | 360 KB | - | 488 KB | 391 KB | None | None
Day 5 | 1.5 MB | 3.8 GB | 66.7 KB | 129 MB | - | 9.8 KB | 45.9 KB | None | 40.2 KB
Day 6 | 218 MB | 31.1 KB | 1.7 MB | 14.3 KB | - | 535 KB | None | None | 112 B
Day 7 | 9.9 MB | 3.3 MB | 922 KB | 170 KB | - | 710 KB | 61.5 KB | None | 101 KB
Day 8 | 64.4 KB | 163 KB | 238 KB | 1.2 MB | - | 798 KB | 116 KB | None | 13.0 KB
Day 9 | 10.6 MB | 163 KB | 22.1 KB | 383 MB | - | 218 KB | 273 KB | None | 326 KB
Day 10 | 185 MB | 5.2 MB | 6.0 KB | **🔴 1.6 GB** | - | 27.0 KB | 14.3 KB | None | 688 KB
Day 11 | 7.9 MB | **🔴 23.0 GB** | 442 KB | 1.6 MB | - | 1.6 MB | 1.1 KB | None | 75.6 KB
Day 12 | 643 KB | 12.0 KB | 2.8 MB | - | - | 20.0 KB | 6.5 MB | None | -
Day 13 | 107 MB | 183 KB | 8.0 KB | 2.9 MB | - | 6.0 KB | 26.6 KB | 578 KB | -
Day 14 | 1.1 MB | 2.7 GB | 27.3 MB | - | - | 2.7 MB | 77.9 KB | None | -
Day 15 | **🔴 3.3 GB** | 6.5 KB | None | - | - | 240 MB | **🔴 103 MB** | None | -
Day 16 | 825 KB | 260 MB | 3.9 MB | - | - | 1.3 MB | 29.0 KB | **🔴 41.8 MB** | -
Day 17 | 2.5 KB | 24.3 MB | 32.3 KB | 12.9 MB | - | 9.0 MB | 192 B | 681 KB | -
Day 18 | 280 MB | 44.8 MB | 14.1 MB | 168 MB | - | 1.4 MB | 1.9 MB | None | -
Day 19 | 411 MB | 48.2 MB | 2.0 MB | 13.7 KB | - | 7.2 MB | 3.5 MB | 7.9 MB | -
Day 20 | **🔴 2.6 GB** | 343 KB | 21.1 MB | 1.0 MB | - | **🔴 1.1 GB** | **🔴 67.1 MB** | None | -
Day 21 | None | 191 MB | 283 MB | 828 KB | - | 372 KB | 2.0 MB | 123 KB | -
Day 22 | 483 MB | 5.3 GB | 6.2 MB | 329 MB | - | 110 MB | 868 KB | 2.0 B | -
Day 23 | 13.5 KB | 3.9 MB | 2.0 MB | 9.1 MB | - | 8.0 MB | 16.1 MB | None | -
Day 24 | 19.4 MB | 35.9 MB | **🔴 1.0 GB** | 188 MB | - | 58.4 MB | 105 KB | 1.0 B | -
Day 25 | None | 290 MB | 430 KB | 1.4 MB | - | 166 KB | 21.4 MB | None | -
*Total* | *8.0 GB* | *35.8 GB* | *1.4 GB* | *2.9 GB* | *None* | *1.6 GB* | *224 MB* | *51.1 MB* | *4.1 MB*