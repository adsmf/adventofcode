# Benchmarks
The following are the benchmarks for the Go implementations of solutions for each day. The results are as measured by a `BenchmarkMain` benchmark in each solution.

## CPU time

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022 | 2023 | 2024
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 88.7µs | 102µs | 46.4µs | 19ms | 11.8µs | 396µs | 14.4µs | 13.9µs | 579µs | 64.5µs
Day 2 | 390µs | 176µs | 47.8µs | 76.4ms | 272ms | 3.06ms | 1.22µs | 4.18µs | 22.2µs | 75µs
Day 3 | 1.53ms | 507µs | 2ns | 45.7ms | 68.8ms | 83.5µs | 48.2µs | 65.7µs | 2.78ms | 25.9µs
Day 4 | 2.24s | 23ms | 1.3ms | 745µs | 35.7ms | 927µs | 692µs | 23.4µs | 62.8µs | 396µs
Day 5 | 2.1ms | 15.8s | 70.5ms | 112ms | 189µs | 28µs | 704µs | 15.7µs | 35.2µs | 35.7µs
Day 6 | 5.73s | 183µs | 11.4ms | 37.5ms | 1.2ms | 1.34ms | 656ns | 14.5µs | 3.44ms | **🔴 77ms**
Day 7 | 22.9ms | 5.81ms | 2.4ms | 3.2ms | 45.2ms | 1.25ms | 57.1µs | 13.2µs | 312µs | 14.2ms
Day 8 | 182µs | 131µs | 2.42ms | 681µs | 93.2µs | 1.81ms | 535µs | 458µs | 1.68ms | 14.1µs
Day 9 | 19.6ms | 5.1ms | 82.8µs | 650ms | 69.4ms | 819µs | 244µs | 446µs | 114µs | 1.29ms
Day 10 | 176ms | 4.04ms | 1.54ms | 348ms | 16.4ms | 117µs | 74.4µs | 919ns | 3.66ms | 80.8µs
Day 11 | 38ms | **🔴 42.5s** | 462µs | 1.78s | 13.2ms | 20.7ms | 360µs | 18.3ms | 274µs | 428ns
Day 12 | 1.82ms | 451ms | 18.4ms | - | 665ms | 61.8µs | 5.09ms | 1.58ms | 53ms | 1.13ms
Day 13 | 270ms | 370µs | **🔴 1.15s** | 6.02ms | 3.65s | 24.7µs | 145µs | 2.59ms | 1.96ms | 53.7µs
Day 14 | 3.54ms | **🔴 16.8s** | 524ms | - | 4.2ms | 6.64ms | 405µs | 8.12ms | 33ms | 266µs
Day 15 | 9.97s | 27.6µs | 440ms | - | 1.38s | **🔴 620ms** | **🔴 296ms** | 2.84µs | 147µs | 605µs
Day 16 | 1.7ms | 115ms | 8.34ms | - | 254ms | 8.75ms | 45.2µs | **🔴 2.05s** | 6.71ms | 3.22ms
Day 17 | 63.6ms | 33.9ms | 566ms | 39ms | 27ms | 25.7ms | 301µs | 1.8ms | **🔴 208ms** | 1.13µs
Day 18 | 680ms | 228ms | 17.7ms | 412ms | 2.25s | 2.89ms | 18.1ms | 1.02ms | 27.6µs | 14.9ms
Day 19 | 4.39s | 205ms | 2.95ms | 54µs | 2.85s | 15.8ms | 41.8ms | **🔴 1.56s** | 293µs | 20.7µs
Day 20 | 5.1s | 452µs | 25.6ms | 2.02ms | 743ms | 224ms | 18.9ms | 594ms | 25.5ms | 29.4µs
Day 21 | 29.8µs | 326ms | 562ms | **🔴 39.1s** | 106ms | 705µs | 7.7ms | 569µs | 10.3ms | 18.3ms
Day 22 | 1.36s | 2.65s | 654ms | 1.71s | - | 288ms | 16.7ms | 172ms | 9.84ms | 57.4ms
Day 23 | 81.4µs | 5.53ms | 4.65ms | 8.02ms | 1.4s | **🔴 682ms** | 92.1ms | 154ms | **🔴 199ms** | **🔴 181ms**
Day 24 | **🔴 31.9s** | 211ms | 438ms | 101ms | 246ms | 101ms | 1.42ms | 190ms | 1.04ms | 257µs
Day 25 | 69.9ms | 432ms | 561ms | 31.7ms | **🔴 11.7s** | 179µs | **🔴 139ms** | 4.48µs | - | 375µs
*Total* | *1m2s* | *1m19.8s* | *5.06s* | *44.5s* | *25.8s* | *2.01s* | *640ms* | *4.75s* | *562ms* | *370ms*

