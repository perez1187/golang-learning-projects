package service

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/polly"
// 	"github.com/aws/aws-sdk-go-v2/service/polly/types"
// 	"github.com/aws/aws-sdk-go/aws"
// 	// "github.com/aws/aws-sdk-go/service/polly"
// )

// // we create an interface
// type PollyService interface {
// 	// text what we want to senthezi and filename
// 	// error, in case of any errors
// 	Synthesize(text string, fileName string) error
// }

// // struct with configuration of voices
// type pollyConfig struct {
// 	voice string
// }

// const (
// 	// AUDIO_FORMAT = "mp3"
// 	KIMBERLY_VOICE = "Kimberly"
// 	JACEK_VOICE    = "Jacek"
// )

// // constructor function for many voices (optional?)
// func NewKimberlyPollyService() PollyService {
// 	return &pollyConfig{
// 		voice: KIMBERLY_VOICE,
// 	}
// }

// func NewJacekPollyService() PollyService {
// 	return &pollyConfig{
// 		voice: JACEK_VOICE,
// 	}
// }

// func createPollyClient() *polly.Client {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		log.Fatalf("failed to load configuration, %v", err)
// 	}

// 	fmt.Print(cfg)
// 	// client := polly.NewFromConfig(cfg)
// 	return polly.NewFromConfig(cfg)
// }

// func (config *pollyConfig) Synthesize(text string, fileName string) error {
// 	pollyClient := createPollyClient()

// 	// now we create input
// 	input := &polly.SynthesizeSpeechInput{
// 		// we need to pass 3 properties
// 		OutputFormat: types.OutputFormatMp3,
// 		Text:         aws.String(text),
// 		VoiceId:      types.VoiceIdJoanna,
// 	}

// 	output, err := pollyClient.SynthesizeSpeech(context.TODO(), input)
// 	if err != nil {
// 		return err
// 	}

// 	// he is creating mp3 file
// 	outFile, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}

// 	defer outFile.Close()

// 	// we copy what we fetch from polly to file
// 	_, err = io.Copy(outFile, output.AudioStream)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
