#!/bin/bash
echo "initialization turbo jobs services."
cd /var/turbo-bin 
./main-jobs-production production
echo "[production] turbo bpo jobs  sevices is running..."