package student

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
)

var tt = []struct {
	name string
	req  *ReqProfile
	res  *ResProfile
}{
	{
		name: "success",
		req: &ReqProfile{
			Id: "0001",
		},
	},
}

func TestStudent(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, tc := range tt {
		abc := &mockSchoolStudent{}
		abc.On("getName", tc.req.Id).Return("Abhinav").Once()

		stud := &GRPCStudent{
			schoolStudent: abc,
		}

		t.Run(tc.name, func(t *testing.T) {
			a, _ := stud.Profile(ctx, tc.req)
			fmt.Printf("%+v", a)
		})
	}
}