## Heap memory

 &nbsp;  | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022 | 2023 | 2024
 ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---:  | ---: 
Day 1 | 17.0 KB | 97.4 KB | 4.9 KB | 6.8 MB | 3.5 KB | 15.0 KB | None | None | None | None
Day 2 | 372 KB | 12.5 KB | 17.6 KB | 18.9 MB | 255 MB | 545 KB | None | None | None | None
Day 3 | 513 KB | 659 KB | None | 57.3 MB | 102 MB | 55.0 KB | 24.6 KB | None | 2.8 MB | None
Day 4 | 454 MB | 10.9 MB | 504 KB | 360 KB | 22.9 MB | 488 KB | 391 KB | None | None | None
Day 5 | 1.5 MB | 3.8 GB | 66.7 KB | 129 MB | 230 KB | 9.8 KB | 45.9 KB | None | 40.2 KB | None
Day 6 | 218 MB | 31.1 KB | 1.7 MB | 14.3 KB | 1.6 MB | 535 KB | None | None | 112 B | None
Day 7 | 9.9 MB | 3.3 MB | 922 KB | 170 KB | 61.9 MB | 710 KB | 61.5 KB | None | 101 KB | None
Day 8 | 64.4 KB | 163 KB | 238 KB | 1.2 MB | 104 KB | 798 KB | 116 KB | None | 13.0 KB | None
Day 9 | 10.6 MB | 163 KB | 22.1 KB | 383 MB | 46.4 MB | 218 KB | 273 KB | None | 326 KB | None
Day 10 | 185 MB | 5.2 MB | 6.0 KB | **🔴 1.6 GB** | 22.6 MB | 27.0 KB | 14.3 KB | None | 688 KB | None
Day 11 | 7.9 MB | **🔴 23.0 GB** | 442 KB | 1.6 MB | 2.5 MB | 1.6 MB | 1.1 KB | None | 75.6 KB | None
Day 12 | 643 KB | 12.0 KB | 2.8 MB | - | 227 MB | 20.0 KB | 6.5 MB | None | **🔴 108 MB** | None
Day 13 | 107 MB | 183 KB | 8.0 KB | 2.9 MB | **🔴 14.3 GB** | 6.0 KB | 26.6 KB | 578 KB | 1.3 MB | None
Day 14 | 1.1 MB | 2.7 GB | 27.3 MB | - | 360 KB | 2.7 MB | 77.9 KB | None | 87.6 KB | None
Day 15 | **🔴 3.3 GB** | 6.5 KB | None | - | 1.4 GB | 240 MB | **🔴 103 MB** | None | 17.9 KB | None
Day 16 | 825 KB | 260 MB | 3.9 MB | - | 481 MB | 1.3 MB | 29.0 KB | **🔴 41.8 MB** | 44.5 MB | None
Day 17 | 2.5 KB | 24.3 MB | 32.3 KB | 12.9 MB | 24.9 MB | 9.0 MB | 192 B | 681 KB | **🔴 76.9 MB** | None
Day 18 | 280 MB | 44.8 MB | 14.1 MB | 168 MB | 1.1 GB | 1.4 MB | 1.9 MB | None | None | None
Day 19 | 411 MB | 48.2 MB | 2.0 MB | 13.7 KB | 2.9 GB | 7.2 MB | 3.7 MB | 7.9 MB | 252 KB | None
Day 20 | **🔴 2.6 GB** | 343 KB | 21.1 MB | 1.0 MB | 554 MB | **🔴 1.1 GB** | **🔴 67.1 MB** | None | 27.4 KB | None
Day 21 | None | 191 MB | 283 MB | 828 KB | 60.6 MB | 372 KB | 2.0 MB | 123 KB | 8.2 MB | 82.9 KB
Day 22 | 483 MB | 5.3 GB | 6.2 MB | 329 MB | - | 110 MB | 868 KB | 2.0 B | 434 KB | None
Day 23 | 13.5 KB | 3.9 MB | 2.0 MB | 9.1 MB | 801 MB | 8.0 MB | 16.1 MB | None | 9.4 MB | **🔴 133 MB**
Day 24 | 19.4 MB | 35.9 MB | **🔴 1.0 GB** | 188 MB | 183 MB | 58.4 MB | 105 KB | 1.0 B | 128 KB | 50.1 KB
Day 25 | None | 290 MB | 430 KB | 1.4 MB | 101 MB | 166 KB | 21.4 MB | None | - | None
*Total* | *8.0 GB* | *35.8 GB* | *1.4 GB* | *2.9 GB* | *22.6 GB* | *1.6 GB* | *224 MB* | *51.1 MB* | *253 MB* | *133 MB*