# Advent of Code
My solutions to puzzles from http://adventofcode.com/

## Goals

In various years I have had different goals when taking part in AoC:

* 2018 - learned about AoC somewhere around day 18 - had fun with a few challenges and made a note to come back the next year
* 2019 - Took part every day.  Enjoyed enough to start working through the backlog for previous years
* 2020 - Looking to make more performant code but without sacrificing too much readability
* 2021 - Self imposed challenge: runtime for all solutions should be <= 1s as measured by go benchmark run on my laptop
  * Result: ✅ total execution time for all 25 days solutions is 707ms
* 2022 - Self imposed challenge: try and eliminate all heap memory allocations
  * Result:
    * ✅ 18 days with zero heap allocation
    * ❌ 7 days needed heap allocation with a total of 51.1MB of heap memory used
    * ❌ Very ugly code!
* 2023 - A planned return to better looking code.  Will still try and maintain performance but not at all costs.
