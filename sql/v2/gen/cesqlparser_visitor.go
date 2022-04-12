<<<<<<< HEAD
// Code generated from CESQLParser.g4 by ANTLR 4.10.1. DO NOT EDIT.
=======
/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

// Code generated from CESQLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.
>>>>>>> bcb75b3 (wip: feat: supports data in sql)

package gen // CESQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CESQLParserParser.
type CESQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CESQLParserParser#cesql.
	VisitCesql(ctx *CesqlContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#binaryComparisonExpression.
	VisitBinaryComparisonExpression(ctx *BinaryComparisonExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#functionInvocationExpression.
	VisitFunctionInvocationExpression(ctx *FunctionInvocationExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#unaryLogicExpression.
	VisitUnaryLogicExpression(ctx *UnaryLogicExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#inExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#atomExpression.
	VisitAtomExpression(ctx *AtomExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#existsExpression.
	VisitExistsExpression(ctx *ExistsExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#binaryLogicExpression.
	VisitBinaryLogicExpression(ctx *BinaryLogicExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#likeExpression.
	VisitLikeExpression(ctx *LikeExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#binaryMultiplicativeExpression.
	VisitBinaryMultiplicativeExpression(ctx *BinaryMultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#dataExistsExpression.
	VisitDataExistsExpression(ctx *DataExistsExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#unaryNumericExpression.
	VisitUnaryNumericExpression(ctx *UnaryNumericExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#subExpression.
	VisitSubExpression(ctx *SubExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#binaryAdditiveExpression.
	VisitBinaryAdditiveExpression(ctx *BinaryAdditiveExpressionContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#integerAtom.
	VisitIntegerAtom(ctx *IntegerAtomContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#dataIdentifierAtom.
	VisitDataIdentifierAtom(ctx *DataIdentifierAtomContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#identifierAtom.
	VisitIdentifierAtom(ctx *IdentifierAtomContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#dataIdentifier.
	VisitDataIdentifier(ctx *DataIdentifierContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#functionIdentifier.
	VisitFunctionIdentifier(ctx *FunctionIdentifierContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#functionParameterList.
	VisitFunctionParameterList(ctx *FunctionParameterListContext) interface{}

	// Visit a parse tree produced by CESQLParserParser#setExpression.
	VisitSetExpression(ctx *SetExpressionContext) interface{}
}
