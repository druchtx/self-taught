# Emulator

## Using gcloud cli

```bash
# install the emulator component
gcloud components install cloud-firestore-emulator

# start the emulator
 gcloud emulators firestore start --host-port=0.0.0.0:8080 --export-on-exit=. --verbosity=debug

 # access the emulator
 export FIRESTORE_EMULATOR_HOST=localhost:8080
 ```

 ## Using firebae-tools

 ...

 ## Docker image

 ```bash
docker build -t druchtx/firestore-emulator .
docker run --name firestore-emulator -d -p8061:8061 -e PROJECT_ID=druchtx-local druchtx/firestore-emulator
 ```