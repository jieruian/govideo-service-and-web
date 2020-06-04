#! /bin/bash

# Build web UI
cd ~/go/src/govideo/video_server/web/
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui/web
cp -R ~/go/src/govideo/video_server/templates ~/go/bin/video_server_web_ui/web