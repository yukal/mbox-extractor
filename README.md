# mbox-extractor
Extracts and splits email messages from a large .mbox file into smaller, chunked files

|            |      sec/op        |       B/op       |   allocs/op   |     runs       |
| :---       |               ---: |             ---: |          ---: |           ---: |
| benchmark1 | 1806867.067µ ± 6%  | 90089.070Ki ± 0% | 4131.000 ± 0% | (p=0.000 n=10) |
| benchmark2 |       5.926µ ± 23% |     4.345Ki ± 2% |    1.000 ± 0% | (p=0.000 n=10) |
| vs base    |    -100.00%        |  -100.00%        |  -99.98%      | (p=0.000 n=10) |
