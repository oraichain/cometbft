package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/tendermint/tendermint/types"
)

func testRamUsage() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	_, err = types.GenesisDocFromFile(path.Join(dir, "/.oraid-fork-backup/config/genesis.json"))
	if err != nil {
		panic(err)
	}
	// genDoc.AppState = json.RawMessage([]byte("{}"))
}

func main() {
	testRamUsage()
	fmt.Println("after reading genesis doc")
	time.Sleep(1 * time.Minute)
}
