package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(password(input()))
	part2(input())
}

func password(s string) string {

	pass := []string{}
	prefix := []byte{0, 0}
	i := 0
	for {
		if len(pass) == 8 {
			break
		}
		str := s + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, str)
		sum := h.Sum(nil)
		if bytes.Compare(sum[:2], prefix) == 0 && (sum[2]>>4) == 0 {
			fmt.Printf("%x\n", sum)
			char := fmt.Sprintf("%x", sum[2])
			pass = append(pass, char)
		}
		i++
	}
	return fmt.Sprintf("%s", strings.Join(pass, ""))
}

func part2(s string) {

	pass := [8]string{}
	prefix := []byte{0, 0}
	i := 0
	seen := map[byte]struct{}{}
	count := 0
	for {
		if count == 8 {
			break
		}
		str := s + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, str)
		sum := h.Sum(nil)
		if bytes.Compare(sum[:2], prefix) == 0 && (sum[2]>>4) == 0 {
			if sum[2] <= 7 {
				if _, ok := seen[sum[2]]; !ok {
					seventh := fmt.Sprintf("%x", sum[3]>>4)
					pass[sum[2]] = seventh
					seen[sum[2]] = struct{}{}
					count++
				}
			}
		}
		i++
	}
	fmt.Printf("%s\n", strings.Join(pass[:], ""))
}
func input() string {
	return `wtnhxymk`
}
