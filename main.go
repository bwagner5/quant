package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\tquant takes a value and converts it to an SI unit w/ the smallest whole number.")
		fmt.Fprintln(os.Stderr, "\t\t> quant 1068Mi -> 1.04Gi")
		fmt.Fprintln(os.Stderr, "\tYou can also choose the SI output to perform base-10 and base-2 unit conversions")
		fmt.Fprintln(os.Stderr, "\t\t> quant -si binary 1068M -> 1018.52Mi")
		fmt.Fprint(os.Stderr, "\tquant can read from arguments or stdin for piping\n\n")

		flag.PrintDefaults()
	}
	si := flag.String("si", "auto", "SI unit to use for output: binary, decimal, or auto which uses the input unit to determine the output")
	flag.Parse()
	input := strings.ReplaceAll(strings.TrimSpace(strings.Join(flag.Args(), "")), ",", "")
	if input == "" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		input = strings.ReplaceAll(strings.TrimSpace(string(stdin)), ",", "")
	}

	quantity := resource.MustParse(input)
	if *si == "binary" {
		quantity.Format = resource.BinarySI
	} else if *si == "decimal" {
		quantity.Format = resource.DecimalSI
	} else if _, err := strconv.Atoi(input); err == nil {
		quantity.Format = resource.BinarySI
	}
	fmt.Println(LargestUnit(quantity))
}

func LargestUnit(quantity resource.Quantity) string {
	var scaledValue float64
	var largestScale int
	if quantity.Format == resource.BinarySI {
		for _, scale := range []int{60, 50, 40, 30, 20, 10, 0, -10} {
			scaledValue = quantity.AsApproximateFloat64() / math.Pow(2, float64(scale))
			if scaledValue < 1 {
				continue
			}
			largestScale = scale
			break
		}
	} else {
		var bytez resource.Scale = 0
		for _, scale := range []resource.Scale{resource.Exa, resource.Peta, resource.Tera, resource.Giga, resource.Mega, resource.Kilo, bytez, resource.Milli} {
			scaledValue := quantity.AsApproximateFloat64() / math.Pow(10, float64(scale))
			if scaledValue < 1 {
				continue
			}
			largestScale = int(scale)
			break
		}
	}
	return fmt.Sprintf("%s%s", formatFloat(scaledValue), GetUnit(quantity, int(largestScale)))
}

func GetUnit(quantity resource.Quantity, scale int) string {
	// <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei
	//
	//	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)
	//
	// <decimalSI>       ::= m | "" | k | M | G | T | P | E
	if quantity.Format == resource.DecimalSI {
		switch scale {
		case 18:
			return "E"
		case 15:
			return "P"
		case 12:
			return "T"
		case 9:
			return "G"
		case 6:
			return "M"
		case 3:
			return "k"
		case -3:
			return "m"
		}
	} else if quantity.Format == resource.BinarySI {
		switch scale {
		case 60:
			return "Ei"
		case 50:
			return "Pi"
		case 40:
			return "Ti"
		case 30:
			return "Gi"
		case 20:
			return "Mi"
		case 10:
			return "Ki"
		}
	}
	return ""
}

func formatFloat(f float64) string {
	s := strconv.FormatFloat(f, 'f', 5, 64)
	parts := strings.Split(s, ".")
	if len(parts) == 1 {
		return s
	}
	reversed := reverse(parts[0])
	withCommas := ""
	for i, p := range reversed {
		if i%3 == 0 && i != 0 {
			withCommas += ","
		}
		withCommas += string(p)
	}
	s = strings.Join([]string{reverse(withCommas), parts[1]}, ".")
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
