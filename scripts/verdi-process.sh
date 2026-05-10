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

if [ -z "$1" ]; then
    echo -e "${C_YELLOW}⚠️  Usage: verdi-process <target_directory>${C_DEF}"
    echo -e "   Example: verdi-process \"/music/Artist/Album\"\n"
    exit 1
fi

TARGET_DIR=$(echo "$1" | tr -s '/')
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
TARGET_DIR_440="${ORIGINAL_DIR} [440 Hz]"
TARGET_DIR_432="${ORIGINAL_DIR} [432 Hz]"

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
find "$TARGET_DIR_440" -type f -name '*.flac' | while read file; do
    if ! metaflac --show-tag=VERSION "$file" | grep -qi VERSION; then
        metaflac --set-tag="VERSION=440 Hz" "$file"
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
VERDI_IN="" VERDI_OUT="" /usr/local/bin/verdi-pitch-engine -in "$TARGET_DIR_440" -out "$TARGET_DIR_432"
echo -e "${C_DIM}------------------------------------------------${C_DEF}"

echo -e "\n🖼️  ${C_BLUE}Step 5: Migrating Sidecar Metadata (Art, Lyrics, PDFs)...${C_DEF}"
(
    cd "$TARGET_DIR_440" || exit
    # Find all non-FLAC files and copy them to the 432 Hz directory, preserving folder structure
    find . -type f ! -iname "*.flac" -exec cp --parents "{}" "$TARGET_DIR_432/" \;
)
echo -e "   ${C_GREEN}✔ Sidecar migration complete.${C_DEF}"

echo -e "\n✨ ${C_GREEN}PROCESSING COMPLETE FOR: ${ALBUM_NAME}${C_DEF} ✨\n"
