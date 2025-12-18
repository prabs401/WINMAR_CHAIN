#!/bin/bash

# Configuration
BACKUP_DIR="backups"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
FILENAME="wnc_backup_$TIMESTAMP.tar.gz"

# Create backup directory if it doesn't exist
mkdir -p $BACKUP_DIR

# Create the backup
# We backup:
# 1. chaindata.json (The most important file containing balances and blocks)
# 2. config/ (Your node configuration)
# 3. build/wnc-node (The program binary itself)

echo "📦 Starting Winmar Chain Backup..."

if [ -f "chaindata.json" ]; then
    tar -czvf "$BACKUP_DIR/$FILENAME" chaindata.json config/ build/wnc-node
    echo "✅ Backup created successfully: $BACKUP_DIR/$FILENAME"
    echo "💡 Please download this file to your local computer!"
else
    echo "⚠️  Warning: chaindata.json not found. Have you started the node yet?"
fi
