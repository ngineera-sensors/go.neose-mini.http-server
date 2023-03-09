# Copy the frontend part so it could be embedded into the executable
rm -r build;
cp -r ..\ts.neose-mini.react-app\build\ ./build;

go-winres make;
go build -o neose-mini-controller.exe -ldflags '-s';

echo "Attempting to publish the tool on NAS"
cp neose-mini-controller.exe \\srv-shares\SHARES\commun\Kirill\_neose-mini-controller\