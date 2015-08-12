package main

import (
	"crypto/rand"
	"encoding/base64"
	//"github.com/nu7hatch/gouuid"
	"github.com/pborman/uuid"
	"strconv"
)

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func randomize2(strSize int) string {

	var dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func randomize() string {
	size := 12 // change the length of the generated random string here

	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		tracer.Error("error: " + err.Error())
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs
}

func RandomsSample1() {

	tracer.News("Alphanum : " + randStr(16, "alphanum"))

	tracer.News("Alpha : " + randStr(16, "alpha"))

	tracer.News("Numbers : " + randStr(16, "number"))

}

func RandomsSample2() {
	/*
		u4, err := uuid.NewV4()
		if err != nil {
			tracer.Error("error: " + err.Error())
			return
		}
		tracer.News("v1 - " + u4.String())
	*/
}

func RandomsSample3() {

	for i := 0; i < 100; i++ {
		tracer.News(strconv.Itoa(i) + " - " + randomize())
	}

	tracer.News("------")

	for i := 0; i < 100; i++ {
		tracer.News(strconv.Itoa(i) + " - " + randomize2(12))
	}

}

func RandomsSample4() {

	tracer.News("----------------------")
	for i := 0; i < 100; i++ {
		tracer.News(strconv.Itoa(i) + " - " + uuid.New())
	}

}
