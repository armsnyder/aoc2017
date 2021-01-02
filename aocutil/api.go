package aocutil

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetInput(day int) io.ReadCloser {
	puzzleInputFilename := filepath.Join("inputs", fmt.Sprintf("%02d.txt", day))

	if _, err := os.Lstat(puzzleInputFilename); errors.Is(err, os.ErrNotExist) {
		session, err := ioutil.ReadFile("session.txt")
		if err != nil {
			panic(err)
		}
		session = bytes.TrimSpace(session)

		resp := makeAOCRequest(http.MethodGet, fmt.Sprintf("/2017/day/%d/input", day), "", nil)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			panic(resp.Status)
		}

		if err := os.MkdirAll("inputs", 0755); err != nil {
			panic(err)
		}

		out, err := os.Create(puzzleInputFilename)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			panic(err)
		}
	}

	file, err := os.Open(puzzleInputFilename)
	if err != nil {
		panic(err)
	}

	return file
}

func Submit(day int, part2 bool, v interface{}) {
	answer := fmt.Sprint(v)

	solutions, err := os.OpenFile("solutions.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer solutions.Close()

	level := "1"
	if part2 {
		level = "2"
	}
	key := fmt.Sprintf("%02d.%s=", day, level)

	scanner := bufio.NewScanner(solutions)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), key) {
			correctAnswer := strings.TrimPrefix(scanner.Text(), key)
			if answer == correctAnswer {
				fmt.Print("Correct")
			} else {
				fmt.Print("Incorrect")
			}
			fmt.Println(" (Already submitted)")
			return
		}
	}

	fmt.Println("Submitting...")

	values := make(url.Values)
	values.Set("answer", answer)
	values.Set("level", level)

	resp := makeAOCRequest(http.MethodPost, fmt.Sprintf("/2017/day/%d/answer", day),
		"application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if _, err := io.Copy(os.Stderr, resp.Body); err != nil {
			panic(err)
		}
		panic(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	message := string(regexp.MustCompile("<p[^>]*>(.*)</p>").FindSubmatch(body)[1])
	fmt.Println(message)

	if strings.Contains(message, "That's the right answer!") ||
		strings.Contains(message, "Did you already complete it?") {
		if _, err := solutions.WriteString(key + answer + "\n"); err != nil {
			panic(err)
		}
	}
}

func makeAOCRequest(method, path, contentType string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, "https://adventofcode.com"+path, body)
	if err != nil {
		panic(err)
	}

	if contentType != "" {
		req.Header.Set("content-type", contentType)
	}

	session, err := ioutil.ReadFile("session.txt")
	if err != nil {
		panic(err)
	}
	session = bytes.TrimSpace(session)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(session),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}
