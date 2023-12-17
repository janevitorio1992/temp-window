var testCasesDecomposition = []struct {
	number   Input // given input as a decimal number
	base     int   // given base to calculate from
	expected []int //
}

func TestSmallestBasesWithEndBaseTooLow(t *testing.T) {
	for _, test := range testCasesBaseTooLow {
		observed, _ := test.number.smallesPalindBase(test.notHighEnoughEndBase)
		if observed == test.expected {
			t.Errorf("smallesPalindBase of %d returned %d, want 0",
				test.number, observed)
		}
	}
}

func TestCreate(t *testing.T) {
	preTestValidate(t)

//	net := network.New(t)
//	val := net.Validators[0]
//	ctx := val.ClientCtx
//	cli.ForceSkipConfirm = true

  	t.Run("Recipe", func(t *testing.T) {
		args := []string{"NewUser0", goodPLR}
		cmd := DevCreate()
		_, err := clitestutil.ExecTestCLICmd(ctx, cmd, args)
		assert.Nil(t, err)
	})
}
