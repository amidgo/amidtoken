cd eth-net && ./restore.sh
gnome-terminal -- ./bootnode.sh && gnome-terminal -- ./node1.sh && gnome-terminal -- ./node2.sh
cd ../ && go run main.go