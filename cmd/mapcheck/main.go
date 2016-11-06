// mapcheck
// Verify the tile layer data from a .tmx file. The XML might look like this:
//
//		<layer name="foo" width="12" height="24">
//			<data encoding="base64" compression="zlib">
//				eJxjZGBgYBzFo3gUj+IBwACO3gEh
//			</data>
//		</layer>
//
// When decoded and decompressed, this data turns into a vector of 1s which
// represent tiles. No structure data is encoded, so the game engine must read
// from the layer metadata to determine structure. All we can verify at this stage
// is that the number of tiles is correctly determined.
package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

var (
	mapData       string = "eJxjZGBgYBzFo3gUj+IBwACO3gEh"
	expectedTiles int    = 288 // 12x24
)

func verifyTileCount() {
	b := []byte(string(mapData))
	if _, err := base64.StdEncoding.Decode(b, b); err != nil {
		log.Fatalln("error:", err.Error())
	}
	r := bytes.NewReader(b)
	z, err := zlib.NewReader(r)
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	defer z.Close()
	markers := make([]uint32, 0)
	var next uint32
	for {
		err = binary.Read(z, binary.LittleEndian, &next)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln("error:", err.Error())
		}
		markers = append(markers, next)
	}
	if len(markers) == expectedTiles {
		fmt.Println("[PASS] Detected 288/288 tiles.")
	} else {
		fmt.Printf("[FAIL] Detected %d/288 tiles\n.", len(markers))
	}
}

func main() {
	verifyTileCount()
}
