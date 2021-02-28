package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    /*
     * Write your code here.
     */


     if s[8] == 'A' {
        
        return s
    } else {
        var i int
        i , _ = strconv.Atoi(s[0:2])
        
        i = i + 12
        temp_string := strconv.Itoa(i)
        
        var ss string = s
        temp_step := replaceAtIndex(ss, rune(temp_string[0]), 0)
        returnluk := replaceAtIndex(temp_step, rune(temp_string[1]), 1)
        
        return returnluk[0:8]
    }
    
} 


func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    
    return string(out)
}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
