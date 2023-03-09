# Copy the frontend part so it could be embedded into the executable
rm -r build; cp -r ..\ts.neose-mini.react-app\build\ ./build; go build -o neose-mini-controller.exe; ./neose-mini-controller.exe