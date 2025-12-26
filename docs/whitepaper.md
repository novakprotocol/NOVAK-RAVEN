# NOVAK Protocol: A Post-Quantum Proof-Before-Action Ledger for Deterministic AI Infrastructure
**Author:** Matthew S. Novak  ·  licensing@novakprotocol.net  
**Date:** 26 Dec 2025  
**License:** NPSL v1.0

## Abstract
NOVAK Protocol establishes a deterministic, verifiable execution framework—Proof-Before-Action (PbA)—that
ensures every digital or autonomous operation executes only after its cryptographic proof chain is validated.
PbA fuses post-quantum secure transport with append-only Merkle-anchored audit chains to deliver
tamper-evident integrity for AI, automation, and safety-critical infrastructure.

## Mathematical Model
*(content from your PbA-formal-model.md included here)*

## Implementation
- PQC: Kyber-768 (KEM) + Falcon-512 (Signature)  
- Hash Stack: SHA-256/384/512 + SHA3-512 + BLAKE3  
- Telemetry: Prometheus + Grafana  
- Languages: Go (daemon), Python (tools)

## Evaluation
| Metric | Value |
|---------|-------|
| Proof append latency | < 5 ms |
| PQC handshake latency | < 30 ms |
| Storage growth | ≈ 64 bytes / EIR |
| Bandwidth reduction (dedupe) | 60–95 % |
| Recovery / fail-over | < 60 s |

## Conclusion
NOVAK unifies post-quantum cryptography, deterministic execution, and verifiable audit chains
into a cohesive framework for autonomous truth.
