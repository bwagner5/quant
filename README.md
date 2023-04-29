# Quant

Quant is a numeric conversion tool. Its primary purpose is for human readable approximations of resource quantities. It can also help with base-2 and base-10 unit conversions. 

## Usage

```
> quant --help
Usage of quant:
	quant takes a value and converts it to an SI unit w/ the smallest whole number.
		> quant 1068Mi -> 1.04Gi
	You can also choose the SI output to perform base-10 and base-2 unit conversions
		> quant -si binary 1068M -> 1018.52Mi
	quant can read from arguments or stdin for piping

  -si string
    	SI unit to use for output: binary, decimal, or auto which uses the input unit to determine the output (default "auto")
```

## Examples:

```
> quant 1024
1.00Ki
> quant 1024Ki
1.00Mi
> quant -si decimal 1024Ki
1.05M
> echo "1,023,399T" | quant -si binary
908.96Pi
> quant 100,000,000,001m
100.00M
```