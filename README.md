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

## Installation:

```
$ brew tap bwagner5/wagner/quant
```

Packages, binaries, and archives are published for all major platforms (Mac amd64/arm64 & Linux amd64/arm64):

Debian / Ubuntu:

```
[[ `uname -m` == "aarch64" ]] && ARCH="arm64" || ARCH="amd64"
OS=`uname | tr '[:upper:]' '[:lower:]'`
wget https://github.com/bwagner5/quant/releases/download/v0.0.8/quant_0.0.8_${OS}_${ARCH}.deb
dpkg --install quant_0.0.8_linux_amd64.deb
quant --help
```

RedHat:

```
[[ `uname -m` == "aarch64" ]] && ARCH="arm64" || ARCH="amd64"
OS=`uname | tr '[:upper:]' '[:lower:]'`
rpm -i https://github.com/bwagner5/quant/releases/download/v0.0.8/quant_0.0.8_${OS}_${ARCH}.rpm
```

Download Binary Directly:

```
[[ `uname -m` == "aarch64" ]] && ARCH="arm64" || ARCH="amd64"
OS=`uname | tr '[:upper:]' '[:lower:]'`
wget -qO- https://github.com/bwagner5/quant/releases/download/v0.0.8/quant_0.0.8_${OS}_${ARCH}.tar.gz | tar xvz
chmod +x quant
```