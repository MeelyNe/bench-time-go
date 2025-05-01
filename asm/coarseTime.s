TEXT ·GetCoarseTime(SB),$0-12
    MOVL    $5, DI
    LEAQ    ret+0(FP), SI
    MOVQ    runtime·vdsoClockgettimeSym(SB), AX
    CALL    AX
    MOVQ    0(SI), AX
    MOVL    8(SI), DX
    MOVQ    AX, sec+0(FP)
    MOVL    DX, nsec+8(FP)
    RET
