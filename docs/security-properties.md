# Security Properties
- Integrity: append-only RGAC + periodic Merkle roots
- Authenticity: hybrid PQC (Kyber/Falcon) for node auth/KE
- Non-repudiation: EIR + HVET with UTC anchors
- Freshness: 30s key rotation; acceptance windows
- Availability: 2-node mesh, telemetry observability
Threat Model: replay, reordering, equivocation, tamper; mitigations documented.
