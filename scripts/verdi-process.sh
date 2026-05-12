#!/bin/bash
# scripts/verdi-process.sh
# A self-contained orchestration script to run VerdiPitchEngine inside Container Station

# Colors
C_DEF="\033[0m"
C_CYAN="\033[1;36m"
C_MAGENTA="\033[1;35m"
C_GREEN="\033[1;32m"
C_YELLOW="\033[1;33m"
C_BLUE="\033[1;34m"
C_DIM="\033[1;30m"

echo -e "\n${C_MAGENTA}🚀 INITIATING VERDI PITCH ENGINE 🚀${C_DEF}"
echo -e "${C_DIM}------------------------------------------------${C_DEF}"

# Execution Blocker: Ensure media servers are stopped to prevent metadata interference
echo -e "🔍 Verifying media server status..."
if ps | grep -v grep | grep -qE "RoonServer|RoonAppliance|Plex Media Server|mserver|twonkymediaserver|jellyfin|emby-server|squeezeboxserver"; then
    echo -e "${C_MAGENTA}🚨 ERROR: An active Media Server (Roon, Plex, MinimServer, etc.) was detected!${C_DEF}"
    echo -e "${C_MAGENTA}🛑 Verdi Pitch Engine will not continue due to possible database interference.${C_DEF}"
    echo -e "${C_YELLOW}💡 Please stop any active media servers in the QNAP App Center before proceeding.${C_DEF}"
    exit 1
fi

if [ -z "$1" ]; then
    echo -e "${C_YELLOW}⚠️  Usage: verdi-process <target_directory> [strategy]${C_DEF}"
    echo -e "   Example: verdi-process \"/music/Artist/Album\" asetrate\n"
    exit 1
fi

TARGET_DIR=$(echo "$1" | tr -s '/')
STRATEGY="${2:-rubberband}"
TARGET_DIR="${TARGET_DIR%/}"

# Auto-correct full QNAP host paths to the internal container mount path
if [[ "$TARGET_DIR" == "/share/Multimedia/Audio/Music"* ]]; then
    TARGET_DIR="${TARGET_DIR/\/share\/Multimedia\/Audio\/Music/\/music}"
    echo -e "${C_YELLOW}💡 Auto-mapped host path to internal container path: ${TARGET_DIR}${C_DEF}\n"
