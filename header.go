package main

const (
	TTYPE_IDENT int = iota
	TTYPE_PUNCT
	TTYPE_INT
	TTYPE_CHAR
	TTYPE_STRING
)

type Token struct {
	typ int
	v   struct { // wanna be Union
		ival  int
		sval  Cstring
		punct int
		c     byte
	}
}

const (
	AST_LITERAL int = iota
	AST_STRING
	AST_LVAR
	AST_GVAR
	AST_FUNCALL
	AST_FUNC
	AST_DECL
	AST_ARRAY_INIT
	AST_ADDR
	AST_DEREF
	AST_IF
	AST_TERNARY
	AST_FOR
	AST_RETURN
	AST_COMPOUND_STMT
	AST_STRUCT_REF
	PUNCT_EQ
	PUNCT_INC
	PUNCT_DEC
	PUNCT_LOGAND
	PUNCT_LOGOR
)

const (
	CTYPE_VOID int = iota
	CTYPE_INT
	CTYPE_CHAR
	CTYPE_ARRAY
	CTYPE_PTR
	CTYPE_STRUCT
)

type Ctype struct {
	typ    int
	ptr    *Ctype  // pointer or array
	size   int     // array
	name   Cstring // struct field
	tag    Cstring // struct
	fields []*Ctype
	offset int // struct
}

type Ast struct {
	typ   int
	ctype *Ctype
	// want to be "union"
	// Integer
	ival int
	// Char
	c byte
	// String
	str struct {
		val    Cstring
		slabel Cstring
	}
	// Local/Global variable
	variable struct {
		varname Cstring
		loff    int
		glabel  Cstring
	}
	// Binary operator
	binop struct {
		left  *Ast
		right *Ast
	}
	// Unary operator
	unary struct {
		operand *Ast
	}
	// Function call or function declaration
	fnc struct {
		fname     Cstring
		args      []*Ast
		params    []*Ast
		localvars []*Ast
		body      *Ast
	}
	// Declaration
	decl struct {
		declvar  *Ast
		declinit *Ast
	}
	// Array initializer
	array_initializer struct {
		arrayinit []*Ast
	}
	// If statement or ternary operator
	_if struct {
		cond *Ast
		then *Ast
		els  *Ast
	}
	// For statement
	_for struct {
		init *Ast
		cond *Ast
		step *Ast
		body *Ast
	}
	_return struct {
		retval *Ast
	}
	compound struct {
		stmts []*Ast
	}
	structref struct {
		struc *Ast
		field *Ctype
	}
}

type Env struct {
	vars    []*Ast
	next    *Env
	structs []*Ast
}

var EMPTY_ENV = Env{}
