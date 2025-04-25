// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Image struct {
	width, height int
	filename      string
}

const maxWorkers = 10

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <image files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	files, err := commandLineFiles(os.Args[1:])
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	resultChan := readFiles(ctx, files)
	writeResult(ctx, resultChan)
}

// This function takes a slice of strings as an argument and returns a slice of strings
func commandLineFiles(files []string) ([]string, error) {
	if runtime.GOOS != "windows" {
		return files, nil
	}

	var expanded []string
	var errs []error

	for _, name := range files {
		matches, err := filepath.Glob(name)
		if err != nil {
			errs = append(errs, fmt.Errorf("invalid pattern %q: %w", name, err))
			continue
		}
		if matches == nil {
			expanded = append(expanded, name)
			continue
		}
		expanded = append(expanded, matches...)
	}

	return expanded, errors.Join(errs...)
}

func readFiles(ctx context.Context, files []string) <-chan Image {
	resultChan := make(chan Image, len(files))
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	go func() {
		defer close(resultChan)

		for _, filename := range files {
			if ctx.Err() != nil {
				break
			}

			wg.Add(1)
			sem <- struct{}{}

			go func(f string) {
				defer wg.Done()
				defer func() { <-sem }()
				process(ctx, f, resultChan)
			}(filename)
		}

		wg.Wait()
	}()

	return resultChan
}
func process(ctx context.Context, filename string, result chan<- Image) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	info, err := os.Stat(filename)
	if err != nil {
		log.Printf("Error stating file %s: %v", filename, err)
		return
	}
	if info.Mode()&os.ModeType != 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file %s: %v", filename, err)
		return
	}
	defer file.Close()

	done := make(chan struct{})
	var config image.Config
	var decodeErr error

	go func() {
		config, _, decodeErr = image.DecodeConfig(file)
		close(done)
	}()

	select {
	case <-done:
		if decodeErr != nil {
			log.Printf("Error decoding %s: %v", filename, decodeErr)
			return
		}
	case <-ctx.Done():
		return
	}

	select {
	case result <- Image{config.Width, config.Height, filename}:
	case <-ctx.Done():
	}
}
func writeResult(ctx context.Context, result <-chan Image) {
	for {
		select {
		case img, ok := <-result:
			if !ok {
				return
			}
			fmt.Printf(`<img src="%s" width="%d" height="%d" />`+"\n",
				filepath.Base(img.filename), img.width, img.height)
		case <-ctx.Done():
			log.Println("Processing cancelled during output")
			return
		}
	}
}
