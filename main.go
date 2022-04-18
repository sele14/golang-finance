package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filePath string) []float64 {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()

    var priceList []float64

    for i, row := range records {
        if i > 0 { // drop header
            priceList = append(priceList, s2f(row[0]))
        }
    }

    // ignore header
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return priceList
}

// s2f converts string to float64
func s2f(str string) float64 {
    f, _ := strconv.ParseFloat(str, 64)
    return f
}

func sum(array []float64, winSize int) float64 {
    var res float64
    for i, val := range array {
        // sum up until window size
        if i >= winSize {
            res += val
        }
    }
    return res
}

func movingAverage(winSize int, prices []float64) []float64 {

    var MAs []float64

    for i := 1; i <= winSize; i++ {
        // get sum
        s := sum(prices, i)
        ma := s/float64(i)

        MAs = append(MAs, ma)
    }
    return MAs
}

func main() {
    prices := readCsvFile("prices1.csv")
    fmt.Println(prices)
    
    // calculate moving average
    MA := movingAverage(10, prices)
    fmt.Println(MA)
}