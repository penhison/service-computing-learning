package main

import (
    flag "github.com/spf13/pflag"
    "fmt"
    "os"
    "io"
    "os/exec"
)

const INBUFSIZ = 16 * 1024

type selpg_args struct {
	start_page int
	end_page int
	page_len int
	page_type bool
	in_filename string
	print_dest string

}

var sp_args selpg_args

func init() {
	flag.IntVarP(&sp_args.start_page, "start", "s", 1, "start page")
	flag.IntVarP(&sp_args.end_page, "end", "e", 10000000, "end page")
	flag.IntVarP(&sp_args.page_len, "pagelen", "l", 72, "page fixed line count")
	flag.BoolVarP(&sp_args.page_type, "f", "f", false, "page seperated by \"\\f\"")
	flag.StringVarP(&sp_args.print_dest, "destination", "d", "", "send to printer destination")
	flag.Usage = usage
	// flag.SortFlags = false
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: selpg -s startpage -e endpage [-l pageline | -f] [-d Destination] [inputfile]\nOptions: \n")
	flag.PrintDefaults()
}

func checkArg() bool {
	res := true
	if sp_args.start_page < 1 {
		fmt.Fprintf(os.Stderr, "start_page must >= 1")
		res = false
	}
	if sp_args.end_page < sp_args.start_page {
		fmt.Fprintf(os.Stderr, "sp_args.end_page must >= sp_args.start_page")
		res = false
	}
	if sp_args.page_len < 1 {
		fmt.Fprintf(os.Stderr, "page_len must >= 1")
		res = false
	}
	if !res {
		os.Exit(1)
	}
	return res
}

func parse() {
	flag.Parse()
	if flag.NArg() > 0 {
		sp_args.in_filename = flag.Arg(0)
	}
}

func process() {
	fin, fout, ferr := os.Stdin, io.WriteCloser(os.Stdout), os.Stderr
	var err error
	if sp_args.in_filename != "" {
		fin, err = os.Open(sp_args.in_filename)
		if err != nil {
			fmt.Fprintf(ferr, "%v\n", err)
			os.Exit(1)
		}
		defer fin.Close()
	}

	if sp_args.print_dest != "" {
		cmd := exec.Command("lp", "-d", sp_args.print_dest)
		fout, err = cmd.StdinPipe()
		if err != nil {
			fmt.Fprintf(ferr, "%v\n", err)
			os.Exit(1)
		}
		defer fout.Close()
		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(ferr, "%v\n", err)
			os.Exit(1)
		}
	}

	count := INBUFSIZ
	line_count, page_count := 0, 1
	start := 0
	// fmt.Fprintf(fout, "out")
	data := make([]byte, INBUFSIZ)
	if sp_args.page_type == false {
		stop := false
		for !stop {
			count, err = fin.Read(data)
			start = 0
			for i := 0; i < count; i++ {
				if data[i] == '\n' {
					if page_count >= sp_args.start_page && page_count <= sp_args.end_page {
						fout.Write(data[start:i+1])
					}
					start = i + 1
					line_count++
					if line_count >= sp_args.page_len {
						page_count++
						line_count = 0
						if page_count > sp_args.end_page {
							stop = true
							break;
						}
					}
				}
			}
			if start < count && page_count >= sp_args.start_page && page_count <= sp_args.end_page {
				fout.Write(data[start:count])
			}
			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(ferr, "%v\n", err)
				}
				break;
			}
		}
	} else {
		stop := false
		for !stop {
			count, err = fin.Read(data)
			start = 0
			for i := 0; i < count; i++ {
				if data[i] == '\f' {
					if page_count >= sp_args.start_page && page_count <= sp_args.end_page {
						fout.Write(data[start:i+1])
					}
					start = i + 1
					page_count++;
					if page_count > sp_args.end_page {
						stop = true
						break
					}
				}
			}
			if start < count && page_count >= sp_args.start_page && page_count <= sp_args.end_page {
				fout.Write(data[start:count])
			}
			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(ferr, "%v\n", err)
				}
				break;
			}
		}
	}
}

func main() {
	parse()
	checkArg()
	process()
}