package calc

import "golang.org/x/net/context"

type EntityCalc struct{}

func (*EntityCalc) Percent(ctx context.Context, req *ReqPercent) (*ResPercent, error) {
	resVals := []float32{}
	for _, v := range req.Val {
		resVals = append(resVals, v*req.Percent/100)
	}
	return &ResPercent{Val: resVals}, nil
}
