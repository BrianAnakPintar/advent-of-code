package day3

import (
	"strconv"
	"unicode"
)

type ParserState int

const (
  idle    ParserState = iota
  m_state
  u_state
  l_state
  open_bracket_state
  m1_state
  m2_state
  m3_state
  comma_state
  n1_state
  n2_state
  n3_state
)

type MulParser struct {
  state ParserState
  val_one int
  val_one_str []rune
  val_two int
  val_two_str []rune
  sum int
}

func NewMulParser() *MulParser {
  parser := MulParser{state: idle}
  parser.val_one = 0
  parser.val_two = 0
  parser.val_one_str = make([]rune, 0)
  parser.val_two_str = make([]rune, 0)
  return &parser
}

func (this *MulParser) ParseVal(char rune) {
  switch this.state {
  case idle:
  this.IdleState(char)
  case m_state:
  this.MState(char)
  case u_state:
  this.UState(char)
  case l_state:
  this.LState(char)
  case open_bracket_state:
  this.OpenBracketState(char)
  case m1_state:
  this.M1State(char)
  case m2_state:
  this.M2State(char)
  case m3_state:
  this.M3State(char)
  case comma_state:
  this.CommaState(char)
  case n1_state:
  this.N1State(char)
  case n2_state:
  this.N2State(char)
  case n3_state:
  this.N3State(char)
  }
}

// INVARIANT: Each State function is gurranteed to be in the state they represent.

func (this *MulParser) IdleState(char rune) {
  if char == 'm' {
    this.state = m_state
  }
  this.val_one = 0
  this.val_two = 0
  this.val_one_str = make([]rune, 0)
  this.val_two_str = make([]rune, 0)
}

func (this *MulParser) MState(char rune) {
  if char == 'u' {
    this.state = u_state
  } else {
    this.state = idle
  }
}

func (this *MulParser) UState(char rune) {
  if char == 'l' {
    this.state = l_state
  } else {
    this.state = idle
  }
}

func (this *MulParser) LState(char rune) {
  if char == '(' {
    this.state = open_bracket_state
  } else {
    this.state = idle
  }
}

func (this *MulParser) OpenBracketState(char rune) {
  if unicode.IsDigit(char) {
    this.state = m1_state 
    this.val_one_str = append(this.val_one_str, char)
  } else {
    this.state = idle
  }
}

func (this *MulParser) M1State(char rune) {
  if unicode.IsDigit(char) {
    this.state = m2_state
    this.val_one_str = append(this.val_one_str, char)
  } else if char == ',' {
    this.state = comma_state
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle
  }
}

func (this *MulParser) M2State(char rune) {
  if unicode.IsDigit(char) {
    this.state = m3_state
    this.val_one_str = append(this.val_one_str, char)
  } else if char == ',' {
    this.state = comma_state
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle
  }
}

func (this *MulParser) M3State(char rune) {
  if char == ',' {
    this.state = comma_state
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle
  }
}

func (this *MulParser) CommaState(char rune) {
  if unicode.IsDigit(char) {
    this.state = n1_state 
    this.val_two_str = append(this.val_two_str, char)
  } else {
    this.state = idle
  }
}

func (this *MulParser) N1State(char rune) {
  if unicode.IsDigit(char) {
    this.state = n2_state
    this.val_two_str = append(this.val_two_str, char)
  } else if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle
  }
}

func (this *MulParser) N2State(char rune) {
  if unicode.IsDigit(char) {
    this.state = n3_state
    this.val_two_str = append(this.val_two_str, char)
  } else if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle
  }
}

func (this *MulParser) N3State(char rune) {
  if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle
  }
}

func (this *MulParser) HandleCloseBracket() {
  this.sum += this.val_one * this.val_two
  this.state = idle
}
