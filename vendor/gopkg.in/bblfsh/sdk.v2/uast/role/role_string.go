// Code generated by "stringer -type=Role"; DO NOT EDIT.

package role

import "strconv"

const _Role_name = "InvalidIdentifierQualifiedOperatorBinaryUnaryLeftRightInfixPostfixBitwiseBooleanUnsignedLeftShiftRightShiftOrXorAndExpressionStatementEqualNotLessThanLessThanOrEqualGreaterThanGreaterThanOrEqualIdenticalContainsIncrementDecrementNegativePositiveDereferenceTakeAddressFileAddSubstractMultiplyDivideModuloPackageDeclarationImportPathnameAliasFunctionBodyNameReceiverArgumentValueArgsListBaseImplementsInstanceSubtypeSubpackageModuleFriendWorldIfConditionThenElseSwitchCaseDefaultForInitializationUpdateIteratorWhileDoWhileBreakContinueGotoBlockScopeReturnTryCatchFinallyThrowAssertCallCalleePositionalNoopLiteralByteByteStringCharacterListMapNullNumberRegexpSetStringTupleTypeEntryKeyPrimitiveAssignmentThisCommentDocumentationWhitespaceIncompleteUnannotatedVisibilityAnnotationAnonymousEnumerationArithmeticRelationalVariable"

var _Role_index = [...]uint16{0, 7, 17, 26, 34, 40, 45, 49, 54, 59, 66, 73, 80, 88, 97, 107, 109, 112, 115, 125, 134, 139, 142, 150, 165, 176, 194, 203, 211, 220, 229, 237, 245, 256, 267, 271, 274, 283, 291, 297, 303, 310, 321, 327, 335, 340, 348, 352, 356, 364, 372, 377, 385, 389, 399, 407, 414, 424, 430, 436, 441, 443, 452, 456, 460, 466, 470, 477, 480, 494, 500, 508, 513, 520, 525, 533, 537, 542, 547, 553, 556, 561, 568, 573, 579, 583, 589, 599, 603, 610, 614, 624, 633, 637, 640, 644, 650, 656, 659, 665, 670, 674, 679, 682, 691, 701, 705, 712, 725, 735, 745, 756, 766, 776, 785, 796, 806, 816, 824}

func (i Role) String() string {
	if i < 0 || i >= Role(len(_Role_index)-1) {
		return "Role(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
