package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PolarBearAndrew/go-s3-viewer/downloader"
)

type S3ViewerServer struct {
	port     string
	s3Reader *downloader.S3Reader
}

type S3ViewerServConf struct {
	Port   string
	Bucket string
	Region string
}

func NewS3ViewerServer(conf S3ViewerServConf) *S3ViewerServer {
	return &S3ViewerServer{
		port:     conf.Port,
		s3Reader: downloader.NewS3Reader(conf.Region, conf.Bucket),
	}
}

func renderErr(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	_, _ = fmt.Fprintf(w, "err %v", err)
}

func (serv *S3ViewerServer) showTheObject(w http.ResponseWriter, req *http.Request) {

	body, err := serv.s3Reader.GetObject(req.URL.Path[1:])

	if err != nil {
		renderErr(w, err)
		return
	}

	bs, err := ioutil.ReadAll(*body)

	if err != nil {
		renderErr(w, err)
		return
	}

	_, _ = w.Write(bs)
}

func (serv *S3ViewerServer) Listen() {
	http.HandleFunc("/", serv.showTheObject)
	_ = http.ListenAndServe(serv.port, nil)
}
