package cfg

import (
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
	"github.com/samber/lo"
)

var Default = &pb.Namespaces{
	Namespaces: []*pb.Namespace{
		HTML,
		SVG,
		MathML,
		ShoelaceExtensions,
	},
	Attributes: lo.Flatten([][]*pb.Attribute{
		DatastarExtensions,
	}),
}
