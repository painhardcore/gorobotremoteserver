package gorobotremoteserver

import (
	"net/http"
)

type Service struct {
	kws map[string]Userkeyword
}

func (sc *Service) add(kw Userkeyword) {
	sc.kws[kw.Name()] = kw
}

// GetKeywordNames is function needed to be compatible with Robot Interface
func (sc *Service) GetKeywordNames(r *http.Request, args *struct{}, reply *struct{ Reply []string }) error {
	names := make([]string, 0, len(sc.kws))
	for i := range sc.kws {
		names = append(names, sc.kws[i].Name())
	}
	reply.Reply = names
	return nil
}

// RunKeyword is function needed to be compatible with Robot Interface
func (sc *Service) RunKeyword(r *http.Request, args *struct {Name string; Args []string}, reply *struct{ Reply Result }) error {
	reply.Reply = sc.kws[args.Name].Run(args.Args)
	return nil
}
