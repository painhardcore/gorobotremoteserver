package gorobotremoteserver

import (
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
)

const (
	Pass = "PASS"
	Fail = "FAIL"
	Port = 8270
)

type server struct {
	service *Service
}

// Register will add keyword by name.
// Can be used many times, only unique name will be used
func (s *server) Register(kw Userkeyword) {
	fmt.Printf("Registering %s\n",kw.Name())
	s.service.add(kw)
}

func (s *server) GetHandler() http.Handler {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	// Make compatible with Robot interface API
	xmlrpcCodec.RegisterAlias("get_keyword_names", "Service.GetKeywordNames")
	xmlrpcCodec.RegisterAlias("run_keyword", "Service.RunKeyword")
	// Not implemented yet since it's optional
	//xmlrpcCodec.RegisterAlias("get_keyword_arguments","Service.GetKeywordArguments")
	//xmlrpcCodec.RegisterAlias("get_keyword_documentation","Service.GetKeywordDocumentation")
	//xmlrpcCodec.RegisterAlias("get_keyword_types","Service.Get_keyword_types")
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(s.service, "")
	return RPC
}

// Start is used to make server on default port
func (s *server) Start() {
	RPC := s.GetHandler()
	http.Handle("/RPC2", RPC)
	log.Println("Starting XML-RPC server on localhost:8270/RPC2")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",Port), nil))
}

func New() *server {
	return &server{service:&Service{kws: make(map[string]Userkeyword)}}
}
