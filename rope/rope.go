package rope

type Rope struct {
	Data   string
	Weight uint
	Left   *Rope
	Right  *Rope
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

// TODO: finish this function
func Split(r Rope, idx uint) (Rope, Rope) {
	if r.Data != "" {

	}

	return left, right
}

func Insert(r Rope, idx uint, s string) Rope {
	left, right := Split(r, idx)
	insert_rope := CreateNewRope(s, nil, nil)
	left_attached := Concatenate(&left, &insert_rope)
	new_rope := Concatenate(&left_attached, right)

	return new_rope
}
