package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() float64 {
	reader := bufio.NewReader(os.Stdin)
	read, err := reader.ReadString('\n')
	if err != nil {
		err.Error()
	}
	read = strings.TrimSpace(read)
	input, err := strconv.ParseFloat(read, 64)
	if err != nil {
		err.Error()
	}
	return input
}
