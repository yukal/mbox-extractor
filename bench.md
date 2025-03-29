Testing a 20 MB file:

|            | runs | cursor | file size | sec/op  |  B/op   | allocs/op |
| :---       | ---: |   ---: | :---      |    ---: |    ---: |      ---: |
| benchmark1 |   10 |     no | 22.1MB    | 1.7808s | 87.98Mi |    4.128k |
| benchmark2 |   10 |    yes | 22.1MB    | 0.1665s | 87.98Mi |    4.128k |
| vs base    |      |        |           | -90.65% |     ±0% |       ±0% |

Testing a 100 MB file:

|            | runs | cursor | file size | sec/op  |  B/op   | allocs/op |
| :---       | ---: |   ---: | :---      |    ---: |    ---: |      ---: |
| benchmark1 |   10 |     no | 110.6MB   | 9.1512s | 439.2Mi |    22.16k |
| benchmark2 |   10 |    yes | 110.6MB   | 0.8505s | 439.3Mi |    22.13k |
| vs base    |      |        |           | -90.71% |  +0.01% |    -0.14% |

Testing a 500 MB file:

|            | runs | cursor | file size | sec/op  |  B/op   | allocs/op |
| :---       | ---: |   ---: | :---      |    ---: |    ---: |      ---: |
| benchmark1 |   10 |     no | 509MB     | 42.005s | 1.973Gi |    103.1k |
| benchmark2 |   10 |    yes | 509MB     |  2.537s | 1.973Gi |    103.0k |
| vs base    |      |        |           | -93.96% |  +0.01% |    -0.12% |

Testing a 1 GB file:

|            | runs | cursor | file size | sec/op  |  B/op   | allocs/op |
| :---       | ---: |   ---: | :---      |    ---: |    ---: |      ---: |
| benchmark1 |    6 |     no | 1GB       | 84.803s | 4.032Gi |    211.2k |
| benchmark2 |    6 |    yes | 1GB       |  5.017s | 4.032Gi |    210.9k |
| vs base    |      |        |           | -94.08% |  +0.01% |    -0.15% |
