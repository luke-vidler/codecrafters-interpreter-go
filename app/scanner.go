package main

import "fmt"

type TokenType string

const (
	// Single-character tokens
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"

	// Special token
	EOF TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal string
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	// Add EOF token
	s.tokens = append(s.tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: "null",
	})

	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN, "null")
	case ')':
		s.addToken(RIGHT_PAREN, "null")
	default:
		// Ignore other characters for now
	}
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) addToken(tokenType TokenType, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
	})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, t.Literal)
}
