/* A Bison parser, made by GNU Bison 2.3.  */

/* Skeleton implementation for Bison's Yacc-like parsers in C

   Copyright (C) 1984, 1989, 1990, 2000, 2001, 2002, 2003, 2004, 2005, 2006
   Free Software Foundation, Inc.

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2, or (at your option)
   any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program; if not, write to the Free Software
   Foundation, Inc., 51 Franklin Street, Fifth Floor,
   Boston, MA 02110-1301, USA.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.

   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* C LALR(1) parser skeleton written by Richard Stallman, by
   simplifying the original so-called "semantic" parser.  */

/* All symbols defined below should begin with yy or YY, to avoid
   infringing on user name space.  This should be done even for local
   variables, as they might otherwise be expanded by user macros.
   There are some unavoidable exceptions within include files to
   define necessary library symbols; they are noted "INFRINGES ON
   USER NAME SPACE" below.  */

/* Identify Bison output.  */
#define YYBISON 1

/* Bison version.  */
#define YYBISON_VERSION "2.3"

/* Skeleton name.  */
#define YYSKELETON_NAME "yacc.c"

/* Pure parsers.  */
#define YYPURE 0

/* Using locations.  */
#define YYLSP_NEEDED 0



/* Tokens.  */
#ifndef YYTOKENTYPE
# define YYTOKENTYPE
   /* Put the tokens into the symbol table, so that GDB and other debuggers
      know about them.  */
   enum yytokentype {
     USE = 258,
     SHOW = 259,
     DATABASES = 260,
     TABLES = 261,
     INDEX = 262,
     FROM = 263,
     DESC = 264,
     CREATE = 265,
     DATABASE = 266,
     TABLE = 267,
     ON = 268,
     INSERT = 269,
     INTO = 270,
     VALUES = 271,
     SELECT = 272,
     DISTINCT = 273,
     AS = 274,
     WHERE = 275,
     AND = 276,
     OR = 277,
     LIKE = 278,
     TOP = 279,
     LIMIT = 280,
     HELP = 281,
     EXIT = 282,
     NOT = 283,
     TOKEN_NULL = 284,
     UNIQUE = 285,
     PRIMARY = 286,
     KEY = 287,
     DEFAULT = 288,
     AUTO_INCREMENT = 289,
     COUNT = 290,
     SUM = 291,
     AVG = 292,
     MIN = 293,
     MAX = 294,
     JOIN = 295,
     CROSS = 296,
     INNER = 297,
     LEFT = 298,
     RIGHT = 299,
     FULL = 300,
     NATURAL = 301,
     OUTER = 302,
     CONCAT = 303,
     INT = 304,
     FLOAT = 305,
     DOUBLE = 306,
     CHAR = 307,
     VARCHAR = 308,
     TEXT = 309,
     DATE = 310,
     DATETIME = 311,
     NAME = 312,
     STRINGVAL = 313,
     COMPARISON = 314,
     INTVAL = 315,
     FLOATVAL = 316
   };
#endif
/* Tokens.  */
#define USE 258
#define SHOW 259
#define DATABASES 260
#define TABLES 261
#define INDEX 262
#define FROM 263
#define DESC 264
#define CREATE 265
#define DATABASE 266
#define TABLE 267
#define ON 268
#define INSERT 269
#define INTO 270
#define VALUES 271
#define SELECT 272
#define DISTINCT 273
#define AS 274
#define WHERE 275
#define AND 276
#define OR 277
#define LIKE 278
#define TOP 279
#define LIMIT 280
#define HELP 281
#define EXIT 282
#define NOT 283
#define TOKEN_NULL 284
#define UNIQUE 285
#define PRIMARY 286
#define KEY 287
#define DEFAULT 288
#define AUTO_INCREMENT 289
#define COUNT 290
#define SUM 291
#define AVG 292
#define MIN 293
#define MAX 294
#define JOIN 295
#define CROSS 296
#define INNER 297
#define LEFT 298
#define RIGHT 299
#define FULL 300
#define NATURAL 301
#define OUTER 302
#define CONCAT 303
#define INT 304
#define FLOAT 305
#define DOUBLE 306
#define CHAR 307
#define VARCHAR 308
#define TEXT 309
#define DATE 310
#define DATETIME 311
#define NAME 312
#define STRINGVAL 313
#define COMPARISON 314
#define INTVAL 315
#define FLOATVAL 316




/* Copy the first part of user declarations.  */
#line 1 "yacc_redisql.y"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "parse_redisql.h"

extern int yylex(void);
void yyerror(const char *s);


/* Enabling traces.  */
#ifndef YYDEBUG
# define YYDEBUG 1
#endif

/* Enabling verbose error messages.  */
#ifdef YYERROR_VERBOSE
# undef YYERROR_VERBOSE
# define YYERROR_VERBOSE 1
#else
# define YYERROR_VERBOSE 0
#endif

/* Enabling the token table.  */
#ifndef YYTOKEN_TABLE
# define YYTOKEN_TABLE 0
#endif

#if ! defined YYSTYPE && ! defined YYSTYPE_IS_DECLARED
typedef union YYSTYPE
#line 11 "yacc_redisql.y"
{
	char *key;
	char *strDateType;
	char *strDatabase;
	char *strTable;
	char *strColumn;
	char *strIndex;
	char *strOption;
	char *strVal;
	int nVal;
	double fVal;
	FieldAlias stFieldAlias;
}
/* Line 193 of yacc.c.  */
#line 242 "y.tab.c"
	YYSTYPE;
# define yystype YYSTYPE /* obsolescent; will be withdrawn */
# define YYSTYPE_IS_DECLARED 1
# define YYSTYPE_IS_TRIVIAL 1
#endif



/* Copy the second part of user declarations.  */


/* Line 216 of yacc.c.  */
#line 255 "y.tab.c"

#ifdef short
# undef short
#endif

#ifdef YYTYPE_UINT8
typedef YYTYPE_UINT8 yytype_uint8;
#else
typedef unsigned char yytype_uint8;
#endif

#ifdef YYTYPE_INT8
typedef YYTYPE_INT8 yytype_int8;
#elif (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
typedef signed char yytype_int8;
#else
typedef short int yytype_int8;
#endif

#ifdef YYTYPE_UINT16
typedef YYTYPE_UINT16 yytype_uint16;
#else
typedef unsigned short int yytype_uint16;
#endif

#ifdef YYTYPE_INT16
typedef YYTYPE_INT16 yytype_int16;
#else
typedef short int yytype_int16;
#endif

#ifndef YYSIZE_T
# ifdef __SIZE_TYPE__
#  define YYSIZE_T __SIZE_TYPE__
# elif defined size_t
#  define YYSIZE_T size_t
# elif ! defined YYSIZE_T && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
#  include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  define YYSIZE_T size_t
# else
#  define YYSIZE_T unsigned int
# endif
#endif

#define YYSIZE_MAXIMUM ((YYSIZE_T) -1)

#ifndef YY_
# if defined YYENABLE_NLS && YYENABLE_NLS
#  if ENABLE_NLS
#   include <libintl.h> /* INFRINGES ON USER NAME SPACE */
#   define YY_(msgid) dgettext ("bison-runtime", msgid)
#  endif
# endif
# ifndef YY_
#  define YY_(msgid) msgid
# endif
#endif

/* Suppress unused-variable warnings by "using" E.  */
#if ! defined lint || defined __GNUC__
# define YYUSE(e) ((void) (e))
#else
# define YYUSE(e) /* empty */
#endif

/* Identity function, used to suppress warnings about constant conditions.  */
#ifndef lint
# define YYID(n) (n)
#else
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static int
YYID (int i)
#else
static int
YYID (i)
    int i;
#endif
{
  return i;
}
#endif

#if ! defined yyoverflow || YYERROR_VERBOSE

/* The parser invokes alloca or malloc; define the necessary symbols.  */

# ifdef YYSTACK_USE_ALLOCA
#  if YYSTACK_USE_ALLOCA
#   ifdef __GNUC__
#    define YYSTACK_ALLOC __builtin_alloca
#   elif defined __BUILTIN_VA_ARG_INCR
#    include <alloca.h> /* INFRINGES ON USER NAME SPACE */
#   elif defined _AIX
#    define YYSTACK_ALLOC __alloca
#   elif defined _MSC_VER
#    include <malloc.h> /* INFRINGES ON USER NAME SPACE */
#    define alloca _alloca
#   else
#    define YYSTACK_ALLOC alloca
#    if ! defined _ALLOCA_H && ! defined _STDLIB_H && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
#     include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
#     ifndef _STDLIB_H
#      define _STDLIB_H 1
#     endif
#    endif
#   endif
#  endif
# endif

# ifdef YYSTACK_ALLOC
   /* Pacify GCC's `empty if-body' warning.  */
#  define YYSTACK_FREE(Ptr) do { /* empty */; } while (YYID (0))
#  ifndef YYSTACK_ALLOC_MAXIMUM
    /* The OS might guarantee only one guard page at the bottom of the stack,
       and a page size can be as small as 4096 bytes.  So we cannot safely
       invoke alloca (N) if N exceeds 4096.  Use a slightly smaller number
       to allow for a few compiler-allocated temporary stack slots.  */
#   define YYSTACK_ALLOC_MAXIMUM 4032 /* reasonable circa 2006 */
#  endif
# else
#  define YYSTACK_ALLOC YYMALLOC
#  define YYSTACK_FREE YYFREE
#  ifndef YYSTACK_ALLOC_MAXIMUM
#   define YYSTACK_ALLOC_MAXIMUM YYSIZE_MAXIMUM
#  endif
#  if (defined __cplusplus && ! defined _STDLIB_H \
       && ! ((defined YYMALLOC || defined malloc) \
	     && (defined YYFREE || defined free)))
