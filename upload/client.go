package upload

import (
	"net/http"

	"github.com/bleenco/go-resumable"
)

func NewClient(url, filepath string) *resumable.Resumable {

	chunksize := int(1 * (1 << 20)) // 1MB

	return resumable.New(url, filepath, &http.Client{}, chunksize, false)
}

func SendData(client *resumable.Resumable) {
	client.Init()
	client.Start()
	resumable.WG.Wait()
}
