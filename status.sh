#!/bin/bash
BIN=/opt/novak/bin/novakd
echo "=== Novak Status $(date -u) ==="
systemctl is-active --quiet novakd && echo "Service: ACTIVE" || echo "Service: INACTIVE"
echo "Commit: $(git -C /opt/novak-src rev-parse --short HEAD)"
[ -f "$BIN" ] && echo "Binary: $(stat -c %y "$BIN" | cut -d. -f1) | Size: $(du -h "$BIN" | cut -f1)"
tail -n3 /var/log/novak/novakd.log 2>/dev/null || true
