package main

//go:generate go run .

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/log/jot"
)

const (
	cefBaseDir    = "J:/zmcommand.sciter/msys2/usr/local/cef"
	outputBaseDir = "../../cef"
)

var (
	sdefsMap = make(map[string]*structDef)
	edefsMap = make(map[string]*enumDef)
	tdefsMap = make(map[string]*typeDef)
	fdefsMap = make(map[string]*funcDef)
)

type lineInfo struct {
	Line     string
	Position position
}

func main() {
	headers := capiHeaders()
	examineCEFSource(headers)
	cleanOutput()
	createCommonHeader(headers)
	dumpStructs()
	dumpEnums()
	dumpTypedefs()
	dumpFunctions()
}

func cleanOutput() {
	f, err := os.Open(outputBaseDir)
	jot.FatalIfErr(err)
	list, err := f.Readdir(-1)
	jot.FatalIfErr(err)
	jot.FatalIfErr(f.Close())
	for _, one := range list {
		path := filepath.Join(outputBaseDir, one.Name())
		if strings.HasSuffix(path, "_gen.go") || strings.HasSuffix(path, "_gen.c") || strings.HasSuffix(path, "_gen.h") {
			jot.FatalIfErr(os.Remove(path))
		}
	}
}

func examineCEFSource(headers []string) {
	var err error
	cmd := exec.Command("clang", clangArgs(headers)...)
	cmd.Dir, err = filepath.Abs(cefBaseDir)
	jot.FatalIfErr(err)
	stdout, err := cmd.StdoutPipe()
	jot.FatalIfErr(err)
	stderr, err := cmd.StderrPipe()
	jot.FatalIfErr(err)
	var wg sync.WaitGroup
	wg.Add(2)
	go scanStdout(&wg, stdout)
	go scanStderr(&wg, stderr)
	jot.FatalIfErr(cmd.Start())
	wg.Wait()
}

func scanStdout(wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()
	var pos position
	var curBlock, prevBlock []lineInfo
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		pos.update(line)
		if len(line) > 1 {
			if strings.HasPrefix(line[1:], "-") { // Look for top-level declarations (start with '|-' or '`-')
				if len(curBlock) > 0 {
					processBlock(curBlock, prevBlock)
				}
				prevBlock = curBlock
				curBlock = make([]lineInfo, 0, 16)
			}
			curBlock = append(curBlock, lineInfo{Line: line, Position: pos})
		}
	}
	if len(curBlock) > 0 {
		processBlock(curBlock, prevBlock)
	}
}

func processBlock(curBlock, prevBlock []lineInfo) {
	first := curBlock[0].Line[1:]
	switch {
	case strings.HasPrefix(first, "-RecordDecl "):
		processRecordDecl(curBlock)
	case len(prevBlock) > 0 && strings.HasPrefix(prevBlock[0].Line, "|-EnumDecl ") && strings.HasPrefix(first, "-TypedefDecl "):
		processTypedefDeclEnumDecl(curBlock, prevBlock)
	case strings.HasPrefix(first, "-TypedefDecl "):
		processTypedefDecl(curBlock)
	case strings.HasPrefix(first, "-FunctionDecl "):
		processFunctionDecl(curBlock)
	}
}

func genSourceFile(tmpl *template.Template, tmplName, fileName string, data interface{}) {
	var buffer bytes.Buffer
	jot.FatalIfErr(tmpl.ExecuteTemplate(&buffer, tmplName, data))
	path := filepath.Join(outputBaseDir, fileName)
	var d []byte
	if strings.HasSuffix(fileName, ".go") {
		var err error
		d, err = format.Source(buffer.Bytes())
		if err != nil {
			jot.Error(errs.NewWithCause(path, err))
			d = buffer.Bytes()
		}
	} else {
		d = buffer.Bytes()
	}
	jot.FatalIfErr(ioutil.WriteFile(path, d, 0644))
}

func scanStderr(wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		jot.Error(scanner.Text())
	}
}

func capiHeaders() []string {
	var headers []string
	for _, one := range headerList(filepath.Join(cefBaseDir, "include", "capi")) {
		if !one.IsDir() {
			name := one.Name()
			if name != "cef_parser_capi.h" &&
				name != "cef_thread_capi.h" &&
				name != "cef_trace_capi.h" &&
				!strings.HasSuffix(name, "_reader_capi.h") &&
				!strings.HasSuffix(name, "_util_capi.h") {
				header := filepath.Join("include", "capi", name)
				headers = append(headers, header)
			}
		}
	}
	for _, one := range headerList(filepath.Join(cefBaseDir, "include", "capi", "views")) {
		if !one.IsDir() {
			headers = append(headers, filepath.Join("include", "capi", "views", one.Name()))
		}
	}
	sort.Strings(headers)
	return headers
}

func headerList(dir string) []os.FileInfo {
	f, err := os.Open(dir)
	jot.FatalIfErr(err)
	list, err := f.Readdir(-1)
	jot.FatalIfErr(err)
	jot.FatalIfErr(f.Close())
	return list
}

func clangArgs(headers []string) []string {
	args := make([]string, 0, len(headers)+7)
	args = append(args, "-I")
	args = append(args, ".")
	args = append(args, "-Xclang")
	args = append(args, "-ast-dump")
	args = append(args, "-fsyntax-only")
	args = append(args, "-fno-color-diagnostics")
	args = append(args, "-Wno-visibility")
	args = append(args, headers...)
	return args
}

func createCommonHeader(headers []string) {
	f, err := os.Create(filepath.Join(outputBaseDir, "capi_gen.h"))
	jot.FatalIfErr(err)
	f.WriteString(`// Code generated - DO NOT EDIT.

#ifndef GOCEF_CAPI_H_
#define GOCEF_CAPI_H_
#pragma once

#include <stdlib.h>
`)
	for _, header := range headers {
		fmt.Fprintf(f, `#include "%s"
`, header)
	}
	f.WriteString(`
#endif // GOCEF_CAPI_H_
`)
	jot.FatalIfErr(f.Close())
}
