package day3

import (
	"strconv"
	"unicode"
)

type ParserTwoState int

const (
  idle_2    ParserTwoState = iota
  m_state_2
  u_state_2
  l_state_2
  open_bracket_state_2
  m1_state_2
  m2_state_2
  m3_state_2
  comma_state_2
  n1_state_2
  n2_state_2
  n3_state_2

  d_state_2 // d
  o_state_2 // o
  do_open_bracket_state_2 // (
  do_close_bracket_state_2 // )
  n_state_2 // n
  apost_state_2 // '
  t_state_2 // t
  dont_open_bracket_state_2 // (
  dont_close_bracket_state_2 // )
)

type MulParserTwo struct {
  state ParserTwoState
  val_one int
  val_one_str []rune
  val_two int
  val_two_str []rune
  sum int

  enableParsing bool
}

func NewMulParserTwo() *MulParserTwo {
  parser := MulParserTwo{state: idle_2}
  parser.val_one = 0
  parser.val_two = 0
  parser.val_one_str = make([]rune, 0)
  parser.val_two_str = make([]rune, 0)
  parser.enableParsing = true
  return &parser
}

func (this *MulParserTwo) ParseVal(char rune) {
  // I'll handle it here because it's convenient to add this to all states.
  if (char == 'd') {
    this.state = d_state_2
    return
  }
  // Special case for "do and don't"
  switch this.state {

  }
  switch this.state {
  case idle_2:
  this.IdleState(char)
  case m_state_2:
  this.MState(char)
  case u_state_2:
  this.UState(char)
  case l_state_2:
  this.LState(char)
  case open_bracket_state_2:
  this.OpenBracketState(char)
  case m1_state_2:
  this.M1State(char)
  case m2_state_2:
  this.M2State(char)
  case m3_state_2:
  this.M3State(char)
  case comma_state_2:
  this.CommaState(char)
  case n1_state_2:
  this.N1State(char)
  case n2_state_2:
  this.N2State(char)
  case n3_state_2:
  this.N3State(char)

  // For 
  case d_state_2:
  this.DState(char)
  case o_state_2:
  this.OState(char)
  case do_open_bracket_state_2:
  this.OpenBracketDoState(char)
  case n_state_2:
  this.NState(char)
  case apost_state_2:
  this.AposState(char)
  case t_state_2:
  this.TState(char)
  case dont_open_bracket_state_2:
  this.OpenBracketDontState(char)
  }
}

// INVARIANT: Each State function is gurranteed to be in the state they represent.

func (this *MulParserTwo) IdleState(char rune) {
  if (!this.enableParsing) {
    return
  }
  if char == 'm' {
    this.state = m_state_2
  }
  this.val_one = 0
  this.val_two = 0
  this.val_one_str = make([]rune, 0)
  this.val_two_str = make([]rune, 0)
}

func (this *MulParserTwo) MState(char rune) {
  if char == 'u' {
    this.state = u_state_2
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) UState(char rune) {
  if char == 'l' {
    this.state = l_state_2
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) LState(char rune) {
  if char == '(' {
    this.state = open_bracket_state_2
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) OpenBracketState(char rune) {
  if unicode.IsDigit(char) {
    this.state = m1_state_2
    this.val_one_str = append(this.val_one_str, char)
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) M1State(char rune) {
  if unicode.IsDigit(char) {
    this.state = m2_state_2
    this.val_one_str = append(this.val_one_str, char)
  } else if char == ',' {
    this.state = comma_state_2
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) M2State(char rune) {
  if unicode.IsDigit(char) {
    this.state = m3_state_2
    this.val_one_str = append(this.val_one_str, char)
  } else if char == ',' {
    this.state = comma_state_2
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) M3State(char rune) {
  if char == ',' {
    this.state = comma_state_2
    val, err := strconv.Atoi(string(this.val_one_str))
    if err != nil {
      panic(err)
    }
    this.val_one = val
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) CommaState(char rune) {
  if unicode.IsDigit(char) {
    this.state = n1_state_2 
    this.val_two_str = append(this.val_two_str, char)
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) N1State(char rune) {
  if unicode.IsDigit(char) {
    this.state = n2_state_2
    this.val_two_str = append(this.val_two_str, char)
  } else if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) N2State(char rune) {
  if unicode.IsDigit(char) {
    this.state = n3_state_2
    this.val_two_str = append(this.val_two_str, char)
  } else if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) N3State(char rune) {
  if char == ')' {
    val, err := strconv.Atoi(string(this.val_two_str))
    if err != nil {
      panic(err)
    }
    this.val_two = val

    this.HandleCloseBracket()
  } else {
    this.state = idle_2
  }
}

func (this *MulParserTwo) HandleCloseBracket() {
  this.sum += this.val_one * this.val_two
  this.state = idle_2
}

func (this *MulParserTwo) DState(char rune) {
  if char == 'o' {
    this.state = o_state_2
  } else {
    this.state = idle_2
    this.IdleState(char)
  }
}

func (this *MulParserTwo) OState(char rune) {
  if char == 'n' {
    this.state = n_state_2
  } else if char == '(' {
    this.state = do_open_bracket_state_2
  } else {
    this.state = idle_2
    this.IdleState(char)
  }
}

func (this *MulParserTwo) NState(char rune) {
  if char == '\'' {
    this.state = apost_state_2
  } else {
    this.state = idle_2
    this.IdleState(char)
  }
}

func (this *MulParserTwo) AposState(char rune) {
  if char == 't' {
    this.state = t_state_2
  } else {
    this.state = idle_2
    this.IdleState(char)
  }
}

func (this *MulParserTwo) TState(char rune) {
  if char == '(' {
    this.state = dont_open_bracket_state_2
  } else {
    this.state = idle_2
    this.IdleState(char)
  }
}

func (this *MulParserTwo) OpenBracketDontState(char rune) {
  if char == ')' {
    this.enableParsing = false
  } 
  this.state = idle_2
  this.IdleState(char)
}

func (this *MulParserTwo) OpenBracketDoState(char rune) {
  if (char == ')') {
    this.enableParsing = true
  }
  this.state = idle_2
  this.IdleState(char)
}

