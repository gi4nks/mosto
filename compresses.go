package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
//	"io"
	"io/ioutil"
//	"os"
)

func Compress(value string) string {
    var b bytes.Buffer
    gz := gzip.NewWriter(&b)
    if _, err := gz.Write([]byte(value)); err != nil {
        panic(err)
    }
    if err := gz.Flush(); err != nil {
        panic(err)
    }
    if err := gz.Close(); err != nil {
        panic(err)
    }
    return b.String()
}


func Uncompress(value string) string {
	buf := bytes.NewBufferString(value)
	
	gz, err := gzip.NewReader(buf)
	
	if err != nil {
        panic(err)
    }
	
	b, err := ioutil.ReadAll(gz)
    if err != nil {
    	fmt.Errorf("ReadAll: %v", err)
    }
	
    //io.Copy(os.Stdout, gz)
	
	fmt.Println(string(b))
	
	gz.Close()
	
	return string(b)
}