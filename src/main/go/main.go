package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

/*
#cgo CFLAGS: -I/usr/lib/jvm/java-8-openjdk-amd64/include
#cgo CFLAGS: -I/usr/lib/jvm/java-8-openjdk-amd64/include/linux
#include <jni.h>
#include <stdio.h>
#include <stdlib.h>
static const char* convert_to_cstring(JNIEnv *env, jstring javaString)
{
   return (*env)->GetStringUTFChars(env, javaString, 0);
}
static jstring convert_to_jstring(JNIEnv *env, char* buf)
{
	return (*env)->NewStringUTF(env, buf);
}
*/
import "C"

var (
	javaResponse = make(chan string)
	goRequest    = make(chan string)
)

func communicateJava(input interface{}) (interface{}, error) {
	inBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	inputStr := string(inBytes)

	// Write to Go
	log.Println("Received request")
	goRequest <- inputStr

	// Read from Java
	respStr := <-javaResponse

	var resp interface{}
	err = json.Unmarshal([]byte(respStr), &resp)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

//export Java_example_Handler_start
func Java_example_Handler_start(env *C.JNIEnv, clazz C.jclass) {
	log.SetPrefix("GO - ")

	go func() {
		lambda.Start(communicateJava)
	}()
}

//export Java_example_Handler_writeResponse
func Java_example_Handler_writeResponse(env *C.JNIEnv, clazz C.jclass, input C.jstring) {
	a := C.convert_to_cstring(env, input)
	b := C.GoString(a)
	javaResponse <- b
}

//export Java_example_Handler_readRequest
func Java_example_Handler_readRequest(env *C.JNIEnv, clazz C.jclass) C.jstring {
	input := <-goRequest
	cstr := C.CString(input)
	cjstring := C.convert_to_jstring(env, cstr)
	return cjstring
}

func main() {}
