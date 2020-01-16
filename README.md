# csv2json
convert csv to json using header column as json key

```
go install github.com/tmrtmhr/csv2json
```

# Usage

```
csv2json < sample.tsv
```

# Sample Input/Output

## Input

```
col1	col2	col3
1	2	3
4	"5
67"	8
9	10	11
```

## Output(JSON)

Command

```
csv2json < sample.csv
```

```
[
  {
    "col1": "1",
    "col2": "2",
    "col3": "3"
  },
  {
    "col1": "4",
    "col2": "5\n67",
    "col3": "8"
  },
  {
    "col1": "9",
    "col2": "10",
    "col3": "11"
  }
]
```

## Output(JSON Seq)

Command

```
csv2json -jsonSeq < sample.csv
```

Output

```
{"col1":"1","col2":"2","col3":"3"}
{"col1":"4","col2":"5\n67","col3":"8"}
{"col1":"9","col2":"10","col3":"11"}
```
