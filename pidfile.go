// Copyright 2019 tinystack/dreamans Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pidfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type PID struct {
	path string
}

func New(path string) *PID {
	p := &PID{
		path: path,
	}
	return p
}

func (p *PID) Create() error {
	if pid := p.ProcessID(); pid > 0 {
		return errors.New("pidfile already exists")
	}
	if err := ioutil.WriteFile(p.path, []byte(fmt.Sprintf("%d", os.Getpid())), 0755); err != nil {
		return err
	}
	return nil
}

func (p *PID) Remove() error {
	return os.Remove(p.path)
}

func (p *PID) ProcessID() int {
	if pidByte, err := ioutil.ReadFile(p.path); err == nil {
		pidString := strings.TrimSpace(string(pidByte))
		if pid, err := strconv.Atoi(pidString); err == nil && processExists(pid) {
			return pid
		}
	}
	return 0
}
