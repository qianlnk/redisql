/* A Bison parser, made by GNU Bison 2.3.  */

/* Skeleton interface for Bison's Yacc-like parsers in C

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
/* Line 1529 of yacc.c.  */
#line 185 "y.tab.h"
	YYSTYPE;
# define yystype YYSTYPE /* obsolescent; will be withdrawn */
# define YYSTYPE_IS_DECLARED 1
# define YYSTYPE_IS_TRIVIAL 1
#endif

extern YYSTYPE yylval;

