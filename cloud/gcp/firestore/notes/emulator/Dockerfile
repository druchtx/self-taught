FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine

#  Firestore emulator needs JRE8+
RUN apk update && apk add openjdk21-jre-headless && \
    gcloud components install cloud-firestore-emulator --quiet && \
    gcloud components install beta --quiet

# Copy the entrypoint script and make it executable
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Declare the default data directory as a volume for persistence
VOLUME /root/.config/gcloud/emulators

# Set the entrypoint to our script
ENTRYPOINT ["/entrypoint.sh"]

