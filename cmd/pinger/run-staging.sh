#!/bin/bash
echo "initialization turbo jobs pinger services."
cd /var/turbo-bpo-be/cmd/pinger
./main-jobs-ping-staging staging
echo "[staging] turbo bpo jobs pinger sevices is running..."