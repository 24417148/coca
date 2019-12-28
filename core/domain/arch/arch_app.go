package arch

import (
	"fmt"
	"github.com/phodal/coca/core/domain/arch/tequila"
	"github.com/phodal/coca/core/models"
)

type ArchApp struct {
}

func NewArchApp() ArchApp {
	return *&ArchApp{}
}

func (a ArchApp) Analysis(deps []models.JClassNode, identifiersMap map[string]models.JIdentifier) tequila.FullGraph {
	fullGraph := tequila.FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*tequila.Relation),
	}

	for _, clz := range deps {
		if clz.Class == "Main" {
			continue
		}

		src := clz.Package + "." + clz.Class
		fullGraph.NodeList[src] = src

		if clz.Extend != "" {
			fmt.Println(clz.Extend)
			relation := &tequila.Relation{
				From:  src,
				To:    clz.Extend,
				Style: "\"solid\"",
			}

			fullGraph.RelationList[relation.From+"->"+relation.To] = relation
		}

		for _, method := range clz.Methods {
			if method.Name == "main" {
				continue
			}

			// TODO: add implements, extends support
			for _, call := range method.MethodCalls {
				dst := call.Package + "." + call.Class
				if _, ok := identifiersMap[dst]; ok {
					relation := &tequila.Relation{
						From:  src,
						To:    dst,
						Style: "\"solid\"",
					}

					fullGraph.RelationList[relation.From+"->"+relation.To] = relation
				}
			}
		}
	}

	return *&fullGraph
}
