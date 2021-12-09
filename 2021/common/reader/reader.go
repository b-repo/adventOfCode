package reader

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntFileAndSplitByLine(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func ReadIntFileAndSplitBy(fname, sep string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), sep)
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func ReadStringFileAndSplitByLine(fname string) (strs []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

func ReadByteFileAndSplitByLine(fname string) (bytes []byte, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	d := strings.Split(string(b), "\n")

	r := make([]byte, len(d))

	for j := range d {
		i, err := strconv.ParseInt(d[j], 2, 32)
		if err != nil {
			return nil, err
		}
		r = append(r, byte(i))
	}

	return r, nil
}

func Read2DStringSliceFile(fname string) ([][]string, error) {
	d, err := ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, err
	}

	r := make([][]string, 0)
	for i := range d {
		r = append(r, strings.Split(d[i], ""))
	}

	return r, nil
}

func Read2DIntSliceFile(fname string) ([][]int, error) {
	d, err := ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, err
	}

	r := make([][]int, 0)

	for i := range d {
		ys := strings.Split(d[i], "")
		col := make([]int, 0)
		for j := range ys {
			n, err := strconv.Atoi(ys[j])
			if err != nil {
				panic(err.Error())
			}

			col = append(col, n)
		}

		r = append(r, col)
	}

	return r, nil
}
