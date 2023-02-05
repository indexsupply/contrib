package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"github.com/golang/snappy"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		start, count        uint
		recs, blks          bool
		ancientDir, outfile string
	)
	flag.UintVar(&start, "h", 0, "starting block height")
	flag.UintVar(&count, "n", 1, "get 'n' blocks after specified height")
	flag.BoolVar(&recs, "r", false, "get receipts")
	flag.BoolVar(&blks, "b", false, "get blocks")
	flag.StringVar(&outfile, "o", "", "write results to file instead of STDOUT")
	flag.StringVar(&ancientDir, "dir", "", "loction of geth's ancient data")
	flag.Parse()

	if !recs && !blks {
		fmt.Fprintf(os.Stderr, "must specify -r or -b\n")
		os.Exit(1)
	}
	if blks {
		//TODO(r): support blocks
		fmt.Fprintf(os.Stderr, "blocks not supported\n")
		os.Exit(1)
	}
	if recs {
		for _, r := range receipts(start, count, ancientDir) {
			fmt.Printf("%x\n", r)
		}
	}
}

func receipts(start, count uint, dir string) [][]byte {
	cidx, err := os.ReadFile(dir + "/receipts.cidx")
	check(err)
	type ptr struct {
		file   uint16
		offset uint32
	}
	index := map[uint]ptr{}
	for i, bn := (start * 6), start; bn < start+count+1; i += 6 {
		index[bn] = ptr{
			file:   binary.BigEndian.Uint16(cidx[i : i+2]),
			offset: binary.BigEndian.Uint32(cidx[i+2 : i+6]),
		}
		bn++
	}
	files := map[uint16]*os.File{}
	for _, p := range index {
		if _, ok := files[p.file]; ok {
			continue
		}
		f, err := os.Open(fmt.Sprintf("%s/receipts.%04d.cdat", dir, p.file))
		check(err)
		files[p.file] = f
	}
	var res [][]byte
	for i, bn := uint(0), start; i < count; i++ {
		cur, next := index[bn], index[bn+1]

		var buf []byte
		switch {
		case cur.file != next.file:
			buf = make([]byte, next.offset)
			_, err = files[next.file].ReadAt(buf, 0)
			check(err)
		case cur.file == next.file:
			buf = make([]byte, next.offset-cur.offset)
			_, err = files[cur.file].ReadAt(buf, int64(cur.offset))
			check(err)
		}
		decoded, err := snappy.Decode(nil, buf)
		check(err)
		res = append(res, decoded)
		bn++
	}
	return res
}
