#include "textflag.h"

// Matrix4x4v computes the product of a Matrix and a vector
// rows *[4][4]float32
// x *[4]float32
// y *[4]float32
TEXT ·Matrix4x4v(SB),NOSPLIT,$0
	MOVQ	rows+0(FP), DI
	MOVQ	x+8(FP), SI
	MOVQ	y+16(FP), BX
	MOVAPS	0(SI), X3
	MOVAPS	0(DI), X0
	MULPS	X3, X0		// X0 ← rows[0] * x
	MOVAPS	16(DI), X1
	MULPS	X3, X1		// X1 ← rows[1] * x
	HADDPS	X1, X0		// X0 ← partial sum of products first two rows
	MOVAPS	32(DI), X1
	MULPS	X3, X1		// X1 ← rows[2] * x
	MOVAPS	48(DI), X2
	MULPS	X3, X2		// X2 ← rows[3] * x
	HADDPS	X2, X1		// X1 ← partial sum of products next two rows
	HADDPS	X1, X0
	MOVAPS	X0, 0(BX)	// y ← total sum of products
	RET
