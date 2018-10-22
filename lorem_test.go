// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import "testing"
import "log"

func TestAll(t *testing.T) {
	l := New()
	for i := 1; i < 14; i++ {
		log.Print(l.word(i))
		for j := 1; j < 14; j++ {
			log.Print(l.Word(i, j))
			log.Print(l.Sentence(i, j))
			log.Print(l.Paragraph(i, j))
		}
		log.Print(l.Url())
		log.Print(l.Host())
		log.Print(l.Email())
	}
}
