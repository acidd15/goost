package goost

import (
	"testing"
)

func Test_String(t *testing.T) {
	a := NewString("1235")

	if !(a.Primitive() == "1235") {
		t.Error("Expected result is not matched.")
	}
}

func Test_Compare(t *testing.T) {
	a := NewString("1234")

	if !(a.Compare("123") == 1) {
		t.Error("Expected result is not matched.")
	}

	if !(a.Compare("12345") == -1) {
		t.Error("Expected result is not matched.")
	}

	if !(a.Compare("1234") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_CompareToIgnoreCase(t *testing.T) {
	a := NewString("AbCD")

	if !(a.CompareToIgnoreCase("abc") == 1) {
		t.Error("Expected result is not matched.")
	}

	if !(a.CompareToIgnoreCase("abCde") == -1) {
		t.Error("Expected result is not matched.")
	}

	if !(a.CompareToIgnoreCase("abCd") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_Equals(t *testing.T) {
	a := NewString("AbCD")

	if !(a.Equals(a) == true) {
		t.Error("Expected result is not matched.")
	}

	if !(a.Equals(NewString("AbCD")) == false) {
		t.Error("Expected result is not matched.")
	}
}

func Test_IsEmpty(t *testing.T) {
	a := NewString("")

	if !(a.IsEmpty() == true) {
		t.Error("Expected result is not matched.")
	}

	b := NewString("b")

	if !(b.IsEmpty() == false) {
		t.Error("Expected result is not matched.")
	}
}

func Test_Length(t *testing.T) {
	a := NewString("1")

	if !(a.Length() == 1) {
		t.Error("Expected result is not matched.")
	}

	b := NewString("22")

	if !(b.Length() == 2) {
		t.Error("Expected result is not matched.")
	}
}

func Test_Primitive(t *testing.T) {
	a := NewString("1")

	if !(a.Primitive() == "1") {
		t.Error("Expected result is not matched.")
	}
}

func Test_RuneAt(t *testing.T) {
	a := NewString("가나다라")
	if !(a.RuneAt(0) == '가') {
		t.Error("Expected result is not matched.")
	}
}

func Test_Split(t *testing.T) {
	a := NewString("가:나:다:라")
	s := a.Split(":")
	if !(s.Length() == 4 && s.At(0).Compare("가") == 0 && s.At(1).Compare("나") == 0 &&
		s.At(2).Compare("다") == 0 && s.At(3).Compare("라") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_Join(t *testing.T) {
	a := NewString("가:나:다:라")
	s := a.Split(":").Join(":")
	if !(s.Compare("가:나:다:라") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_ToLower(t *testing.T) {
	a := NewString("ABcDE")

	if !(a.ToLower().Compare("abcde") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_ToUpper(t *testing.T) {
	a := NewString("ABcDE")

	if !(a.ToUpper().Compare("ABCDE") == 0) {
		t.Error("Expected result is not matched.")
	}
}

func Test_TrimSpace(t *testing.T) {
	a := NewString(" ABcDE ")

	if !(a.TrimSpace().Compare("ABcDE") == 0) {
		t.Error("Expected result is not matched.")
	}
}