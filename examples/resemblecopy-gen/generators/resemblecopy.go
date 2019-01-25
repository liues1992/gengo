package generators

import (
	"fmt"
	"github.com/liues1992/gengo/args"
	"github.com/liues1992/gengo/examples/set-gen/sets"
	"github.com/liues1992/gengo/generator"
	"github.com/liues1992/gengo/namer"
	"github.com/liues1992/gengo/types"
	"io"
	"k8s.io/klog"
	"path/filepath"
	"strings"
)

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		klog.Fatalf("Failed loading boilerplate: %v", err)
	}

	inputs := sets.NewString(context.Inputs...)
	packages := generator.Packages{}
	header := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)

	for i := range inputs {
		klog.V(5).Infof("Considering pkg %q", i)
		pkg := context.Universe[i]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}

		klog.V(3).Infof("Package %q needs generation", i)
		path := pkg.Path
		// if the source path is within a /vendor/ directory (for example,
		// k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1), allow
		// generation to output to the proper relative path (under vendor).
		// Otherwise, the generator will create the file in the wrong location
		// in the output directory.
		// TODO: build a more fundamental concept in gengo for dealing with modifications
		// to vendored packages.
		if strings.HasPrefix(pkg.SourcePath, arguments.OutputBase) {
			expandedPath := strings.TrimPrefix(pkg.SourcePath, arguments.OutputBase)
			if strings.Contains(expandedPath, "/vendor/") {
				path = expandedPath
			}
		}
		packages = append(packages,
			&generator.DefaultPackage{
				PackageName: strings.Split(filepath.Base(pkg.Path), ".")[0],
				PackagePath: path,
				HeaderText:  header,
				GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
					return []generator.Generator{
						NewGenResembleCopy(arguments.OutputFileBaseName, pkg.Path),
					}
				},
				FilterFunc: func(c *generator.Context, t *types.Type) bool {
					return t.Name.Package == pkg.Path
				},
			})
	}
	return packages
}

func NewGenResembleCopy(sanitizedName, targetPackage string) generator.Generator {
	return &genResembleCopy{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		targetPackage: targetPackage,
		imports:       generator.NewImportTracker(),
	}
}

type genResembleCopy struct {
	generator.DefaultGen
	targetPackage string
	imports       namer.ImportTracker
}

func (*genResembleCopy) Name() string {
	panic("implement me")
}

func (*genResembleCopy) Filter(*generator.Context, *types.Type) bool {
	panic("implement me")
}

func (*genResembleCopy) Namers(*generator.Context) namer.NameSystems {
	panic("implement me")
}

func (*genResembleCopy) Init(*generator.Context, io.Writer) error {
	panic("implement me")
}

func (*genResembleCopy) Finalize(*generator.Context, io.Writer) error {
	panic("implement me")
}

func (*genResembleCopy) PackageVars(*generator.Context) []string {
	panic("implement me")
}

func (*genResembleCopy) PackageConsts(*generator.Context) []string {
	panic("implement me")
}

func (*genResembleCopy) GenerateType(*generator.Context, *types.Type, io.Writer) error {
	panic("implement me")
}

func (*genResembleCopy) Imports(*generator.Context) []string {
	panic("implement me")
}

func (*genResembleCopy) Filename() string {
	panic("implement me")
}

func (*genResembleCopy) FileType() string {
	panic("implement me")
}