#   include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
#   ifndef _STDLIB_H
#    define _STDLIB_H 1
#   endif
#  endif
#  ifndef YYMALLOC
#   define YYMALLOC malloc
#   if ! defined malloc && ! defined _STDLIB_H && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
void *malloc (YYSIZE_T); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
#  ifndef YYFREE
#   define YYFREE free
#   if ! defined free && ! defined _STDLIB_H && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
void free (void *); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
# endif
#endif /* ! defined yyoverflow || YYERROR_VERBOSE */


#if (! defined yyoverflow \
     && (! defined __cplusplus \
	 || (defined YYSTYPE_IS_TRIVIAL && YYSTYPE_IS_TRIVIAL)))

/* A type that is properly aligned for any stack member.  */
union yyalloc
{
  yytype_int16 yyss;
  YYSTYPE yyvs;
  };

/* The size of the maximum gap between one aligned stack and the next.  */
# define YYSTACK_GAP_MAXIMUM (sizeof (union yyalloc) - 1)

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */
# define YYSTACK_BYTES(N) \
     ((N) * (sizeof (yytype_int16) + sizeof (YYSTYPE)) \
      + YYSTACK_GAP_MAXIMUM)

/* Copy COUNT objects from FROM to TO.  The source and destination do
   not overlap.  */
# ifndef YYCOPY
#  if defined __GNUC__ && 1 < __GNUC__
#   define YYCOPY(To, From, Count) \
      __builtin_memcpy (To, From, (Count) * sizeof (*(From)))
#  else
#   define YYCOPY(To, From, Count)		\
      do					\
	{					\
	  YYSIZE_T yyi;				\
	  for (yyi = 0; yyi < (Count); yyi++)	\
	    (To)[yyi] = (From)[yyi];		\
	}					\
      while (YYID (0))
#  endif
# endif

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */
# define YYSTACK_RELOCATE(Stack)					\
    do									\
      {									\
	YYSIZE_T yynewbytes;						\
	YYCOPY (&yyptr->Stack, Stack, yysize);				\
	Stack = &yyptr->Stack;						\
	yynewbytes = yystacksize * sizeof (*Stack) + YYSTACK_GAP_MAXIMUM; \
	yyptr += yynewbytes / sizeof (*yyptr);				\
      }									\
    while (YYID (0))

#endif

/* YYFINAL -- State number of the termination state.  */
#define YYFINAL  33
/* YYLAST -- Last index in YYTABLE.  */
#define YYLAST   199

/* YYNTOKENS -- Number of terminals.  */
#define YYNTOKENS  71
/* YYNNTS -- Number of nonterminals.  */
#define YYNNTS  52
/* YYNRULES -- Number of rules.  */
#define YYNRULES  118
/* YYNRULES -- Number of states.  */
#define YYNSTATES  207

/* YYTRANSLATE(YYLEX) -- Bison symbol number corresponding to YYLEX.  */
#define YYUNDEFTOK  2
#define YYMAXUTOK   316

#define YYTRANSLATE(YYX)						\
  ((unsigned int) (YYX) <= YYMAXUTOK ? yytranslate[YYX] : YYUNDEFTOK)

/* YYTRANSLATE[YYLEX] -- Bison symbol number corresponding to YYLEX.  */
static const yytype_uint8 yytranslate[] =
{
       0,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
      59,    61,    66,    64,    62,    65,    63,    67,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,    68,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     1,     2,     3,     4,
       5,     6,     7,     8,     9,    10,    11,    12,    13,    14,
      15,    16,    17,    18,    19,    20,    21,    22,    23,    24,
      25,    26,    27,    28,    29,    30,    31,    32,    33,    34,
      35,    36,    37,    38,    39,    40,    41,    42,    43,    44,
      45,    46,    47,    48,    49,    50,    51,    52,    53,    54,
      55,    56,    57,    58,    60,    69,    70
};

#if YYDEBUG
/* YYPRHS[YYN] -- Index of the first RHS symbol of rule number YYN in
   YYRHS.  */
static const yytype_uint16 yyprhs[] =
{
       0,     0,     3,     5,     7,     9,    11,    13,    15,    19,
      21,    23,    25,    29,    33,    39,    43,    45,    47,    49,
      54,    62,    64,    68,    72,    74,    76,    78,    80,    82,
      84,    86,    88,    93,    95,    96,    98,   101,   104,   106,
     109,   112,   114,   124,   134,   138,   139,   141,   145,   147,
     151,   153,   155,   157,   167,   169,   170,   173,   178,   181,
     183,   184,   188,   192,   194,   198,   202,   206,   208,   212,
     215,   217,   219,   221,   223,   228,   230,   234,   236,   238,
     240,   242,   244,   246,   248,   250,   255,   259,   262,   264,
     266,   269,   272,   276,   280,   284,   287,   289,   290,   292,
     293,   296,   298,   300,   302,   304,   306,   307,   310,   311,
     313,   317,   321,   325,   327,   329,   332,   333,   338
};

/* YYRHS -- A `-1'-separated list of the rules' RHS.  */
static const yytype_int8 yyrhs[] =
{
      72,     0,    -1,    73,    -1,    74,    -1,    78,    -1,    79,
      -1,    89,    -1,    94,    -1,     3,   112,   116,    -1,    75,
      -1,    76,    -1,    77,    -1,     4,     5,   116,    -1,     4,
       6,   116,    -1,     4,     7,     8,   113,   116,    -1,     9,
     113,   116,    -1,    80,    -1,    81,    -1,    88,    -1,    10,
      11,   112,   116,    -1,    10,    12,   113,    59,    82,    61,
     116,    -1,    83,    -1,    82,    62,    83,    -1,   114,    84,
      85,    -1,    49,    -1,    50,    -1,    51,    -1,    52,    -1,
      53,    -1,    54,    -1,    55,    -1,    56,    -1,    84,    59,
      69,    61,    -1,    86,    -1,    -1,    87,    -1,    87,    86,
      -1,    28,    29,    -1,    30,    -1,    31,    32,    -1,    33,
      93,    -1,    34,    -1,    10,     7,   115,    13,   113,    59,
      91,    61,   116,    -1,    14,    15,   113,    90,    16,    59,
      92,    61,   116,    -1,    59,    91,    61,    -1,    -1,   114,
      -1,    91,    62,   114,    -1,    93,    -1,    92,    62,    93,
      -1,    69,    -1,    70,    -1,    58,    -1,    17,    95,    96,
       8,   105,   117,   121,   122,   116,    -1,    18,    -1,    -1,
      98,    97,    -1,    96,    62,    98,    97,    -1,    19,    57,
      -1,    57,    -1,    -1,    98,    64,    99,    -1,    98,    65,
      99,    -1,    99,    -1,    99,    66,   100,    -1,    99,    67,
     100,    -1,    99,    48,   100,    -1,   100,    -1,    59,    98,
      61,    -1,    65,   100,    -1,   101,    -1,    93,    -1,    29,
      -1,   102,    -1,   103,    59,    98,    61,    -1,   104,    -1,
     113,    63,   104,    -1,    35,    -1,    36,    -1,    37,    -1,
      38,    -1,    39,    -1,    66,    -1,   114,    -1,   106,    -1,
     105,   107,   106,   110,    -1,   108,   106,   110,    -1,   113,
      97,    -1,    62,    -1,    40,    -1,    41,    40,    -1,    42,
      40,    -1,    43,   109,    40,    -1,    44,   109,    40,    -1,
      45,   109,    40,    -1,    46,    40,    -1,    47,    -1,    -1,
     111,    -1,    -1,    13,   118,    -1,    57,    -1,    57,    -1,
      57,    -1,    57,    -1,    68,    -1,    -1,    20,   118,    -1,
      -1,   119,    -1,   118,   120,   119,    -1,    98,    60,    98,
      -1,    59,   118,    61,    -1,    21,    -1,    22,    -1,    24,
      69,    -1,    -1,    25,    69,    62,    69,    -1,    -1
};

/* YYRLINE[YYN] -- source line where rule number YYN was defined.  */
static const yytype_uint16 yyrline[] =
{
       0,    52,    52,    57,    61,    66,    70,    74,    81,    88,
      92,    96,   104,   111,   118,   125,   132,   137,   142,   149,
     156,   163,   167,   174,   181,   185,   189,   193,   197,   201,
     205,   209,   213,   220,   221,   225,   226,   230,   231,   232,
     233,   234,   238,   246,   253,   258,   264,   268,   276,   280,
     287,   296,   305,   315,   322,   323,   326,   339,   355,   359,
     364,   370,   371,   372,   379,   381,   383,   385,   392,   396,
     400,   407,   414,   418,   422,   429,   437,   448,   449,   450,
     451,   452,   456,   460,   467,   469,   471,   476,   483,   484,
     485,   486,   490,   491,   492,   493,   497,   498,   502,   507,
     513,   520,   527,   534,   541,   548,   553,   559,   564,   569,
     573,   584,   611,   621,   622,   626,   631,   635,   640
};
#endif

#if YYDEBUG || YYERROR_VERBOSE || YYTOKEN_TABLE
/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */
static const char *const yytname[] =
{
  "$end", "error", "$undefined", "USE", "SHOW", "DATABASES", "TABLES",
  "INDEX", "FROM", "DESC", "CREATE", "DATABASE", "TABLE", "ON", "INSERT",
  "INTO", "VALUES", "SELECT", "DISTINCT", "AS", "WHERE", "AND", "OR",
  "LIKE", "TOP", "LIMIT", "HELP", "EXIT", "NOT", "TOKEN_NULL", "UNIQUE",
  "PRIMARY", "KEY", "DEFAULT", "AUTO_INCREMENT", "COUNT", "SUM", "AVG",
  "MIN", "MAX", "JOIN", "CROSS", "INNER", "LEFT", "RIGHT", "FULL",
  "NATURAL", "OUTER", "CONCAT", "INT", "FLOAT", "DOUBLE", "CHAR",
  "VARCHAR", "TEXT", "DATE", "DATETIME", "NAME", "STRINGVAL", "'('",
  "COMPARISON", "')'", "','", "'.'", "'+'", "'-'", "'*'", "'/'", "';'",
  "INTVAL", "FLOATVAL", "$accept", "sql", "use", "show", "show_databases",
  "show_tables", "show_index", "desc", "create", "create_database",
  "create_table", "column_name_type_list", "column_name_type",
  "column_type", "opt_constraint_list", "constraint_list", "constraint",
  "create_index", "insert_into", "opt_column_name_list",
  "column_name_list", "value_list", "value", "select", "opt_distinct",
  "expression_list", "opt_alias", "expression", "mulexp", "primary",
  "term", "column_reference", "function_name", "column_name_or_star",
  "table_list", "table_def", "default_join", "join", "opt_outer",
  "opt_join_condition", "join_condition", "database_name", "table_name",
  "column_name", "index_name", "opt_semicolon", "opt_where_condition",
  "condition", "bool_term", "bool_op", "opt_top", "opt_limit", 0
};
#endif

