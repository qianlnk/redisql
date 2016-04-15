%{
#include "stdio.h"
#include "stdlib.h"
#include "string.h"
#include "syntaxtree.h"

//int yydebug = 1;

struct stnode * create_non_terminal_node(char * pszName);

%}

%union{
	struct stnode * node_ptr;
}

%token <node_ptr> CREATE

%token <node_ptr> TABLE

%token <node_ptr> STRING

%token <node_ptr> INT

%token <node_ptr> FLOAT

%type <node_ptr> table

%type <node_ptr> column

%token <node_ptr> INTVAL

%token <node_ptr> STRINGVAL

%token <node_ptr> FLOATVAL

%token <node_ptr> NAME

%token <node_ptr> '('

%token <node_ptr> ')'

%token <node_ptr> '.'

%token <node_ptr> ','

%type <node_ptr> data_type

%type <node_ptr> base_table_element_commalist

%type <node_ptr> base_table_element

%type <node_ptr> column_def

%type <node_ptr> base_table_def

%token <node_ptr> QUIT

%token <node_ptr> INSERT

%token <node_ptr> INTO

%type <node_ptr> insert_statement

%type <node_ptr> opt_column_commalist

%type <node_ptr> values_or_query_spec

%type <node_ptr> column_commalist

%token <node_ptr> VALUES

%type <node_ptr> insert_atom_commalist

%type <node_ptr> insert_atom

%type <node_ptr> atom

%token <node_ptr> NULLX

%type <node_ptr> literal

%token <node_ptr> DELETE

%token <node_ptr> FROM

%type <node_ptr> delete_statement

%type <node_ptr> opt_where_clause

%type <node_ptr> where_clause

%token <node_ptr> WHERE

%type <node_ptr> search_condition

%token <node_ptr> OR

%token <node_ptr> AND

%token <node_ptr> NOT

%type <node_ptr> predicate

%type <node_ptr> comparison_predicate

%type <node_ptr> scalar_exp

%token <node_ptr> COMPARISON

%token <node_ptr> '+'

%token <node_ptr> '-'

%token <node_ptr> '*'

%token <node_ptr> '/'

%type <node_ptr> column_ref

%token <node_ptr> UPDATE

%token <node_ptr> SET

%type <node_ptr> assignment_commalist

%type <node_ptr> assignment

%type <node_ptr> update_statement

%type <node_ptr> select_statement

%token <node_ptr> SELECT

%type <node_ptr> selection

%type <node_ptr> table_exp

%type <node_ptr> scalar_exp_commalist

%type <node_ptr> from_clause

%type <node_ptr> table_ref_commalist

%type <node_ptr> table_ref

%token <node_ptr> DROP

%type <node_ptr> drop_table_statement

%token <node_ptr> HELP

%type <node_ptr> help_statement

%start sql

%%

sql: base_table_def
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_CREATE_TABLE;
	}
	|
	 QUIT
	 {
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_QUIT;			
	}
	|insert_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_INSERT;
	}
	|delete_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_DELETE;		
	}	
	|update_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_UPDATE;		
	}
	|select_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_SELECT;		
	}
	|drop_table_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_DROP_TABLE;		
	}
	|help_statement
	{
		struct stnode * p = create_non_terminal_node("sql");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		syntax_tree_ptr = p;

		sql_action = SAT_HELP;		
	}
	;

help_statement:HELP
	{
		struct stnode * p = create_non_terminal_node("help_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
		
	}
	;

base_table_def:

	CREATE TABLE table '(' base_table_element_commalist ')'
	{
		struct stnode * p = create_non_terminal_node("base_table_def");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $4))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $5))
		{
			printf("error:append_child\n");

			return 1;
		}	

		if(!append_child(p, $6))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;

