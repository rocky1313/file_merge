package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
    inputDir := "./"
    outputDir := "../output"
    mergeCount := 1000

    files, _ := filepath.Glob(inputDir + "/*.bin")

    for i := 0; i < len(files); i += mergeCount {
        var readers []io.Reader
        for j := i; j < i+mergeCount && j < len(files); j++ {
            file, _ := os.Open(files[j])
            readers = append(readers, file)
        }

        outFilename := fmt.Sprintf("%s/hrng_data_%d.bin", outputDir, (i/mergeCount)+1)
        outFile, _ := os.Create(outFilename)
        defer outFile.Close()

        writer := io.MultiWriter(outFile)
        _, _ = io.Copy(writer, io.MultiReader(readers...))
    }
}