# ifdef YYPRINT
/* YYTOKNUM[YYLEX-NUM] -- Internal token number corresponding to
   token YYLEX-NUM.  */
static const yytype_uint16 yytoknum[] =
{
       0,   256,   257,   258,   259,   260,   261,   262,   263,   264,
     265,   266,   267,   268,   269,   270,   271,   272,   273,   274,
     275,   276,   277,   278,   279,   280,   281,   282,   283,   284,
     285,   286,   287,   288,   289,   290,   291,   292,   293,   294,
     295,   296,   297,   298,   299,   300,   301,   302,   303,   304,
     305,   306,   307,   308,   309,   310,   311,   312,   313,    40,
     314,    41,    44,    46,    43,    45,    42,    47,    59,   315,
     316
};
# endif

/* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */
static const yytype_uint8 yyr1[] =
{
       0,    71,    72,    72,    72,    72,    72,    72,    73,    74,
      74,    74,    75,    76,    77,    78,    79,    79,    79,    80,
      81,    82,    82,    83,    84,    84,    84,    84,    84,    84,
      84,    84,    84,    85,    85,    86,    86,    87,    87,    87,
      87,    87,    88,    89,    90,    90,    91,    91,    92,    92,
      93,    93,    93,    94,    95,    95,    96,    96,    97,    97,
      97,    98,    98,    98,    99,    99,    99,    99,   100,   100,
     100,   101,   101,   101,   101,   102,   102,   103,   103,   103,
     103,   103,   104,   104,   105,   105,   105,   106,   107,   107,
     107,   107,   108,   108,   108,   108,   109,   109,   110,   110,
     111,   112,   113,   114,   115,   116,   116,   117,   117,   118,
     118,   119,   119,   120,   120,   121,   121,   122,   122
};

/* YYR2[YYN] -- Number of symbols composing right hand side of rule YYN.  */
static const yytype_uint8 yyr2[] =
{
       0,     2,     1,     1,     1,     1,     1,     1,     3,     1,
       1,     1,     3,     3,     5,     3,     1,     1,     1,     4,
       7,     1,     3,     3,     1,     1,     1,     1,     1,     1,
       1,     1,     4,     1,     0,     1,     2,     2,     1,     2,
       2,     1,     9,     9,     3,     0,     1,     3,     1,     3,
       1,     1,     1,     9,     1,     0,     2,     4,     2,     1,
       0,     3,     3,     1,     3,     3,     3,     1,     3,     2,
       1,     1,     1,     1,     4,     1,     3,     1,     1,     1,
       1,     1,     1,     1,     1,     4,     3,     2,     1,     1,
       2,     2,     3,     3,     3,     2,     1,     0,     1,     0,
       2,     1,     1,     1,     1,     1,     0,     2,     0,     1,
       3,     3,     3,     1,     1,     2,     0,     4,     0
};

/* YYDEFACT[STATE-NAME] -- Default rule to reduce with in state
   STATE-NUM when YYTABLE doesn't specify something else to do.  Zero
   means the default is an error.  */
static const yytype_uint8 yydefact[] =
{
       0,     0,     0,     0,     0,     0,    55,     0,     2,     3,
       9,    10,    11,     4,     5,    16,    17,    18,     6,     7,
     101,   106,   106,   106,     0,   102,   106,     0,     0,     0,
       0,    54,     0,     1,   105,     8,    12,    13,     0,    15,
     104,     0,   106,     0,    45,    72,    77,    78,    79,    80,
      81,   103,    52,     0,     0,    82,    50,    51,    71,     0,
      60,    63,    67,    70,    73,     0,    75,     0,    83,   106,
       0,    19,     0,     0,     0,     0,    69,     0,     0,     0,
      59,     0,     0,    56,     0,     0,     0,     0,     0,    14,
       0,   103,     0,    21,     0,     0,    46,     0,    68,    97,
      97,    97,     0,   108,    84,     0,    60,    60,    58,    61,
      62,    66,    64,    65,     0,    76,     0,   106,     0,    24,
      25,    26,    27,    28,    29,    30,    31,    34,    44,     0,
       0,    96,     0,     0,     0,    95,     0,    89,     0,     0,
      88,     0,   116,    99,    87,    57,    74,     0,    20,    22,
       0,    38,     0,     0,    41,     0,    23,    33,    35,    47,
       0,    48,    92,    93,    94,     0,     0,   107,   109,    90,
      91,    99,     0,   118,     0,    86,    98,   106,    37,    39,
      40,     0,    36,   106,     0,     0,     0,     0,   113,   114,
       0,    85,   115,     0,   106,   100,    42,    32,    43,    49,
     112,   111,   110,     0,    53,     0,   117
};

/* YYDEFGOTO[NTERM-NUM].  */
static const yytype_int16 yydefgoto[] =
{
      -1,     7,     8,     9,    10,    11,    12,    13,    14,    15,
      16,    92,    93,   127,   156,   157,   158,    17,    18,    74,
      95,   160,    58,    19,    32,    59,    83,   166,    61,    62,
      63,    64,    65,    66,   103,   104,   141,   105,   132,   175,
     176,    21,    67,    68,    41,    35,   142,   167,   168,   190,
     173,   194
};

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */
#define YYPACT_NINF -144
static const yytype_int16 yypact[] =
{
      89,   -39,   162,   -33,   137,    24,    27,    52,  -144,  -144,
    -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,
    -144,   -10,   -10,   -10,    63,  -144,   -10,     3,   -39,   -33,
     -33,  -144,    43,  -144,  -144,  -144,  -144,  -144,   -33,  -144,
    -144,    91,   -10,    25,    35,  -144,  -144,  -144,  -144,  -144,
    -144,    44,  -144,    43,    43,  -144,  -144,  -144,  -144,    -1,
     -11,   -23,  -144,  -144,  -144,    56,  -144,    48,  -144,   -10,
     -33,  -144,    88,    88,   116,    99,  -144,    30,    43,   117,
    -144,    43,    43,  -144,    43,    43,    43,    43,   -38,  -144,
      97,  -144,   -50,  -144,    72,   -25,  -144,   114,  -144,   128,
     128,   128,   136,    -7,  -144,   -33,     2,   -11,  -144,   -23,
     -23,  -144,  -144,  -144,   101,  -144,    88,   -10,    88,  -144,
    -144,  -144,  -144,  -144,  -144,  -144,  -144,    55,  -144,    88,
     -43,  -144,   138,   139,   140,  -144,    81,  -144,   141,   142,
    -144,   -33,   153,   170,  -144,  -144,  -144,     5,  -144,  -144,
     155,  -144,   154,   -43,  -144,   118,  -144,  -144,   103,  -144,
      29,  -144,  -144,  -144,  -144,    81,    94,    75,  -144,  -144,
    -144,   170,   119,   160,    81,  -144,  -144,   -10,  -144,  -144,
    -144,   129,  -144,   -10,   -43,     4,    -5,    43,  -144,  -144,
      81,  -144,  -144,   120,   -10,    75,  -144,  -144,  -144,  -144,
    -144,    65,  -144,   130,  -144,   122,  -144
};

/* YYPGOTO[NTERM-NUM].  */
static const yytype_int16 yypgoto[] =
{
    -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,  -144,
    -144,  -144,    76,  -144,  -144,    37,  -144,  -144,  -144,  -144,
      77,  -144,  -121,  -144,  -144,  -144,    36,   -30,    71,   -44,
    -144,  -144,  -144,   108,  -144,   -91,  -144,  -144,    70,    26,
    -144,   171,     0,   -67,  -144,   -22,  -144,  -143,     8,  -144,
    -144,  -144
};

/* YYTABLE[YYPACT[STATE-NUM]].  What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule which
   number is the opposite.  If zero, do what YYDEFACT says.
   If YYTABLE_NINF, syntax error.  */
#define YYTABLE_NINF -103
static const yytype_int16 yytable[] =
{
      36,    37,    60,    26,    39,    94,    96,    77,    79,   161,
      76,   117,   118,   136,   143,    52,   188,   189,    20,    91,
      71,    79,   186,    75,    25,    84,    56,    57,    55,    43,
      44,   195,   180,   137,   138,   139,   128,   129,    69,    30,
     111,   112,   113,    85,    86,    31,    80,    89,   107,    96,
     171,    94,    33,    81,    82,   140,   200,   114,    34,    80,
      40,    78,   159,   199,   187,    98,   177,   129,    81,    82,
      90,    38,    45,    99,   100,   101,   102,   106,    46,    47,
      48,    49,    50,   150,    72,   151,   152,    25,   153,   154,
     183,   184,     1,     2,    73,   148,   188,   189,     3,     4,
      51,    52,    53,     5,    70,   106,     6,  -102,    54,    55,
      45,    88,    56,    57,   155,    87,    46,    47,    48,    49,
      50,   119,   120,   121,   122,   123,   124,   125,   126,    81,
      82,   150,    97,   151,   152,   185,   153,   154,    51,    52,
     165,   106,   144,   145,    27,    91,    54,    55,    28,    29,
      56,    57,   109,   110,   187,   196,   116,   201,    81,    82,
      98,   198,   146,    81,    82,    81,    82,    22,    23,    24,
     133,   134,   204,   130,   108,   131,   135,   172,   162,   163,
     164,   169,   170,   174,   178,   193,   179,   181,   192,   203,
     197,   206,   205,   147,   149,   182,   115,   191,   202,    42
};

