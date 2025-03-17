package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <file.m3u|file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var songs []Song
	if strings.HasSuffix(filename, ".m3u") {
		songs = readM3uPlaylist(string(data))
		writePlsPlaylist(filename, songs)
	} else if strings.HasSuffix(filename, ".pls") {
		songs = readPlsPlaylist(string(data))
		writeM3uPlaylist(filename, songs)
	} else {
		log.Fatalf("Unsupported file format: %s", filename)
	}
}

func readPlsPlaylist(data string) []Song {
	var songs []Song
	scanner := bufio.NewScanner(strings.NewReader(data))
	var song Song
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "[playlist]") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		key, value := parts[0], parts[1]

		switch {
		case strings.HasPrefix(key, "File"):
			song.Filename = normalizePath(value)
		case strings.HasPrefix(key, "Title"):
			song.Title = value
		case strings.HasPrefix(key, "Length"):
			song.Seconds, _ = strconv.Atoi(value)
		}

		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func readM3uPlaylist(data string) []Song {
	var songs []Song
	scanner := bufio.NewScanner(strings.NewReader(data))
	var song Song

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = normalizePath(line)
		}

		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func parseExtinfLine(line string) (title string, seconds int) {
	parts := strings.SplitN(line[8:], ",", 2)
	if len(parts) != 2 {
		return "", -1
	}
	seconds, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("Invalid duration in EXTINF: %v", err)
		return parts[1], -1
	}
	return parts[1], seconds
}

func normalizePath(path string) string {
	return strings.Map(func(r rune) rune {
		if r == '/' || r == '\\' {
			return filepath.Separator
		}
		return r
	}, path)
}

func writeM3uPlaylist(filename string, songs []Song) {
	newFilename := strings.TrimSuffix(filename, ".pls") + ".m3u"
	writePlaylist(newFilename, "#EXTM3U\n", songs, func(song Song) string {
		return fmt.Sprintf("#EXTINF:%d,%s\n%s\n", song.Seconds, song.Title, song.Filename)
	})
}

func writePlsPlaylist(filename string, songs []Song) {
	newFilename := strings.TrimSuffix(filename, ".m3u") + ".pls"
	writePlaylist(newFilename, "[playlist]\n", songs, func(song Song) string {
		return fmt.Sprintf("File%d=%s\nTitle%d=%s\nLength%d=%d\n",
			len(songs), song.Filename, len(songs), song.Title, len(songs), song.Seconds)
	})
}

func writePlaylist(filename, header string, songs []Song, format func(Song) string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	var b strings.Builder
	b.WriteString(header)
	for _, song := range songs {
		b.WriteString(format(song))
	}
	b.WriteString(fmt.Sprintf("NumberOfEntries=%d\nVersion=2\n", len(songs)))

	if _, err := file.Write([]byte(b.String())); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Printf("%s file content:\n%s\n", filename, b.String())
}
