package main

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	//loading config data ~.aws/config
	// config.DefaultSharedConfigFilename()
	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatalf("failed to load configuration, %v", err)
	// }

	// fmt.Println(cfg)

	// hardocoded credentials
	//https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKIAUTAF6IFYRMUVVA6X", "JZicrSUFI8GBjeSa8kuH3PmJGML8zvIfbsL64Izr", "")), config.WithRegion("eu-central-1"),
	)

	// create polly client
	client := polly.NewFromConfig(cfg)
	// client := polly.New(sess)

	// create input
	input := &polly.SynthesizeSpeechInput{
		// we need to pass 3 properties
		OutputFormat: types.OutputFormatMp3,
		Text:         aws.String("hi ! Yaba daba du"),
		VoiceId:      types.VoiceIdJoanna,
	}

	// catch output
	// SDK V1 client.SynthesizeSpeech(input)
	output, err := client.SynthesizeSpeech(context.TODO(), input)

	if err != nil {
		panic(err)
	}
	// fmt.Print(output)

	// test
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	checkErr(err)
	// 	return
	// }
	// src, err := file.Open()
	// if err != nil {
	// 	checkErr(err)
	// 	return
	// }
	// defer src.Close()

	dir := "tmp"
	// dst, err := os.Create(filepath.Join(dir, filepath.Base(filename.Filename))) // dir is directory where you want to save file.

	// finish test
	fileName := "qork3.mp3"

	// create file
	// outFile, err := os.Create(fileName) // previus
	// https://stackoverflow.com/questions/48349927/how-to-write-a-directory-with-a-file-in-golang
	outFile, err := os.Create(filepath.Join(dir, filepath.Base(fileName)))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// we copy what we fetch from polly to file
	_, err = io.Copy(outFile, output.AudioStream)

	if err != nil {
		panic(err)
	}
}
