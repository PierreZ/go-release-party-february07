//Let's rewrite echo "test" |  gzip --stdout | base64 in Go

func main() {

	pReader, pWriter := io.Pipe()                               // HL
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter) // HL
	gzipWriter := gzip.NewWriter(b64Writer)                     // HL

	go func() {
		gzipWriter.Write([]byte("These words will be compressed and push into pipe")) // HL
		gzipWriter.Close()
		b64Writer.Close()
		pWriter.Close()
	}()

	text, _ := ioutil.ReadAll(pReader)
	fmt.Printf("%s\n", text)
}
