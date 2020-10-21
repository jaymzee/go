#include "textflag.h"

TEXT ·AddInt(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), DX
	ADDQ AX, DX
	MOVQ DX, ret+16(FP)
	RET

TEXT ·MulInt(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), DX
	IMULQ DX
	MOVQ AX, ret+16(FP)
	MOVQ DX, ret+24(FP)
	RET

TEXT ·MultFloat32(SB),NOSPLIT,$0
	MOVSS x+0(FP), X0
	MOVSS y+4(FP), X1
	MULSS X1, X0
	MOVSS X0, ret+8(FP)
	RET

TEXT ·MultFloat64(SB),NOSPLIT,$0
	MOVSD x+0(FP), X0
	MOVSD y+8(FP), X1
	MULSD X1, X0
	MOVSD X0, ret+16(FP)
	RET