static const yytype_uint8 yycheck[] =
{
      22,    23,    32,     3,    26,    72,    73,     8,    19,   130,
      54,    61,    62,    20,   105,    58,    21,    22,    57,    57,
      42,    19,   165,    53,    57,    48,    69,    70,    66,    29,
      30,   174,   153,    40,    41,    42,    61,    62,    38,    15,
      84,    85,    86,    66,    67,    18,    57,    69,    78,   116,
     141,   118,     0,    64,    65,    62,    61,    87,    68,    57,
      57,    62,   129,   184,    60,    61,    61,    62,    64,    65,
      70,     8,    29,    43,    44,    45,    46,    77,    35,    36,
      37,    38,    39,    28,    59,    30,    31,    57,    33,    34,
      61,    62,     3,     4,    59,   117,    21,    22,     9,    10,
      57,    58,    59,    14,    13,   105,    17,    63,    65,    66,
      29,    63,    69,    70,    59,    59,    35,    36,    37,    38,
      39,    49,    50,    51,    52,    53,    54,    55,    56,    64,
      65,    28,    16,    30,    31,   165,    33,    34,    57,    58,
      59,   141,   106,   107,     7,    57,    65,    66,    11,    12,
      69,    70,    81,    82,    60,   177,    59,   187,    64,    65,
      61,   183,    61,    64,    65,    64,    65,     5,     6,     7,
     100,   101,   194,    59,    57,    47,    40,    24,    40,    40,
      40,    40,    40,    13,    29,    25,    32,    69,    69,    69,
      61,    69,    62,   116,   118,   158,    88,   171,   190,    28
};

/* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
   symbol of state STATE-NUM.  */
static const yytype_uint8 yystos[] =
{
       0,     3,     4,     9,    10,    14,    17,    72,    73,    74,
      75,    76,    77,    78,    79,    80,    81,    88,    89,    94,
      57,   112,     5,     6,     7,    57,   113,     7,    11,    12,
      15,    18,    95,     0,    68,   116,   116,   116,     8,   116,
      57,   115,   112,   113,   113,    29,    35,    36,    37,    38,
      39,    57,    58,    59,    65,    66,    69,    70,    93,    96,
      98,    99,   100,   101,   102,   103,   104,   113,   114,   113,
      13,   116,    59,    59,    90,    98,   100,     8,    62,    19,
      57,    64,    65,    97,    48,    66,    67,    59,    63,   116,
     113,    57,    82,    83,   114,    91,   114,    16,    61,    43,
      44,    45,    46,   105,   106,   108,   113,    98,    57,    99,
      99,   100,   100,   100,    98,   104,    59,    61,    62,    49,
      50,    51,    52,    53,    54,    55,    56,    84,    61,    62,
      59,    47,   109,   109,   109,    40,    20,    40,    41,    42,
      62,   107,   117,   106,    97,    97,    61,    91,   116,    83,
      28,    30,    31,    33,    34,    59,    85,    86,    87,   114,
      92,    93,    40,    40,    40,    59,    98,   118,   119,    40,
      40,   106,    24,   121,    13,   110,   111,    61,    29,    32,
      93,    69,    86,    61,    62,    98,   118,    60,    21,    22,
     120,   110,    69,    25,   122,   118,   116,    61,   116,    93,
      61,    98,   119,    69,   116,    62,    69
};

#define yyerrok		(yyerrstatus = 0)
#define yyclearin	(yychar = YYEMPTY)
#define YYEMPTY		(-2)
#define YYEOF		0

#define YYACCEPT	goto yyacceptlab
#define YYABORT		goto yyabortlab
#define YYERROR		goto yyerrorlab


/* Like YYERROR except do call yyerror.  This remains here temporarily
   to ease the transition to the new meaning of YYERROR, for GCC.
   Once GCC version 2 has supplanted version 1, this can go.  */

#define YYFAIL		goto yyerrlab

#define YYRECOVERING()  (!!yyerrstatus)

#define YYBACKUP(Token, Value)					\
do								\
  if (yychar == YYEMPTY && yylen == 1)				\
    {								\
      yychar = (Token);						\
      yylval = (Value);						\
      yytoken = YYTRANSLATE (yychar);				\
      YYPOPSTACK (1);						\
      goto yybackup;						\
    }								\
  else								\
    {								\
      yyerror (YY_("syntax error: cannot back up")); \
      YYERROR;							\
    }								\
while (YYID (0))


#define YYTERROR	1
#define YYERRCODE	256


/* YYLLOC_DEFAULT -- Set CURRENT to span from RHS[1] to RHS[N].
   If N is 0, then set CURRENT to the empty location which ends
   the previous symbol: RHS[0] (always defined).  */

#define YYRHSLOC(Rhs, K) ((Rhs)[K])
#ifndef YYLLOC_DEFAULT
# define YYLLOC_DEFAULT(Current, Rhs, N)				\
    do									\
      if (YYID (N))                                                    \
	{								\
	  (Current).first_line   = YYRHSLOC (Rhs, 1).first_line;	\
	  (Current).first_column = YYRHSLOC (Rhs, 1).first_column;	\
	  (Current).last_line    = YYRHSLOC (Rhs, N).last_line;		\
	  (Current).last_column  = YYRHSLOC (Rhs, N).last_column;	\
	}								\
      else								\
	{								\
	  (Current).first_line   = (Current).last_line   =		\
	    YYRHSLOC (Rhs, 0).last_line;				\
	  (Current).first_column = (Current).last_column =		\
	    YYRHSLOC (Rhs, 0).last_column;				\
	}								\
    while (YYID (0))
#endif


/* YY_LOCATION_PRINT -- Print the location on the stream.
   This macro was not mandated originally: define only if we know
   we won't break user code: when these are the locations we know.  */

#ifndef YY_LOCATION_PRINT
# if defined YYLTYPE_IS_TRIVIAL && YYLTYPE_IS_TRIVIAL
#  define YY_LOCATION_PRINT(File, Loc)			\
     fprintf (File, "%d.%d-%d.%d",			\
	      (Loc).first_line, (Loc).first_column,	\
	      (Loc).last_line,  (Loc).last_column)
# else
#  define YY_LOCATION_PRINT(File, Loc) ((void) 0)
# endif
#endif


/* YYLEX -- calling `yylex' with the right arguments.  */

#ifdef YYLEX_PARAM
# define YYLEX yylex (YYLEX_PARAM)
#else
# define YYLEX yylex ()
#endif

/* Enable debugging if requested.  */
#if YYDEBUG

# ifndef YYFPRINTF
#  include <stdio.h> /* INFRINGES ON USER NAME SPACE */
#  define YYFPRINTF fprintf
# endif

# define YYDPRINTF(Args)			\
do {						\
  if (yydebug)					\
    YYFPRINTF Args;				\
} while (YYID (0))

# define YY_SYMBOL_PRINT(Title, Type, Value, Location)			  \
do {									  \
  if (yydebug)								  \
    {									  \
      YYFPRINTF (stderr, "%s ", Title);					  \
      yy_symbol_print (stderr,						  \
		  Type, Value); \
      YYFPRINTF (stderr, "\n");						  \
    }									  \
} while (YYID (0))


/*--------------------------------.
| Print this symbol on YYOUTPUT.  |
`--------------------------------*/

/*ARGSUSED*/
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
yy_symbol_value_print (FILE *yyoutput, int yytype, YYSTYPE const * const yyvaluep)
#else
static void
yy_symbol_value_print (yyoutput, yytype, yyvaluep)
    FILE *yyoutput;
    int yytype;
    YYSTYPE const * const yyvaluep;
#endif
{
  if (!yyvaluep)
    return;
# ifdef YYPRINT
  if (yytype < YYNTOKENS)
    YYPRINT (yyoutput, yytoknum[yytype], *yyvaluep);
# else
  YYUSE (yyoutput);
# endif
  switch (yytype)
    {
      default:
	break;
    }
}


/*--------------------------------.
| Print this symbol on YYOUTPUT.  |
`--------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
yy_symbol_print (FILE *yyoutput, int yytype, YYSTYPE const * const yyvaluep)
#else
static void
yy_symbol_print (yyoutput, yytype, yyvaluep)
    FILE *yyoutput;
    int yytype;
    YYSTYPE const * const yyvaluep;
#endif
{
  if (yytype < YYNTOKENS)
    YYFPRINTF (yyoutput, "token %s (", yytname[yytype]);
  else
    YYFPRINTF (yyoutput, "nterm %s (", yytname[yytype]);

  yy_symbol_value_print (yyoutput, yytype, yyvaluep);
  YYFPRINTF (yyoutput, ")");
}

/*------------------------------------------------------------------.
| yy_stack_print -- Print the state stack from its BOTTOM up to its |
| TOP (included).                                                   |
`------------------------------------------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
yy_stack_print (yytype_int16 *bottom, yytype_int16 *top)
#else
static void
yy_stack_print (bottom, top)
    yytype_int16 *bottom;
    yytype_int16 *top;
#endif
{
  YYFPRINTF (stderr, "Stack now");
  for (; bottom <= top; ++bottom)
    YYFPRINTF (stderr, " %d", *bottom);
  YYFPRINTF (stderr, "\n");
}

# define YY_STACK_PRINT(Bottom, Top)				\
do {								\
  if (yydebug)							\
    yy_stack_print ((Bottom), (Top));				\
} while (YYID (0))


/*------------------------------------------------.
| Report that the YYRULE is going to be reduced.  |
`------------------------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
yy_reduce_print (YYSTYPE *yyvsp, int yyrule)
#else
static void
yy_reduce_print (yyvsp, yyrule)
    YYSTYPE *yyvsp;
    int yyrule;
#endif
{
  int yynrhs = yyr2[yyrule];
  int yyi;
  unsigned long int yylno = yyrline[yyrule];
  YYFPRINTF (stderr, "Reducing stack by rule %d (line %lu):\n",
	     yyrule - 1, yylno);
  /* The symbols being reduced.  */
  for (yyi = 0; yyi < yynrhs; yyi++)
    {
      fprintf (stderr, "   $%d = ", yyi + 1);
      yy_symbol_print (stderr, yyrhs[yyprhs[yyrule] + yyi],
		       &(yyvsp[(yyi + 1) - (yynrhs)])
		       		       );
      fprintf (stderr, "\n");
    }
}

