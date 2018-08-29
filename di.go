package binu

import "github.com/clavoie/di"

func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewGobber, di.Singleton},
	}
}
