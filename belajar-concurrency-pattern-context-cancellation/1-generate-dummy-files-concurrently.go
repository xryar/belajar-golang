package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000
const timeoutDuration = 3 * time.Second

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.61-pipeline-cancellation-context")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func main() {
	log.Println("start")
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	generateFilesWithContext(ctx)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	generateFilesWithContext(context.Background())
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	done := make(chan int)

	go func() {
		// pipeline 1: job distribution
		chanFileIndex := generateFileIndexes(ctx)

		// pipeline : the main logic (creating files)
		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		//track and print output
		counterTotal := 0
		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
			counterTotal++
		}
		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}
}

func generateFileIndexes(ctx context.Context) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FileInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers)

	go func() {
		// dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						filepath := filepath.Join(tempPath, job.FileName)
						content := randomString(contentLength)
						err := os.WriteFile(filepath, []byte(content), os.ModePerm)

						log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

						// construct the job's result, and send it to `chanOut`
						chanOut <- FileInfo{
							FileName:    job.FileName,
							WorkerIndex: workerIndex,
							Err:         err,
						}
					}

				}

				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
