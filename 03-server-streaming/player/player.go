package player

import (
	"echo/api"
	"io/ioutil"
	"log"
	"math"
)

const chunkSize = 200000

const testFilePath = "./player/music/nate.mp3"

func GetTestFileChunks() ([]api.GlobalRadioResponse, error) {
	var chunks []api.GlobalRadioResponse

	bytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	seek := 0
	fileSize := len(bytes)
	var idx int64
	for idx = 0; seek < fileSize; idx++ {
		newChunkSize := int(math.Min(float64(chunkSize), float64(fileSize-seek)))
		chunks = append(chunks, api.GlobalRadioResponse{
			ChunkIndex: idx,
			MusicBytes: bytes[seek : seek+newChunkSize],
		})
		seek += newChunkSize
	}

	return chunks, nil
}
