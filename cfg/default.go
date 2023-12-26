package cfg

import pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"

var DEFAULT = &pb.Namespaces{
	Namespaces: []*pb.Namespace{
		HTML,
		SVG,
		MathML,
	},
}
