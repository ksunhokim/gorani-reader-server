git submodule update --recursive --remote
cd app
carthage update
cd FolioReaderKit
carthage update
echo "swift project configuration complete"
echo "open app/app.xcworkspace to build the app"
