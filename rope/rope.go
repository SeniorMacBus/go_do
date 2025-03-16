package rope

import "fmt"

type Rope struct {
	Data   string
	Weight uint
	Left   *Rope
	Right  *Rope
}

type spliterror struct{}

func (e *spliterror) Error() string {
	return "There was an error during the split operation!"
}

func (r *Rope) calc_weight() {
	if r.Data != "" {
		r.Weight = uint(len(r.Data))
	} else {
		r.Weight = r.Left.Weight
		if r.Right != nil {
			r.Weight += r.Right.Weight
		}
	}
}

func Print(r *Rope) {
	if r == nil {
		return
	}

	Print(r.Left)
	if r.Data != "" {
		fmt.Print(r.Data)
	}
	Print(r.Right)
}

func ParseRope(r Rope, idx uint) {
	// TODO: finish this function
}

func CreateNewRope(data string, left *Rope, right *Rope) Rope {
	new_rope := Rope{Data: data, Left: left, Right: right}
	new_rope.calc_weight()

	return new_rope
}

func Concatenate(left *Rope, right *Rope) Rope {
	new_rope := CreateNewRope("", left, right)

	return new_rope
}

func CharAt(r Rope, idx uint) byte {
	if r.Data != "" {
		return r.Data[idx]
	}

	if idx < r.Weight {
		return CharAt(*r.Left, idx)
	} else {
		return CharAt(*r.Right, idx-r.Weight)
	}
}

func Substring(r Rope) {

}

func (r *Rope) Insert(to_insert Rope, idx uint) {
}

// TODO: finish this function
func Split(r *Rope, idx uint) (Rope, Rope, error) {
	if r == nil {
		left := CreateNewRope("", nil, nil)
		right := CreateNewRope("", nil, nil)
		err := &spliterror{}
		return left, right, err
	}
	if r.Data != "" {
		left := CreateNewRope(r.Data[:idx], nil, nil)
		right := CreateNewRope(r.Data[idx:], nil, nil)
		return left, right, nil
	}

	return CreateNewRope("", nil, nil), CreateNewRope("", nil, nil), nil
}