# define YY_REDUCE_PRINT(Rule)		\
do {					\
  if (yydebug)				\
    yy_reduce_print (yyvsp, Rule); \
} while (YYID (0))

/* Nonzero means print parse trace.  It is left uninitialized so that
   multiple parsers can coexist.  */
int yydebug;
#else /* !YYDEBUG */
# define YYDPRINTF(Args)
# define YY_SYMBOL_PRINT(Title, Type, Value, Location)
# define YY_STACK_PRINT(Bottom, Top)
# define YY_REDUCE_PRINT(Rule)
#endif /* !YYDEBUG */


/* YYINITDEPTH -- initial size of the parser's stacks.  */
#ifndef	YYINITDEPTH
# define YYINITDEPTH 200
#endif

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

#ifndef YYMAXDEPTH
# define YYMAXDEPTH 10000
#endif



#if YYERROR_VERBOSE

# ifndef yystrlen
#  if defined __GLIBC__ && defined _STRING_H
#   define yystrlen strlen
#  else
/* Return the length of YYSTR.  */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static YYSIZE_T
yystrlen (const char *yystr)
#else
static YYSIZE_T
yystrlen (yystr)
    const char *yystr;
#endif
{
  YYSIZE_T yylen;
  for (yylen = 0; yystr[yylen]; yylen++)
    continue;
  return yylen;
}
#  endif
# endif

# ifndef yystpcpy
#  if defined __GLIBC__ && defined _STRING_H && defined _GNU_SOURCE
#   define yystpcpy stpcpy
#  else
/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static char *
yystpcpy (char *yydest, const char *yysrc)
#else
static char *
yystpcpy (yydest, yysrc)
    char *yydest;
    const char *yysrc;
#endif
{
  char *yyd = yydest;
  const char *yys = yysrc;

  while ((*yyd++ = *yys++) != '\0')
    continue;

  return yyd - 1;
}
#  endif
# endif

# ifndef yytnamerr
/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for yyerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from yytname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */
static YYSIZE_T
yytnamerr (char *yyres, const char *yystr)
{
  if (*yystr == '"')
    {
      YYSIZE_T yyn = 0;
      char const *yyp = yystr;

      for (;;)
	switch (*++yyp)
	  {
	  case '\'':
	  case ',':
	    goto do_not_strip_quotes;

	  case '\\':
	    if (*++yyp != '\\')
	      goto do_not_strip_quotes;
	    /* Fall through.  */
	  default:
	    if (yyres)
	      yyres[yyn] = *yyp;
	    yyn++;
	    break;

	  case '"':
	    if (yyres)
	      yyres[yyn] = '\0';
	    return yyn;
	  }
    do_not_strip_quotes: ;
    }

  if (! yyres)
    return yystrlen (yystr);

  return yystpcpy (yyres, yystr) - yyres;
}
# endif

/* Copy into YYRESULT an error message about the unexpected token
   YYCHAR while in state YYSTATE.  Return the number of bytes copied,
   including the terminating null byte.  If YYRESULT is null, do not
   copy anything; just return the number of bytes that would be
   copied.  As a special case, return 0 if an ordinary "syntax error"
   message will do.  Return YYSIZE_MAXIMUM if overflow occurs during
   size calculation.  */
static YYSIZE_T
yysyntax_error (char *yyresult, int yystate, int yychar)
{
  int yyn = yypact[yystate];

  if (! (YYPACT_NINF < yyn && yyn <= YYLAST))
    return 0;
  else
    {
      int yytype = YYTRANSLATE (yychar);
      YYSIZE_T yysize0 = yytnamerr (0, yytname[yytype]);
      YYSIZE_T yysize = yysize0;
      YYSIZE_T yysize1;
      int yysize_overflow = 0;
      enum { YYERROR_VERBOSE_ARGS_MAXIMUM = 5 };
      char const *yyarg[YYERROR_VERBOSE_ARGS_MAXIMUM];
      int yyx;

# if 0
      /* This is so xgettext sees the translatable formats that are
	 constructed on the fly.  */
      YY_("syntax error, unexpected %s");
      YY_("syntax error, unexpected %s, expecting %s");
      YY_("syntax error, unexpected %s, expecting %s or %s");
      YY_("syntax error, unexpected %s, expecting %s or %s or %s");
      YY_("syntax error, unexpected %s, expecting %s or %s or %s or %s");
# endif
      char *yyfmt;
      char const *yyf;
      static char const yyunexpected[] = "syntax error, unexpected %s";
      static char const yyexpecting[] = ", expecting %s";
      static char const yyor[] = " or %s";
      char yyformat[sizeof yyunexpected
		    + sizeof yyexpecting - 1
		    + ((YYERROR_VERBOSE_ARGS_MAXIMUM - 2)
		       * (sizeof yyor - 1))];
      char const *yyprefix = yyexpecting;

      /* Start YYX at -YYN if negative to avoid negative indexes in
	 YYCHECK.  */
      int yyxbegin = yyn < 0 ? -yyn : 0;

      /* Stay within bounds of both yycheck and yytname.  */
      int yychecklim = YYLAST - yyn + 1;
      int yyxend = yychecklim < YYNTOKENS ? yychecklim : YYNTOKENS;
      int yycount = 1;

      yyarg[0] = yytname[yytype];
      yyfmt = yystpcpy (yyformat, yyunexpected);

      for (yyx = yyxbegin; yyx < yyxend; ++yyx)
	if (yycheck[yyx + yyn] == yyx && yyx != YYTERROR)
	  {
	    if (yycount == YYERROR_VERBOSE_ARGS_MAXIMUM)
	      {
		yycount = 1;
		yysize = yysize0;
		yyformat[sizeof yyunexpected - 1] = '\0';
		break;
	      }
	    yyarg[yycount++] = yytname[yyx];
	    yysize1 = yysize + yytnamerr (0, yytname[yyx]);
	    yysize_overflow |= (yysize1 < yysize);
	    yysize = yysize1;
	    yyfmt = yystpcpy (yyfmt, yyprefix);
	    yyprefix = yyor;
	  }

      yyf = YY_(yyformat);
      yysize1 = yysize + yystrlen (yyf);
      yysize_overflow |= (yysize1 < yysize);
      yysize = yysize1;

      if (yysize_overflow)
	return YYSIZE_MAXIMUM;

      if (yyresult)
	{
	  /* Avoid sprintf, as that infringes on the user's name space.
	     Don't have undefined behavior even if the translation
	     produced a string with the wrong number of "%s"s.  */
	  char *yyp = yyresult;
	  int yyi = 0;
	  while ((*yyp = *yyf) != '\0')
	    {
	      if (*yyp == '%' && yyf[1] == 's' && yyi < yycount)
		{
		  yyp += yytnamerr (yyp, yyarg[yyi++]);
		  yyf += 2;
		}
	      else
		{
		  yyp++;
		  yyf++;
		}
	    }
	}
      return yysize;
    }
}
#endif /* YYERROR_VERBOSE */


/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

/*ARGSUSED*/
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
yydestruct (const char *yymsg, int yytype, YYSTYPE *yyvaluep)
#else
static void
yydestruct (yymsg, yytype, yyvaluep)
    const char *yymsg;
    int yytype;
    YYSTYPE *yyvaluep;
#endif
{
  YYUSE (yyvaluep);

  if (!yymsg)
    yymsg = "Deleting";
  YY_SYMBOL_PRINT (yymsg, yytype, yyvaluep, yylocationp);

  switch (yytype)
    {

      default:
	break;
    }
}


/* Prevent warnings from -Wmissing-prototypes.  */

#ifdef YYPARSE_PARAM
#if defined __STDC__ || defined __cplusplus
int yyparse (void *YYPARSE_PARAM);
#else
int yyparse ();
#endif
#else /* ! YYPARSE_PARAM */
#if defined __STDC__ || defined __cplusplus
int yyparse (void);
#else
int yyparse ();
#endif
#endif /* ! YYPARSE_PARAM */



/* The look-ahead symbol.  */
int yychar;

/* The semantic value of the look-ahead symbol.  */
YYSTYPE yylval;

/* Number of syntax errors so far.  */
int yynerrs;



/*----------.
| yyparse.  |
`----------*/

#ifdef YYPARSE_PARAM
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
int
yyparse (void *YYPARSE_PARAM)
#else
int
yyparse (YYPARSE_PARAM)
    void *YYPARSE_PARAM;
#endif
#else /* ! YYPARSE_PARAM */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
int
yyparse (void)
#else
int
yyparse ()

