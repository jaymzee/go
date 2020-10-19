#include "textflag.h"

TEXT ·addi64(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), DX
	ADDQ AX, DX
	MOVQ DX, ret+16(FP)
	RET

TEXT ·muli64(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), DX
	IMULQ DX
	MOVQ AX, ret+16(FP)
	MOVQ DX, ret+24(FP)
	RET
