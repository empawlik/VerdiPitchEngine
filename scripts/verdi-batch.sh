#!/bin/bash
# -----------------------------------------------------------------------------
# Verdi Pitch Engine - Batch Processing Orchestrator
# Executes batch processing for a specified root folder.
# -----------------------------------------------------------------------------

set -e

# Styling
C_CYAN='\033[1;36m'
C_GREEN='\033[1;32m'
C_RED='\033[1;31m'
C_YELLOW='\033[1;33m'
C_DEF='\033[0m'
C_DIM='\033[2m'

echo -e "\n${C_CYAN}🚀 INITIATING VERDI BATCH ENGINE 🚀${C_DEF}"
echo -e "${C_DIM}------------------------------------------------${C_DEF}"

# Argument parsing
if [ $# -lt 1 ]; then
    echo -e "${C_RED}❌ Error: Missing arguments.${C_DEF}"
    echo -e "Usage: verdi-batch <root-folder> [limit]"
    echo -e "Example: verdi-batch \"Artist\" 10"
    exit 1
fi

RAW_ROOT="$1"
LIMIT="${2:-all}"

if [[ "$LIMIT" != "all" ]] && ! [[ "$LIMIT" =~ ^[0-9]+$ ]]; then
    echo -e "${C_RED}❌ Error: Limit must be an integer or 'all'.${C_DEF}"
    exit 1
fi

# Convert 'all' or '0' to a very large number for comparison logic
if [[ "$LIMIT" == "all" ]] || [[ "$LIMIT" == "0" ]]; then
    MAX_ALBUMS=999999
    LIMIT_LABEL="All"
else
    MAX_ALBUMS="$LIMIT"
    LIMIT_LABEL="$LIMIT"
fi

# Map to container mount
BASE_DIR="/music"
TARGET_ROOT="${BASE_DIR}/${RAW_ROOT}"

# Strip leading slash if user accidentally included it for relative pathing
if [[ "$RAW_ROOT" == /* ]] && [[ ! "$RAW_ROOT" == /music* ]]; then
    # if it's an absolute path but not /music, try stripping /share/...
    CLEAN_PATH=$(echo "$RAW_ROOT" | sed -E 's|^/share/[^/]+/[^/]+/[^/]+/(.*)|\1|')
    TARGET_ROOT="${BASE_DIR}/${CLEAN_PATH}"
fi

if [ ! -d "$TARGET_ROOT" ]; then
    echo -e "${C_RED}❌ Error: Root directory not found: $TARGET_ROOT${C_DEF}"
    exit 1
fi

echo -e "💡 Scanning Root: ${C_DIM}$TARGET_ROOT${C_DEF}"
echo -e "💡 Limit: ${C_DIM}${LIMIT_LABEL} albums${C_DEF}"

# Find all directories containing .flac files
# - We use find to locate directories with .flac files
# - We filter out [440 Hz] and [432 Hz] directories to only target raw master directories
# - We sort them alphabetically
# - We take the top $LIMIT directories

echo -e "\n🔍 Scanning for unprocessed FLAC albums..."

# Temporary file to hold unique album directories
TMP_LIST=$(mktemp)

# Find directories containing FLACs, sort unique
find "$TARGET_ROOT" -type f -name '*.flac' -exec dirname {} \; | sort -u > "$TMP_LIST"

ALBUM_COUNT=0
declare -a ALBUMS_TO_PROCESS

while IFS= read -r dir; do
    # Skip if the directory name contains [440 Hz] or [432 Hz] (it's already a versioned artifact)
    if [[ "$dir" == *"[440 Hz]"* ]] || [[ "$dir" == *"[432 Hz]"* ]]; then
        continue
    fi
    
    # Skip if the hidden [440 Hz] backup directory already exists!
    parent_dir=$(dirname "$dir")
    base_name=$(basename "$dir")
    if [ -d "${parent_dir}/.${base_name} [440 Hz]" ]; then
        continue
    fi

    # Add to our target list
    ALBUMS_TO_PROCESS+=("$dir")
    ALBUM_COUNT=$((ALBUM_COUNT + 1))
    
    # Break if we reach the limit
    if [ "$ALBUM_COUNT" -ge "$MAX_ALBUMS" ]; then
        break
    fi
done < "$TMP_LIST"

rm -f "$TMP_LIST"

if [ ${#ALBUMS_TO_PROCESS[@]} -eq 0 ]; then
    echo -e "${C_GREEN}✔ No pending albums found in $TARGET_ROOT. Everything is up to date!${C_DEF}"
    exit 0
fi

echo -e "📌 Found ${C_GREEN}${#ALBUMS_TO_PROCESS[@]}${C_DEF} albums to process:"
for idx in "${!ALBUMS_TO_PROCESS[@]}"; do
    # Print the relative path for cleaner logging
    REL_PATH="${ALBUMS_TO_PROCESS[$idx]#$BASE_DIR/}"
    echo -e "  $((idx + 1)). ${C_DIM}${REL_PATH}${C_DEF}"
done

echo -e "\n⏳ Beginning batch execution sequence in 5 seconds... (Ctrl+C to abort)"
sleep 5

# Execute verdi-process on each album
for idx in "${!ALBUMS_TO_PROCESS[@]}"; do
    ALBUM_DIR="${ALBUMS_TO_PROCESS[$idx]}"
    REL_PATH="${ALBUM_DIR#$BASE_DIR/}"
    
    echo -e "\n================================================================"
    echo -e "▶️  Processing Album $((idx + 1)) of ${#ALBUMS_TO_PROCESS[@]}: ${C_CYAN}${REL_PATH}${C_DEF}"
    echo -e "================================================================"
    
    # Call verdi-process using the relative path so it correctly maps it
    /usr/local/bin/verdi-process "$REL_PATH"
    
    echo -e "\n✅ Finished: ${C_DIM}${REL_PATH}${C_DEF}"
    sleep 2
done

echo -e "\n🎉 ${C_GREEN}BATCH EXECUTION COMPLETE${C_DEF} 🎉"