#endif
#endif
{
  
  int yystate;
  int yyn;
  int yyresult;
  /* Number of tokens to shift before error messages enabled.  */
  int yyerrstatus;
  /* Look-ahead token as an internal (translated) token number.  */
  int yytoken = 0;
#if YYERROR_VERBOSE
  /* Buffer for error messages, and its allocated size.  */
  char yymsgbuf[128];
  char *yymsg = yymsgbuf;
  YYSIZE_T yymsg_alloc = sizeof yymsgbuf;
#endif

  /* Three stacks and their tools:
     `yyss': related to states,
     `yyvs': related to semantic values,
     `yyls': related to locations.

     Refer to the stacks thru separate pointers, to allow yyoverflow
     to reallocate them elsewhere.  */

  /* The state stack.  */
  yytype_int16 yyssa[YYINITDEPTH];
  yytype_int16 *yyss = yyssa;
  yytype_int16 *yyssp;

  /* The semantic value stack.  */
  YYSTYPE yyvsa[YYINITDEPTH];
  YYSTYPE *yyvs = yyvsa;
  YYSTYPE *yyvsp;



#define YYPOPSTACK(N)   (yyvsp -= (N), yyssp -= (N))

  YYSIZE_T yystacksize = YYINITDEPTH;

  /* The variables used to return semantic value and location from the
     action routines.  */
  YYSTYPE yyval;


  /* The number of symbols on the RHS of the reduced rule.
     Keep to zero when no symbol should be popped.  */
  int yylen = 0;

  YYDPRINTF ((stderr, "Starting parse\n"));

  yystate = 0;
  yyerrstatus = 0;
  yynerrs = 0;
  yychar = YYEMPTY;		/* Cause a token to be read.  */

  /* Initialize stack pointers.
     Waste one element of value and location stack
     so that they stay on the same level as the state stack.
     The wasted elements are never initialized.  */

  yyssp = yyss;
  yyvsp = yyvs;

  goto yysetstate;

/*------------------------------------------------------------.
| yynewstate -- Push a new state, which is found in yystate.  |
`------------------------------------------------------------*/
 yynewstate:
  /* In all cases, when you get here, the value and location stacks
     have just been pushed.  So pushing a state here evens the stacks.  */
  yyssp++;

 yysetstate:
  *yyssp = yystate;

  if (yyss + yystacksize - 1 <= yyssp)
    {
      /* Get the current used size of the three stacks, in elements.  */
      YYSIZE_T yysize = yyssp - yyss + 1;

#ifdef yyoverflow
      {
	/* Give user a chance to reallocate the stack.  Use copies of
	   these so that the &'s don't force the real ones into
	   memory.  */
	YYSTYPE *yyvs1 = yyvs;
	yytype_int16 *yyss1 = yyss;


	/* Each stack pointer address is followed by the size of the
	   data in use in that stack, in bytes.  This used to be a
	   conditional around just the two extra args, but that might
	   be undefined if yyoverflow is a macro.  */
	yyoverflow (YY_("memory exhausted"),
		    &yyss1, yysize * sizeof (*yyssp),
		    &yyvs1, yysize * sizeof (*yyvsp),

		    &yystacksize);

	yyss = yyss1;
	yyvs = yyvs1;
      }
#else /* no yyoverflow */
# ifndef YYSTACK_RELOCATE
      goto yyexhaustedlab;
# else
      /* Extend the stack our own way.  */
      if (YYMAXDEPTH <= yystacksize)
	goto yyexhaustedlab;
      yystacksize *= 2;
      if (YYMAXDEPTH < yystacksize)
	yystacksize = YYMAXDEPTH;

      {
	yytype_int16 *yyss1 = yyss;
	union yyalloc *yyptr =
	  (union yyalloc *) YYSTACK_ALLOC (YYSTACK_BYTES (yystacksize));
	if (! yyptr)
	  goto yyexhaustedlab;
	YYSTACK_RELOCATE (yyss);
	YYSTACK_RELOCATE (yyvs);

#  undef YYSTACK_RELOCATE
	if (yyss1 != yyssa)
	  YYSTACK_FREE (yyss1);
      }
# endif
#endif /* no yyoverflow */

      yyssp = yyss + yysize - 1;
      yyvsp = yyvs + yysize - 1;


      YYDPRINTF ((stderr, "Stack size increased to %lu\n",
		  (unsigned long int) yystacksize));

      if (yyss + yystacksize - 1 <= yyssp)
	YYABORT;
    }

  YYDPRINTF ((stderr, "Entering state %d\n", yystate));

  goto yybackup;

/*-----------.
| yybackup.  |
`-----------*/
yybackup:

  /* Do appropriate processing given the current state.  Read a
     look-ahead token if we need one and don't already have one.  */

  /* First try to decide what to do without reference to look-ahead token.  */
  yyn = yypact[yystate];
  if (yyn == YYPACT_NINF)
    goto yydefault;

  /* Not known => get a look-ahead token if don't already have one.  */

  /* YYCHAR is either YYEMPTY or YYEOF or a valid look-ahead symbol.  */
  if (yychar == YYEMPTY)
    {
      YYDPRINTF ((stderr, "Reading a token: "));
      yychar = YYLEX;
    }

  if (yychar <= YYEOF)
    {
      yychar = yytoken = YYEOF;
      YYDPRINTF ((stderr, "Now at end of input.\n"));
    }
  else
    {
      yytoken = YYTRANSLATE (yychar);
      YY_SYMBOL_PRINT ("Next token is", yytoken, &yylval, &yylloc);
    }

  /* If the proper action on seeing token YYTOKEN is to reduce or to
     detect an error, take that action.  */
  yyn += yytoken;
  if (yyn < 0 || YYLAST < yyn || yycheck[yyn] != yytoken)
    goto yydefault;
  yyn = yytable[yyn];
  if (yyn <= 0)
    {
      if (yyn == 0 || yyn == YYTABLE_NINF)
	goto yyerrlab;
      yyn = -yyn;
      goto yyreduce;
    }

  if (yyn == YYFINAL)
    YYACCEPT;

  /* Count tokens shifted since error; after three, turn off error
     status.  */
  if (yyerrstatus)
    yyerrstatus--;

  /* Shift the look-ahead token.  */
  YY_SYMBOL_PRINT ("Shifting", yytoken, &yylval, &yylloc);

  /* Discard the shifted token unless it is eof.  */
  if (yychar != YYEOF)
    yychar = YYEMPTY;

  yystate = yyn;
  *++yyvsp = yylval;

  goto yynewstate;


/*-----------------------------------------------------------.
| yydefault -- do the default action for the current state.  |
`-----------------------------------------------------------*/
yydefault:
  yyn = yydefact[yystate];
  if (yyn == 0)
    goto yyerrlab;
  goto yyreduce;


/*-----------------------------.
| yyreduce -- Do a reduction.  |
`-----------------------------*/
yyreduce:
  /* yyn is the number of a rule to reduce with.  */
  yylen = yyr2[yyn];

  /* If YYLEN is nonzero, implement the default value of the action:
     `$$ = $1'.

     Otherwise, the following line sets YYVAL to garbage.
     This behavior is undocumented and Bison
     users should not rely upon it.  Assigning to YYVAL
     unconditionally makes the parser a bit smaller, and it avoids a
     GCC warning that YYVAL may be used uninitialized.  */
  yyval = yyvsp[1-yylen];


  YY_REDUCE_PRINT (yyn);
  switch (yyn)
    {
        case 2:
#line 53 "yacc_redisql.y"
    {
		setType(REDISQL_USE);
		setDatabaseName((yyvsp[(1) - (1)].strOption));
	}
    break;

  case 3:
#line 58 "yacc_redisql.y"
    {
		//no code
	}
    break;

  case 4:
#line 62 "yacc_redisql.y"
    {
		setType(REDISQL_DESC);
		setTableName((yyvsp[(1) - (1)].strOption));
	}
    break;

  case 5:
#line 67 "yacc_redisql.y"
    {
		//no code
	}
    break;

  case 6:
#line 71 "yacc_redisql.y"
    {
		setType(REDISQL_INSERT);
	}
    break;

  case 7:
#line 75 "yacc_redisql.y"
    {
		setType(REDISQL_SELECT);
	}
    break;

  case 8:
#line 82 "yacc_redisql.y"
    {
		(yyval.strOption) = (yyvsp[(2) - (3)].strDatabase);
	}
    break;

  case 9:
#line 89 "yacc_redisql.y"
    {
		setType(REDISQL_SHOW_DATABASES);
	}
    break;

  case 10:
#line 93 "yacc_redisql.y"
    {
		setType(REDISQL_SHOW_TABLES);
	}
    break;

  case 11:
#line 97 "yacc_redisql.y"
    {
		setType(REDISQL_SHOW_INDEX);
		setTableName((yyvsp[(1) - (1)].strOption));
	}
    break;

  case 12:
#line 105 "yacc_redisql.y"
    {
		(yyval.strOption) = NULL;
	}
    break;

  case 13:
#line 112 "yacc_redisql.y"
    {
		(yyval.strOption) = NULL;
	}
    break;

  case 14:
#line 119 "yacc_redisql.y"
    {
		(yyval.strOption) = (yyvsp[(4) - (5)].strTable);
	}
    break;

  case 15:
#line 126 "yacc_redisql.y"
    {
		(yyval.strOption) = (yyvsp[(2) - (3)].strTable);
	}
    break;

  case 16:
#line 133 "yacc_redisql.y"
    {
		setType(REDISQL_CREATE_DATABASE);
		setDatabaseName((yyvsp[(1) - (1)].strOption));
	}
    break;

  case 17:
#line 138 "yacc_redisql.y"
    {
		setType(REDISQL_CREATE_TABLE);
		setTableName((yyvsp[(1) - (1)].strOption));
	}
    break;

  case 18:
#line 143 "yacc_redisql.y"
    {
		setType(REDISQL_CREATE_INDEX);
	}
    break;

  case 19:
#line 150 "yacc_redisql.y"
    {
		(yyval.strOption) = (yyvsp[(3) - (4)].strDatabase);
	}
    break;

  case 20:
#line 157 "yacc_redisql.y"
    {
		(yyval.strOption) = (yyvsp[(3) - (7)].strTable);
	}
    break;

  case 21:
#line 164 "yacc_redisql.y"
    {

	}
    break;

  case 22:
#line 168 "yacc_redisql.y"
    {

	}
    break;

  case 23:
#line 175 "yacc_redisql.y"
    {
		addFieldType((yyvsp[(1) - (3)].strColumn), (yyvsp[(2) - (3)].strColumn));
	}
    break;

  case 24:
#line 182 "yacc_redisql.y"
    {
		(yyval.strColumn) = "NUMBER";
	}
    break;

  case 25:
#line 186 "yacc_redisql.y"
    {
		(yyval.strColumn) = "NUMBER";
	}
    break;

  case 26:
#line 190 "yacc_redisql.y"
    {
		(yyval.strColumn) = "NUMBER";
	}
    break;

  case 27:
#line 194 "yacc_redisql.y"
    {
		(yyval.strColumn) = "STRING";
	}
    break;

  case 28:
#line 198 "yacc_redisql.y"
    {
		(yyval.strColumn) = "STRING";
	}
    break;

  case 29:
#line 202 "yacc_redisql.y"
    {
		(yyval.strColumn) = "STRING";
	}
    break;

  case 30:
#line 206 "yacc_redisql.y"
    {
		(yyval.strColumn) = "DATE";
	}
    break;

  case 31:
#line 210 "yacc_redisql.y"
    {
		(yyval.strColumn) = "DATE";
	}
    break;

  case 32:
#line 214 "yacc_redisql.y"
    {
		(yyval.strColumn) = (yyvsp[(1) - (4)].strColumn);
	}
    break;

  case 33:
#line 220 "yacc_redisql.y"
    {}
    break;

  case 34:
#line 221 "yacc_redisql.y"
    {}
    break;

  case 35:
#line 225 "yacc_redisql.y"
    {}
    break;

  case 36:
#line 226 "yacc_redisql.y"
    {}
    break;

  case 37:
#line 230 "yacc_redisql.y"
    {}
    break;

  case 38:
#line 231 "yacc_redisql.y"
    {}
    break;

  case 39:
#line 232 "yacc_redisql.y"
    {}
    break;

  case 40:
#line 233 "yacc_redisql.y"
    {}
    break;

  case 41:
#line 234 "yacc_redisql.y"
    {}
    break;

  case 42:
#line 239 "yacc_redisql.y"
    {
		setIndexName((yyvsp[(3) - (9)].strIndex));
		setTableName((yyvsp[(5) - (9)].strTable));
	}
    break;

  case 43:
#line 247 "yacc_redisql.y"
    {
		setTableName((yyvsp[(3) - (9)].strTable));
	}
    break;

  case 44:
#line 254 "yacc_redisql.y"
    {

	}
    break;

  case 45:
#line 258 "yacc_redisql.y"
    {

	}
    break;

  case 46:
#line 265 "yacc_redisql.y"
    {
		addFieldType((yyvsp[(1) - (1)].strColumn), "");
	}
    break;

  case 47:
#line 269 "yacc_redisql.y"
    {
		addFieldType((yyvsp[(3) - (3)].strColumn), "");
	}
    break;

  case 48:
#line 277 "yacc_redisql.y"
    {

	}
    break;

  case 49:
#line 281 "yacc_redisql.y"
    {

	}
    break;

  case 50:
#line 288 "yacc_redisql.y"
    {
		char pc[1000] = {'\0'};
		union Value uVal;
		uVal.nValue = (yyvsp[(1) - (1)].nVal);
		addFieldValue(REDISQL_INT, uVal);
		sprintf(pc, "%d", (yyvsp[(1) - (1)].nVal));
		(yyval.strVal) = pc;
	}
    break;

  case 51:
#line 297 "yacc_redisql.y"
    {
		char pc[1000] = {'\0'};
		union Value uVal;
		uVal.fValue = (yyvsp[(1) - (1)].fVal);
		addFieldValue(REDISQL_FLOAT, uVal);
		sprintf(pc, "%f", (yyvsp[(1) - (1)].fVal));
		(yyval.strVal) = pc;
	}
    break;

  case 52:
#line 306 "yacc_redisql.y"
    {
		union Value uVal;
		uVal.pcValue = (yyvsp[(1) - (1)].strVal);
		addFieldValue(REDISQL_STRING, uVal);
		(yyval.strVal) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 53:
#line 316 "yacc_redisql.y"
    {
		setType(REDISQL_SELECT);
	}
    break;

  case 54:
#line 322 "yacc_redisql.y"
    {}
    break;

  case 55:
#line 323 "yacc_redisql.y"
    {}
    break;

  case 56:
#line 327 "yacc_redisql.y"
    {
		FieldAlias st;
		st = (yyvsp[(1) - (2)].stFieldAlias);
		if (strcmp((yyvsp[(2) - (2)].strVal), "") != 0)
		{
			addFieldAlias(st.pcTableAlias, st.pcField, (yyvsp[(2) - (2)].strVal));
		}
		else
		{
			addFieldAlias(st.pcTableAlias, st.pcField, st.pcField);
		}
	}
    break;

  case 57:
#line 340 "yacc_redisql.y"
    {
		FieldAlias st;
		st = (yyvsp[(3) - (4)].stFieldAlias);
		if (strcmp((yyvsp[(4) - (4)].strVal), "") != 0)
		{
			addFieldAlias(st.pcTableAlias, st.pcField, (yyvsp[(4) - (4)].strVal));
		}
		else
		{
			addFieldAlias(st.pcTableAlias, st.pcField, st.pcField);
		}
	}
    break;

  case 58:
#line 356 "yacc_redisql.y"
    {
		(yyval.strVal) = (yyvsp[(2) - (2)].strVal);
	}
    break;

  case 59:
#line 360 "yacc_redisql.y"
    {
		(yyval.strVal) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 60:
#line 364 "yacc_redisql.y"
    {
		(yyval.strVal) = "";
	}
    break;

  case 63:
#line 373 "yacc_redisql.y"
    {
		(yyval.stFieldAlias) = (yyvsp[(1) - (1)].stFieldAlias);
	}
    break;

  case 64:
#line 380 "yacc_redisql.y"
    {}
    break;

  case 65:
#line 382 "yacc_redisql.y"
    {}
    break;

  case 66:
#line 384 "yacc_redisql.y"
    {}
    break;

  case 67:
#line 386 "yacc_redisql.y"
    { 
		(yyval.stFieldAlias) = (yyvsp[(1) - (1)].stFieldAlias); 
	}
    break;

  case 68:
#line 393 "yacc_redisql.y"
    { 
		(yyval.stFieldAlias) = (yyvsp[(2) - (3)].stFieldAlias); 
	}
    break;

  case 69:
#line 397 "yacc_redisql.y"
    { 
		(yyval.stFieldAlias) = (yyvsp[(2) - (2)].stFieldAlias); 
	}
    break;

  case 70:
#line 401 "yacc_redisql.y"
    { 
		(yyval.stFieldAlias) = (yyvsp[(1) - (1)].stFieldAlias); 
	}
    break;

  case 71:
#line 408 "yacc_redisql.y"
    {
		FieldAlias st;
		st.pcTableAlias = NULL;
		st.pcField = (yyvsp[(1) - (1)].strVal);
		(yyval.stFieldAlias) = st;
	}
    break;

  case 72:
#line 415 "yacc_redisql.y"
    {

	}
    break;

  case 73:
#line 419 "yacc_redisql.y"
    {
		(yyval.stFieldAlias) = (yyvsp[(1) - (1)].stFieldAlias);
	}
    break;

  case 74:
#line 423 "yacc_redisql.y"
    {

	}
    break;

  case 75:
#line 430 "yacc_redisql.y"
    {
		FieldAlias stTmp;
		stTmp.pcTableAlias = "";
		stTmp.pcField = (yyvsp[(1) - (1)].strVal);
		stTmp.pcAlias = (yyvsp[(1) - (1)].strVal);
		(yyval.stFieldAlias) = stTmp;
	}
    break;

  case 76:
#line 438 "yacc_redisql.y"
    {
		FieldAlias stTmp;
		stTmp.pcTableAlias = (yyvsp[(1) - (3)].strTable);
		stTmp.pcField = (yyvsp[(3) - (3)].strVal);
		stTmp.pcAlias = (yyvsp[(3) - (3)].strVal);
		(yyval.stFieldAlias) = stTmp;
	}
    break;

  case 77:
#line 448 "yacc_redisql.y"
    {}
    break;

  case 78:
#line 449 "yacc_redisql.y"
    {}
    break;

  case 79:
#line 450 "yacc_redisql.y"
    {}
    break;

  case 80:
#line 451 "yacc_redisql.y"
    {}
    break;

  case 81:
#line 452 "yacc_redisql.y"
    {}
    break;

  case 82:
#line 457 "yacc_redisql.y"
    { 
		(yyval.strVal) = "*";
	}
    break;

  case 83:
#line 461 "yacc_redisql.y"
    {
		(yyval.strVal) = (yyvsp[(1) - (1)].strColumn);
	}
    break;

  case 84:
#line 468 "yacc_redisql.y"
    {}
    break;

  case 85:
#line 470 "yacc_redisql.y"
    {}
    break;

  case 86:
#line 472 "yacc_redisql.y"
    {}
    break;

  case 87:
#line 477 "yacc_redisql.y"
    {
		addTableAlias((yyvsp[(1) - (2)].strTable), (yyvsp[(2) - (2)].strVal));
	}
    break;

  case 89:
#line 484 "yacc_redisql.y"
    {}
    break;

  case 90:
#line 485 "yacc_redisql.y"
    {}
    break;

  case 91:
#line 486 "yacc_redisql.y"
    {}
    break;

  case 92:
#line 490 "yacc_redisql.y"
    {}
    break;

  case 93:
#line 491 "yacc_redisql.y"
    {}
    break;

  case 94:
#line 492 "yacc_redisql.y"
    {}
    break;

  case 95:
#line 493 "yacc_redisql.y"
    {}
    break;

  case 96:
#line 497 "yacc_redisql.y"
    {}
    break;

  case 97:
#line 498 "yacc_redisql.y"
    {}
    break;

  case 98:
#line 503 "yacc_redisql.y"
    {

	}
    break;

  case 99:
#line 507 "yacc_redisql.y"
    {

	}
    break;

  case 100:
#line 514 "yacc_redisql.y"
    {

	}
    break;

  case 101:
#line 521 "yacc_redisql.y"
    {
		(yyval.strDatabase) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 102:
#line 528 "yacc_redisql.y"
    {
		(yyval.strTable) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 103:
#line 535 "yacc_redisql.y"
    {
		(yyval.strColumn) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 104:
#line 542 "yacc_redisql.y"
    {
		(yyval.strIndex) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 105:
#line 549 "yacc_redisql.y"
    {
		(yyval.strVal) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 106:
#line 553 "yacc_redisql.y"
    {
		(yyval.strVal) = "";
	}
    break;

  case 107:
#line 560 "yacc_redisql.y"
    {
		setWhere((yyvsp[(2) - (2)].strVal));
	}
    break;

  case 108:
#line 564 "yacc_redisql.y"
    {
	}
    break;

  case 109:
#line 570 "yacc_redisql.y"
    { 
		(yyval.strVal) = (yyvsp[(1) - (1)].strVal);
	}
    break;

  case 110:
#line 574 "yacc_redisql.y"
    { 
   		char *pc = (char *)malloc(1000);
   		memset(pc, '\0', 1000);
   		sprintf(pc, "%s %s %s", (yyvsp[(1) - (3)].strVal), (yyvsp[(2) - (3)].strVal), (yyvsp[(3) - (3)].strVal));
   		(yyval.strVal) = pc;
   	}
    break;

  case 111:
#line 585 "yacc_redisql.y"
    {
   		FieldAlias st1, st3;
   		char *pc = (char *)malloc(1000);
   		memset(pc, '\0', 1000);
   		st1 = (yyvsp[(1) - (3)].stFieldAlias);
   		st3 = (yyvsp[(3) - (3)].stFieldAlias);
   		if (strcmp(st1.pcTableAlias, "") != 0)
   		{
   			strcat(pc, st1.pcTableAlias);
   			strcat(pc, ".");
   		}
   		strcat(pc, st1.pcField);
   		strcat(pc, " ");
   		strcat(pc, (yyvsp[(2) - (3)].strVal));
   		strcat(pc, " ");
   		if (st3.pcTableAlias != NULL)
   		{
	   		if (strcmp(st3.pcTableAlias, "") != 0)
	   		{
	   			strcat(pc, st3.pcTableAlias);
	   			strcat(pc, ".");
	   		}
   		}
   		strcat(pc, st3.pcField);
   		(yyval.strVal) = pc;
   	}
    break;

  case 112:
#line 612 "yacc_redisql.y"
    {
   		char *pc = (char *)malloc(1000);
   		memset(pc, '\0', 1000);
   		sprintf(pc, "( %s )", (yyvsp[(2) - (3)].strVal));
   		(yyval.strVal) = pc;
   	}
    break;

  case 113:
#line 621 "yacc_redisql.y"
    { (yyval.strVal) = "AND"; }
    break;

  case 114:
#line 622 "yacc_redisql.y"
    { (yyval.strVal) = "OR"; }
    break;

  case 115:
#line 627 "yacc_redisql.y"
    {
		setTop((yyvsp[(2) - (2)].nVal));
	}
    break;

  case 116:
#line 631 "yacc_redisql.y"
    {}
    break;

  case 117:
#line 636 "yacc_redisql.y"
    {
		setLimit((yyvsp[(2) - (4)].nVal), (yyvsp[(4) - (4)].nVal));
	}
    break;

  case 118:
#line 640 "yacc_redisql.y"
    {}
    break;


/* Line 1267 of yacc.c.  */
#line 2455 "y.tab.c"
      default: break;
    }
  YY_SYMBOL_PRINT ("-> $$ =", yyr1[yyn], &yyval, &yyloc);

  YYPOPSTACK (yylen);
  yylen = 0;
  YY_STACK_PRINT (yyss, yyssp);

  *++yyvsp = yyval;


  /* Now `shift' the result of the reduction.  Determine what state
     that goes to, based on the state we popped back to and the rule
     number reduced by.  */

  yyn = yyr1[yyn];

  yystate = yypgoto[yyn - YYNTOKENS] + *yyssp;
  if (0 <= yystate && yystate <= YYLAST && yycheck[yystate] == *yyssp)
    yystate = yytable[yystate];
  else
    yystate = yydefgoto[yyn - YYNTOKENS];

  goto yynewstate;


/*------------------------------------.
| yyerrlab -- here on detecting error |
`------------------------------------*/
yyerrlab:
  /* If not already recovering from an error, report this error.  */
  if (!yyerrstatus)
    {
      ++yynerrs;
#if ! YYERROR_VERBOSE
      yyerror (YY_("syntax error"));
#else
      {
	YYSIZE_T yysize = yysyntax_error (0, yystate, yychar);
	if (yymsg_alloc < yysize && yymsg_alloc < YYSTACK_ALLOC_MAXIMUM)
	  {
	    YYSIZE_T yyalloc = 2 * yysize;
	    if (! (yysize <= yyalloc && yyalloc <= YYSTACK_ALLOC_MAXIMUM))
	      yyalloc = YYSTACK_ALLOC_MAXIMUM;
	    if (yymsg != yymsgbuf)
	      YYSTACK_FREE (yymsg);
	    yymsg = (char *) YYSTACK_ALLOC (yyalloc);
	    if (yymsg)
	      yymsg_alloc = yyalloc;
	    else
	      {
		yymsg = yymsgbuf;
		yymsg_alloc = sizeof yymsgbuf;
	      }
	  }

	if (0 < yysize && yysize <= yymsg_alloc)
	  {
	    (void) yysyntax_error (yymsg, yystate, yychar);
	    yyerror (yymsg);
	  }
	else
	  {
	    yyerror (YY_("syntax error"));
	    if (yysize != 0)
	      goto yyexhaustedlab;
	  }
      }
#endif
    }



  if (yyerrstatus == 3)
    {
      /* If just tried and failed to reuse look-ahead token after an
	 error, discard it.  */

      if (yychar <= YYEOF)
	{
	  /* Return failure if at end of input.  */
	  if (yychar == YYEOF)
	    YYABORT;
	}
      else
	{
	  yydestruct ("Error: discarding",
		      yytoken, &yylval);
	  yychar = YYEMPTY;
	}
    }

  /* Else will try to reuse look-ahead token after shifting the error
     token.  */
  goto yyerrlab1;


/*---------------------------------------------------.
| yyerrorlab -- error raised explicitly by YYERROR.  |
`---------------------------------------------------*/
yyerrorlab:

  /* Pacify compilers like GCC when the user code never invokes
     YYERROR and the label yyerrorlab therefore never appears in user
     code.  */
  if (/*CONSTCOND*/ 0)
     goto yyerrorlab;

  /* Do not reclaim the symbols of the rule which action triggered
     this YYERROR.  */
  YYPOPSTACK (yylen);
  yylen = 0;
  YY_STACK_PRINT (yyss, yyssp);
  yystate = *yyssp;
  goto yyerrlab1;


/*-------------------------------------------------------------.
| yyerrlab1 -- common code for both syntax error and YYERROR.  |
`-------------------------------------------------------------*/
yyerrlab1:
  yyerrstatus = 3;	/* Each real token shifted decrements this.  */

  for (;;)
    {
      yyn = yypact[yystate];
      if (yyn != YYPACT_NINF)
	{
	  yyn += YYTERROR;
	  if (0 <= yyn && yyn <= YYLAST && yycheck[yyn] == YYTERROR)
	    {
	      yyn = yytable[yyn];
	      if (0 < yyn)
		break;
	    }
	}

      /* Pop the current state because it cannot handle the error token.  */
      if (yyssp == yyss)
	YYABORT;


      yydestruct ("Error: popping",
		  yystos[yystate], yyvsp);
      YYPOPSTACK (1);
      yystate = *yyssp;
      YY_STACK_PRINT (yyss, yyssp);
    }

  if (yyn == YYFINAL)
    YYACCEPT;

  *++yyvsp = yylval;


  /* Shift the error token.  */
  YY_SYMBOL_PRINT ("Shifting", yystos[yyn], yyvsp, yylsp);

  yystate = yyn;
  goto yynewstate;


/*-------------------------------------.
| yyacceptlab -- YYACCEPT comes here.  |
`-------------------------------------*/
yyacceptlab:
  yyresult = 0;
  goto yyreturn;

/*-----------------------------------.
| yyabortlab -- YYABORT comes here.  |
`-----------------------------------*/
yyabortlab:
  yyresult = 1;
  goto yyreturn;

#ifndef yyoverflow
/*-------------------------------------------------.
| yyexhaustedlab -- memory exhaustion comes here.  |
`-------------------------------------------------*/
yyexhaustedlab:
  yyerror (YY_("memory exhausted"));
  yyresult = 2;
  /* Fall through.  */
#endif

yyreturn:
  if (yychar != YYEOF && yychar != YYEMPTY)
     yydestruct ("Cleanup: discarding lookahead",
		 yytoken, &yylval);
  /* Do not reclaim the symbols of the rule which action triggered
     this YYABORT or YYACCEPT.  */
  YYPOPSTACK (yylen);
  YY_STACK_PRINT (yyss, yyssp);
  while (yyssp != yyss)
    {
      yydestruct ("Cleanup: popping",
		  yystos[*yyssp], yyvsp);
      YYPOPSTACK (1);
    }
#ifndef yyoverflow
  if (yyss != yyssa)
    YYSTACK_FREE (yyss);
#endif
#if YYERROR_VERBOSE
  if (yymsg != yymsgbuf)
    YYSTACK_FREE (yymsg);
#endif
  /* Make sure YYID is used.  */
  return YYID (yyresult);
}


#line 643 "yacc_redisql.y"

#include "lex.yy.c"
void yyerror(const char * s)
{
	printf("%s\n", s);
}

int redisql_parse(const char * sql)
{
	init();
	if(!sql)
	{
		printf("sql is null\n");
	}
	if (strcmp(sql, "") == 0)
	{
		return 0;
	}
	int len = strlen(sql);
	YY_BUFFER_STATE state = yy_scan_string(sql);
	yy_switch_to_buffer(state);
	int n = yyparse();
	yy_delete_buffer(state);
	return n;
}
