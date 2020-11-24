package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	pb "google.golang.org/Tarea2SD/Client/Servicio"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func uploadFile(ctx context.Context, c pb.EstructuraCentralizadaClient, f string, name string, etx string) (stats pb.UploadStatus, err error) {
	fileToBeChunked := f
	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	const fileChunk = 256000 //250KB
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		re, _ := c.Subir(ctx, &pb.Chunk{Contenido: partBuffer, TotalChunks: int32(totalPartsNum), NumeroChunk: int32(i), Nombre: name, Ext: etx})
		log.Printf("M: %v", re.GetMensaje())
	}
	return

}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEstructuraCentralizadaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uploadFile(ctx, c, "./Book/Frankenstein-Mary_Shelley.pdf", "Frankenstein-Mary_Shelley", "pdf")
}
