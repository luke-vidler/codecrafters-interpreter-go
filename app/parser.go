package main

// Parser implements a recursive descent parser
type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// Parse parses the tokens and returns an expression
func (p *Parser) Parse() Expr {
	return p.expression()
}

// expression parses an expression
func (p *Parser) expression() Expr {
	return p.term()
}

// term parses addition and subtraction expressions (+, -)
func (p *Parser) term() Expr {
	expr := p.factor()

	// Left-associative: keep consuming + and - operators
	for p.match(PLUS, MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

// factor parses multiplication and division expressions (*, /)
func (p *Parser) factor() Expr {
	expr := p.unary()

	// Left-associative: keep consuming * and / operators
	for p.match(STAR, SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

// unary parses unary expressions (!, -)
func (p *Parser) unary() Expr {
	// Check for unary operators
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary() // Right-associative, so we call unary() recursively
		return &Unary{Operator: operator, Right: right}
	}

	// No unary operator, move to primary
	return p.primary()
}

// primary parses primary expressions (literals and grouping)
func (p *Parser) primary() Expr {
	// Handle TRUE
	if p.match(TRUE) {
		return &Literal{Value: true}
	}

	// Handle FALSE
	if p.match(FALSE) {
		return &Literal{Value: false}
	}

	// Handle NIL
	if p.match(NIL) {
		return &Literal{Value: nil}
	}

	// Handle NUMBER
	if p.match(NUMBER) {
		// The previous token is the number we just matched
		return &Literal{Value: p.previous().Literal}
	}

	// Handle STRING
	if p.match(STRING) {
		return &Literal{Value: p.previous().Literal}
	}

	// Handle LEFT_PAREN - grouping expression
	if p.match(LEFT_PAREN) {
		expr := p.expression()
		// Consume the closing RIGHT_PAREN
		p.match(RIGHT_PAREN)
		return &Grouping{Expression: expr}
	}

	// If we get here, we couldn't parse anything
	return nil
}

// match checks if the current token matches any of the given types
func (p *Parser) match(types ...TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

// check returns true if the current token is of the given type
func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

// advance consumes the current token and returns it
func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// isAtEnd returns true if we're at the end of the token list
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

// peek returns the current token without consuming it
func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

// previous returns the most recently consumed token
func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}
