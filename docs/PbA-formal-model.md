# NOVAK Proof-Before-Action (PbA) — Formal Model

## Data
EIR_i = ⟨R_i, D_i, O_i, T_i, HVET_i⟩
HVET_i = H( H(R_i) || H(D_i) || H(O_i) || H(T_i) )

RGAC chain:
C_0 = H("NOVAK|RGAC|v1")
C_i = H( C_{i-1} || encode(EIR_i) )

Merkle checkpoints (batch of size N):
Root_k = Merkle( EIR_{kN+1} … EIR_{kN+N} )

## Invariants
- verify(HVET_i) = true ⇒ execute(action_i)
- UTC time anchored; monotonic acceptance window
- append-only chain; inclusion proofs via RGAC & Merkle

## Timing
- Heartbeat 5s; PQC rotation 30s; checkpoint N=1000
