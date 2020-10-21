#include "textflag.h"

/* R10 and R11 are reserved by the compiler */
/* R10 is the g (goroutine) structure pointer */

TEXT ·AddInt(SB),NOSPLIT,$0
	MOVW x+0(FP), R1
	MOVW y+4(FP), R2
	ADD R2, R1, R0
	MOVW R0, ret+8(FP)
	RET

TEXT ·MultInt(SB),NOSPLIT,$0
	MOVW x+0(FP), R1
	MOVW y+4(FP), R2
	MUL R2, R1, R0
	EOR  R1,R1		// cheating by clearing the upper 32 bits
	MOVW R0, ret+8(FP)
	MOVW R1, ret+12(FP)
	RET

TEXT ·MultFloat32(SB),NOSPLIT,$0
	MOVF x+0(FP), F1
	MOVF y+4(FP), F2
	MULF F2, F1, F0
	MOVF F0, ret+8(FP)
	RET

TEXT ·MultFloat64(SB),NOSPLIT,$0
	MOVD x+0(FP), F1
	MOVD y+8(FP), F2
	MULD F2, F1, F0
	MOVD F0, ret+16(FP)
	RET
