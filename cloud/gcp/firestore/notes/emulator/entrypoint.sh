#!/bin/sh
set -e

# Check for mandatory PROJECT_ID
# if [ -z "$PROJECT_ID" ]; then
#   echo "Error: The PROJECT_ID environment variable must be set." >&2
#   exit 1
# fi

# Set host and port, with a default for the port
HOST=${HOST:-0.0.0.0}
PORT=${PORT:-8061}

# Execute the emulator using the beta command, storing data in the default location
exec gcloud beta emulators firestore start \
  # --project="$PROJECT_ID" \
  --host-port="$HOST:$PORT" \
  --quiet --verbosity=debug