package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func compileContract(sol, contract string) {
	// solc --abi name.sol --overwrite -o name
	output, err := exec.Command(
		"solc",
		"--abi",
		fmt.Sprintf("%s.sol", sol),
		"--overwrite",
		"-o",
		fmt.Sprintf("%s", sol),
	).Output()
	if err != nil {
		panic(string(output))
	}

	// solc --bin name.sol --overwrite -o name
	output, err = exec.Command(
		"solc",
		"--bin", fmt.Sprintf("%s.sol", sol),
		"--overwrite",
		"-o", sol,
	).Output()
	if err != nil {
		panic(string(output))
	}

	// abigen --bin=name/name.bin --abi=name/name.abi --pkg=name --out=name/name.go
	output, err = exec.Command(
		"abigen",
		fmt.Sprintf("--bin=%s/%s.bin", sol, contract),
		fmt.Sprintf("--abi=%s/%s.abi", sol, contract),
		fmt.Sprintf("--pkg=%s", sol),
		fmt.Sprintf("--out=%s/%s.go", sol, strings.ToLower(sol)),
	).Output()
	if err != nil {
		panic(string(output))
	}
}

// Generates contracts
func main() {
	sol := flag.String("sol", "", "Name of the solidity file (without the extension).")
	contract := flag.String("contract", "", "Name of contract in solidity file.")

	flag.Parse()
	compileContract(*sol, *contract)
}
