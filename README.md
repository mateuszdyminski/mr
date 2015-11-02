# Dummy Map-Reduce implementation written in GO

## Goal 

Write simple Map-Reduce to calculate Lotto lottery histogram.
In file `data` you can find all historically results of Lotto. File `extracted` contains only numbers without the date of the result.

## Requirements 

* Go >= 1.3 

## Run

```
go run main.go master extracted sequential
```

## Results

To sum up results: 

In next lottery we should choose numbers: `48, 43, 12, 47, 44, 49`


#### Histogram
![alt text](https://github.com/mateuszdyminski/mr/raw/master/histogram.png "Histogram")

#### Sorted results

| Numbers | Count  |
| ------- | ------:|
| 48      | 623    |
| 43      | 630    |
| 12      | 648    |
| 47      | 652    |
| 44      | 654    |
| 49      | 659    |
| 7       | 662    |
| 16      | 665    |
| 8       | 666    |
| 10      | 674    |
| 23      | 675    |
| 39      | 681    |
| 30      | 682    |
| 41      | 682    |
| 46      | 683    |
| 3       | 689    |
| 2       | 691    |
| 35      | 691    |
| 14      | 692    |
| 5       | 692    |
| 11      | 695    |
| 9       | 695    |
| 37      | 696    |
| 40      | 696    |
| 26      | 699    |
| 33      | 699    |
| 20      | 704    |
| 19      | 705    |
| 1       | 706    |
| 32      | 706    |
| 22      | 708    |
| 25      | 708    |
| 15      | 709    |
| 38      | 714    |
| 28      | 717    |
| 36      | 717    |
| 45      | 721    |
| 42      | 723    |
| 18      | 724    |
| 29      | 728    |
| 31      | 728    |
| 6       | 728    |
| 4       | 729    |
| 13      | 730    |
| 27      | 734    |
| 24      | 737    |
| 21      | 740    |
| 34      | 748    |
| 17      | 753    |

## Implementation in based on MIT 6.824 course 

Details here:
http://css.csail.mit.edu/6.824/2014/

