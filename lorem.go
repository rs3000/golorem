// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import (
	"math/rand"
	"strings"
	"time"
)

type Lorem struct {
	r *rand.Rand
}

func New() *Lorem {
	return &Lorem{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate a natural word len.
func (self *Lorem) genWordLen() int {
	n := self.r.Intn(10000)
	// a table of word lengths and their frequencies.
	switch {
	case n < 193:
		return 1
	case n < 1901:
		return 2
	case n < 3800:
		return 3
	case n < 5041:
		return 4
	case n < 6100:
		return 5
	case n < 7009:
		return 6
	case n < 7897:
		return 7
	case n < 8565:
		return 8
	case n < 9087:
		return 9
	case n < 9505:
		return 10
	case n < 9727:
		return 11
	case n < 9867:
		return 12
	default:
		return 13
	}
}

func (self *Lorem) intRange(min, max int) int {
	if min == max {
		return self.intRange(min, min+1)
	}
	if min > max {
		return self.intRange(max, min)
	}
	n := self.r.Int() % (max - min)
	return n + min
}

func (self *Lorem) word(wordLen int) string {
	if wordLen < 1 {
		wordLen = 1
	}
	if wordLen > 13 {
		wordLen = 13
	}

	n := self.r.Int() % len(wordlist)
	for {
		if n >= len(wordlist)-1 {
			n = 0
		}
		if len(wordlist[n]) == wordLen {
			return wordlist[n]
		}
		n++
	}
	return ""
}

// Generate a word in a specfied range of letters.
func (self *Lorem) Word(min, max int) string {
	n := self.intRange(min, max)
	return self.word(n)
}

// Generate a sentence with a specified range of words.
func (self *Lorem) Sentence(min, max int) string {
	n := self.intRange(min, max)

	// grab some words
	ws := []string{}
	maxcommas := 2
	numcomma := 0
	for i := 0; i < n; i++ {
		word := self.word(self.genWordLen())
		if i == 0 {
			word = strings.ToUpper(word[:1]) + word[1:]
		}
		ws = append(ws, word)

		// maybe insert a comma, if there are currently < 2 commas, and
		// the current word is not the last or first
		if (self.r.Int()%n == 0) && numcomma < maxcommas && i < n-1 && i > 2 {
			ws[i-1] += ","
			numcomma += 1
		}

	}

	return strings.Join(ws, " ") + "."
}

// Generate a paragraph with a specified range of sentenences.
const (
	minwords = 5
	maxwords = 22
)

func (self *Lorem) Paragraph(min, max int) string {
	n := self.intRange(min, max)

	p := []string{}
	for i := 0; i < n; i++ {
		p = append(p, self.Sentence(minwords, maxwords))
	}
	return strings.Join(p, " ")
}

// Generate a random URL
func (self *Lorem) Url() string {
	n := self.intRange(0, 3)

	base := `http://www.` + self.Host()

	switch n {
	case 0:
		break
	case 1:
		base += "/" + self.Word(2, 8)
	case 2:
		base += "/" + self.Word(2, 8) + "/" + self.Word(2, 8) + ".html"
	}
	return base
}

// Host
func (self *Lorem) Host() string {
	n := self.intRange(0, 3)
	tld := ""
	switch n {
	case 0:
		tld = ".com"
	case 1:
		tld = ".net"
	case 2:
		tld = ".org"
	}

	parts := []string{self.Word(2, 8), self.Word(2, 8), tld}
	return strings.Join(parts, ``)
}

// Email
func (self *Lorem) Email() string {
	return self.Word(4, 10) + `@` + self.Host()
}
