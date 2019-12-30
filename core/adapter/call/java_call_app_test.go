package call

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter/identifier"
	"github.com/phodal/coca/core/models"
	"path/filepath"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(identifier.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaCallApp()
	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)

	g.Expect(len(callNodes)).To(Equal(1))
}

func TestJavaCallListener_EnterConstructorDeclaration(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/suggest/factory"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)
	g.Expect(len(callNodes[0].Methods)).To(Equal(3))
}

func getCallNodes(codePath string) []models.JClassNode {
	identifierApp := new(identifier.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaCallApp()

	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)
	return callNodes
}

func TestLambda_Express(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/lambda"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[1].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(methodMap["save"].MethodCalls[0].MethodName).To(Equal("of"))
	g.Expect(methodMap["findById"].MethodCalls[3].MethodName).To(Equal("toDomainModel"))
}

func TestInterface(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/interface"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Methods)).To(Equal(6))
	g.Expect(methodMap["count"].Name).To(Equal("count"))
}

func TestAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Annotations)).To(Equal(0))
	g.Expect(methodMap["macOsXPositiveTest"].Name).To(Equal("macOsXPositiveTest"))

	for _, call := range methodMap["macOsXPositiveTest"].MethodCalls {
		fmt.Println(call.Class)
	}

	g.Expect(true).To(Equal(true))
}

func Test_ShouldHaveOnlyOneAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/regression/CallAssertInClassTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["supportsEventType"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["genericListenerRawTypeTypeErasure"].Annotations)).To(Equal(1))
}

func Test_ShouldHaveOnlyOneAnnotationWithMultipleSame(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/regression/EnvironmentSystemIntegrationTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["setUp"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionMatchOr"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionMatchAnd"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionNoMatchAnd"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionNoMatchNone"].Annotations)).To(Equal(1))
}

func Test_CreatorAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["macOsXPositiveTest"].Annotations)).To(Equal(0))
}

func Test_ShouldGetMethodCreators(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	fmt.Println(methodMap["macOsXPositiveTest"])
	g.Expect(len(methodMap["macOsXPositiveTest"].Creators)).To(Equal(2))
}
