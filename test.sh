#!/bin/bash
cd /Users/mioi/co/whimsy
echo "Testing whimsy library..."
go test -v
echo ""
echo "Running example..."
cd example && go run main.go