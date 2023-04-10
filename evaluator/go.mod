module monkey/evaluator

go 1.20

replace monkey/lexer => ../lexer

replace monkey/parser => ../parser

replace monkey/object => ../object

replace monkey/token => ../token

replace monkey/ast => ../ast

require (
	monkey/lexer v0.0.0-00010101000000-000000000000
	monkey/object v0.0.0-00010101000000-000000000000
	monkey/parser v0.0.0-00010101000000-000000000000
)

require (
	monkey/ast v0.0.0-00010101000000-000000000000 // indirect
	monkey/token v0.0.0-00010101000000-000000000000 // indirect
)