# Auto-correct relative paths to default to the internal container mount path
elif [[ "$TARGET_DIR" != /* ]]; then
    TARGET_DIR="/music/${TARGET_DIR}"
    echo -e "${C_YELLOW}💡 Auto-mapped relative path to: ${TARGET_DIR}${C_DEF}\n"
fi

ALBUM_NAME=$(basename "$TARGET_DIR")

if [[ "$TARGET_DIR" == *" [440 Hz]" ]] || [[ "$TARGET_DIR" == *" [432 Hz]" ]]; then
    # Automatically strip the tag so we can use the base directory name for orchestration
    TARGET_DIR="${TARGET_DIR% \[440 Hz\]}"
    TARGET_DIR="${TARGET_DIR% \[432 Hz\]}"
    echo -e "${C_YELLOW}💡 Auto-detected version tag. Resolving to base directory: ${TARGET_DIR}${C_DEF}\n"
fi

ALBUM_NAME=$(basename "$TARGET_DIR")

ORIGINAL_DIR="$TARGET_DIR"
# Hide the 440 Hz backup so Roon ignores it to prevent duplicates
TARGET_DIR_440="$(dirname "$TARGET_DIR")/.$(basename "$TARGET_DIR") [440 Hz]"
# Output directly back into the exact original path to retain Roon database entries (favorites, playlists)
TARGET_DIR_432="$ORIGINAL_DIR"

if [ ! -d "$ORIGINAL_DIR" ] && [ ! -d "$TARGET_DIR_440" ]; then
    echo -e "${C_YELLOW}❌ Error: Directory not found: ${ORIGINAL_DIR} (nor its 440 Hz backup)${C_DEF}"
    exit 1
fi

echo -e "🎶 ${C_CYAN}Target Album:${C_DEF} ${ALBUM_NAME}"

if [ -d "$ORIGINAL_DIR" ]; then
    echo -e "\n📦 ${C_BLUE}Step 1: Securing Original Master...${C_DEF}"
    echo -e "   Moving to: ${C_DIM}${TARGET_DIR_440}${C_DEF}"
    mv "$ORIGINAL_DIR" "$TARGET_DIR_440"
else
    echo -e "\n📦 ${C_BLUE}Step 1: Original Master already secured... skipping move.${C_DEF}"
fi

echo -e "\n🏷️  ${C_BLUE}Step 2: Enforcing Version Metadata (440 Hz)...${C_DEF}"
find "$TARGET_DIR_440" -type d \( -name ".@__*" -o -name ".AppleDouble" \) -prune -o -type f -name '*.flac' -print | while read file; do
    if ! metaflac --show-tag=VERSION "$file" | grep -qi VERSION; then
        metaflac --preserve-modtime --set-tag="VERSION=440 Hz" "$file"
    fi
done
echo -e "   ${C_GREEN}✔ Metadata enforcement complete.${C_DEF}"

echo -e "\n📂 ${C_BLUE}Step 3: Provisioning Output Enclave...${C_DEF}"
echo -e "   Creating: ${C_DIM}${TARGET_DIR_432}${C_DEF}"
mkdir -p "$TARGET_DIR_432"
chmod 777 "$TARGET_DIR_432"

# Clean up any residual temporary files from forcefully aborted runs
find "$TARGET_DIR_432" -type f -name '*.tmp' -delete 2>/dev/null || true

echo -e "\n⚙️  ${C_BLUE}Step 4: Executing High-Fidelity DSP TSM...${C_DEF}"
echo -e "${C_DIM}------------------------------------------------${C_DEF}"
VERDI_IN="" VERDI_OUT="" VERDI_STRATEGY="$STRATEGY" /usr/local/bin/verdi-pitch-engine -in "$TARGET_DIR_440" -out "$TARGET_DIR_432"
echo -e "${C_DIM}------------------------------------------------${C_DEF}"

echo -e "\n🖼️  ${C_BLUE}Step 5: Migrating Sidecar Metadata (Art, Lyrics, PDFs)...${C_DEF}"
(
    cd "$TARGET_DIR_440" || exit
    # Find all non-FLAC files and copy them to the 432 Hz directory, preserving timestamps and folder structure
    find . -type d \( -name ".@__*" -o -name ".AppleDouble" \) -prune -o -type f ! -iname "*.flac" -exec cp -p --parents "{}" "$TARGET_DIR_432/" \;
)
echo -e "   ${C_GREEN}✔ Sidecar migration complete.${C_DEF}"

echo -e "\n🏷️  ${C_BLUE}Step 6: Enforcing Version Metadata (432 Hz)...${C_DEF}"
if [ "$STRATEGY" == "asetrate" ]; then
    VERSION_STRING="432 Hz (Asetrate)"
else
    VERSION_STRING="432 Hz"
fi

find "$TARGET_DIR_432" -type d \( -name ".@__*" -o -name ".AppleDouble" \) -prune -o -type f -name '*.flac' -print | while read -r file; do
    metaflac --preserve-modtime --remove-tag=VERSION "$file"
    metaflac --preserve-modtime --set-tag="VERSION=${VERSION_STRING}" "$file"
done
echo -e "   ${C_GREEN}✔ Output metadata enforcement complete (${VERSION_STRING}).${C_DEF}"

echo -e "\n🕰️  ${C_BLUE}Step 7: Synchronizing File & Directory Timestamps...${C_DEF}"
find "$TARGET_DIR_440" -type d \( -name ".@__*" -o -name ".AppleDouble" \) -prune -o -print | while read -r path; do
    rel_path="${path#$TARGET_DIR_440}"
    if [ -e "$TARGET_DIR_432$rel_path" ]; then
        touch -r "$path" "$TARGET_DIR_432$rel_path"
    fi
done
echo -e "   ${C_GREEN}✔ Original timestamps preserved to prevent Roon 'Recent' flagging.${C_DEF}"

echo -e "\n✨ ${C_GREEN}PROCESSING COMPLETE FOR: ${ALBUM_NAME}${C_DEF} ✨\n"