base_table_element_commalist:
	base_table_element
	{
		struct stnode * p = create_non_terminal_node("base_table_element_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		$$ = p;
	}
	| base_table_element_commalist ',' base_table_element
	{
		struct stnode * p = create_non_terminal_node("base_table_element_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}	
		
		$$ = p;
	}
	;

base_table_element:
	column_def
	{
		struct stnode * p = create_non_terminal_node("base_table_element");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		$$ = p;
	}
	;

column_def:
	column data_type
	{
		struct stnode * p = create_non_terminal_node("column_def");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}
		
		$$ = p;
	}
	;
	
table:
	NAME
	{	struct stnode * p = create_non_terminal_node("table");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	|
	NAME '.' NAME
	{
		struct stnode * p = create_non_terminal_node("table");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

column:
	NAME 
	{	struct stnode * p = create_non_terminal_node("column");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

data_type:
	  STRING 
		{
			struct stnode * p = create_non_terminal_node("data_type");

			if(!p)
			{
				printf("error:create_non_terminal_node\n");

				return 1;
			}

			if(!append_child(p, $1))
			{
				printf("error:append_child\n");

				return 1;
			}

			$$ = p;
		}	
	| INT
		{
			struct stnode * p = create_non_terminal_node("data_type");

			if(!p)
			{
				printf("error:create_non_terminal_node\n");

				return 1;
			}

			if(!append_child(p, $1))
			{
				printf("error:append_child\n");

				return 1;
			}


			$$ = p;		
		}
	| FLOAT
		{
			struct stnode * p = create_non_terminal_node("data_type");

			if(!p)
			{
				printf("error:create_non_terminal_node\n");

				return 1;
			}

			if(!append_child(p, $1))
			{
				printf("error:append_child\n");

				return 1;
			}

			$$ = p;
		}
	;

insert_statement:
	INSERT INTO table opt_column_commalist values_or_query_spec
	{
		struct stnode * p = create_non_terminal_node("insert_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}			

		if(!append_child(p, $4))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $5))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;

opt_column_commalist:
	/*空值*/
	{
		struct stnode * p = create_non_terminal_node("opt_column_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		$$ = p;
	}
	| '(' column_commalist ')'
	{
		struct stnode * p = create_non_terminal_node("opt_column_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

column_commalist:
	column
	{
		struct stnode * p = create_non_terminal_node("column_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| column_commalist','column
	{
		struct stnode * p = create_non_terminal_node("column_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}
		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

values_or_query_spec:
	VALUES'('insert_atom_commalist')'
	{
		struct stnode * p = create_non_terminal_node("values_or_query_spec");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}
		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}
		if(!append_child(p, $4))
		{
			printf("error:append_child\n");

			return 1;
		}		

		$$ = p;
		
	}
	;
insert_atom_commalist:
	insert_atom
	{
		struct stnode * p = create_non_terminal_node("insert_atom_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	| insert_atom_commalist','insert_atom
	{
		struct stnode * p = create_non_terminal_node("insert_atom_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;

insert_atom:
	atom
	{
		struct stnode * p = create_non_terminal_node("insert_atom");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;	
	}
	|NULLX
	{
		struct stnode * p = create_non_terminal_node("insert_atom");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
atom: literal
	{
		struct stnode * p = create_non_terminal_node("atom");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;	
	}
	;
literal:
	STRINGVAL
	{
		struct stnode * p = create_non_terminal_node("literal");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;	
	}
	|INTVAL
	{
		struct stnode * p = create_non_terminal_node("literal");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}	

		$$ = p;
	}
	|FLOATVAL
	{
		struct stnode * p = create_non_terminal_node("literal");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

delete_statement:
	DELETE FROM table opt_where_clause
	{
		struct stnode * p = create_non_terminal_node("delete_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}		
		
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $4))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

opt_where_clause:
	/*空值*/
	{
		struct stnode * p = create_non_terminal_node("opt_where_clause");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		$$ = p;
	}
	|where_clause
	{
		struct stnode * p = create_non_terminal_node("opt_where_clause");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;

where_clause:
	WHERE search_condition
	{		
		struct stnode * p = create_non_terminal_node("where_clause");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
search_condition:
	search_condition OR search_condition
	{
		struct stnode * p = create_non_terminal_node("search_condition");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
		
	}
	| search_condition AND search_condition
	{
		struct stnode * p = create_non_terminal_node("search_condition");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	| NOT search_condition
	{
		struct stnode * p = create_non_terminal_node("search_condition");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}		

		$$ = p;
	}
	| '(' search_condition ')'
	{
		struct stnode * p = create_non_terminal_node("search_condition");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| predicate
	{
		struct stnode * p = create_non_terminal_node("search_condition");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
predicate:
	comparison_predicate
	{
		struct stnode * p = create_non_terminal_node("predicate");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

comparison_predicate:
	scalar_exp COMPARISON scalar_exp
	{
		struct stnode * p = create_non_terminal_node("comparison_predicate");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;

scalar_exp:
	scalar_exp'+'scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| scalar_exp'-'scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| scalar_exp'*'scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| scalar_exp'/'scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| '('scalar_exp')'
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| atom/*常量*/
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	| column_ref/*变量*/
	{
		struct stnode * p = create_non_terminal_node("scalar_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
column_ref:
	NAME
	{
		struct stnode * p = create_non_terminal_node("column_ref");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
		
	}
	| NAME'.'NAME
	{
		struct stnode * p = create_non_terminal_node("column_ref");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}
	
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
update_statement:
	UPDATE table SET assignment_commalist opt_where_clause
	{
		struct stnode * p = create_non_terminal_node("update_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}
	
		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $4))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $5))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
assignment_commalist:
	assignment
	{
		struct stnode * p = create_non_terminal_node("assignment_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	|assignment_commalist','assignment
	{
		struct stnode * p = create_non_terminal_node("assignment_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
assignment:
	column COMPARISON scalar_exp
	{
		struct stnode * p = create_non_terminal_node("assignment");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if($2 == NULL)
		{
			return 1;
		}

		if(strcmp($2->m_strName, "="))
		{
			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	|column COMPARISON NULLX
	{
		struct stnode * p = create_non_terminal_node("assignment");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if($2 == NULL)
		{
			return 1;
		}

		if(strcmp($2->m_strName, "="))
		{
			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;
	}
	;
select_statement:
	SELECT selection table_exp
	{
		struct stnode * p = create_non_terminal_node("select_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}		

		$$ = p;
	}
	;
selection:
	scalar_exp_commalist
	{
		struct stnode * p = create_non_terminal_node("selection");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	| '*'
	{
		struct stnode * p = create_non_terminal_node("selection");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
scalar_exp_commalist:
	scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	| scalar_exp_commalist','scalar_exp
	{
		struct stnode * p = create_non_terminal_node("scalar_exp_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;	
	}
	;
table_exp:
	from_clause
	opt_where_clause
	{
		struct stnode * p = create_non_terminal_node("table_exp");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
from_clause:
	FROM table_ref_commalist
	{
		struct stnode * p = create_non_terminal_node("from_clause");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
table_ref_commalist:
	table_ref
	{
		struct stnode * p = create_non_terminal_node("table_ref_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	| table_ref_commalist','table_ref
	{
		struct stnode * p = create_non_terminal_node("table_ref_commalist");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
table_ref:
	table
	{
		struct stnode * p = create_non_terminal_node("table_ref");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;
drop_table_statement:
	DROP TABLE table
	{
		struct stnode * p = create_non_terminal_node("drop_table_statement");

		if(!p)
		{
			printf("error:create_non_terminal_node\n");

			return 1;
		}

		if(!append_child(p, $1))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $2))
		{
			printf("error:append_child\n");

			return 1;
		}

		if(!append_child(p, $3))
		{
			printf("error:append_child\n");

			return 1;
		}

		$$ = p;		
	}
	;

%%

//创建非终端结点

struct stnode * create_non_terminal_node(char * pszName)
{
	if(!pszName)
	{
		return NULL;
	}	
	
	struct stnode * p = malloc_node();

	char * pBuf = NULL;

	if(!p)
	{
		printf("创建结点失败\n");

		return NULL;
	}

	//名称

	p->m_nType = BRANCH;

	int nLen = strlen(pszName);

	pBuf = 	(char * )malloc(nLen + 1);

	strcpy(pBuf, pszName);

	pBuf[nLen] = 0;

	p->m_strName = pBuf;	

	return p;
}

yyerror(char * s)
{
	printf("%s\n", s);
}

typedef YY_BUFFER_STATE;

int sql_parse(const char * sql)
{
	if(!sql)
	{
		printf("没有输入SQL语句\n");
	}

	int len = strlen(sql);

	YY_BUFFER_STATE state = yy_scan_bytes(sql, len);

	yy_switch_to_buffer(state);

	int n = yyparse();

	yy_delete_buffer(state);

	return n;
}